package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/thoas/go-funk"
)

func ResolveMultiplexedResourceID(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return resource.Set(c.Name, client.MultiplexedResourceId)
}

func ResolveProtoTimestamp(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		ts, ok := data.(*timestamppb.Timestamp)
		if !ok {
			return fmt.Errorf("unexpected type, wanted \"*timestamppb.Timestamp\", have \"%T\"", data)
		}
		return resource.Set(c.Name, ts.AsTime())
	}
}

type LabeledStruct interface {
	GetLabels() map[string]string
}

func ResolveLabels(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(LabeledStruct)
	rawLabels := r.GetLabels()

	labels := make(map[string]*string, len(rawLabels))
	for k, v := range rawLabels {
		labels[k] = &v
	}

	return resource.Set(c.Name, labels)
}

type Converter func(from interface{}) (interface{}, error)

func dictConverter(from interface{}) (interface{}, error) {
	if from == nil {
		return nil, nil
	}

	var mapValue map[string]string
	var ok bool
	if mapValue, ok = from.(map[string]string); !ok {
		return nil, errors.New("not a value of type map[string]string")
	}

	res := make(map[string]*string, len(mapValue))
	for k, v := range mapValue {
		res[k] = &v
	}
	return res, nil
}

func timeConverter(from interface{}) (interface{}, error) {
	if from == nil {
		return nil, nil
	}

	var protots *timestamp.Timestamp
	var ok bool
	if protots, ok = from.(*timestamp.Timestamp); !ok {
		return nil, errors.New("not a value of type *timestamp.Timestamp")
	}

	if !protots.IsValid() {
		return nil, errors.New("invalid proto timestamp")
	}

	ts := protots.AsTime()
	return &ts, nil
}

func resolveAsSmth(converter Converter) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		resolver := resolvePathAsSmth(strcase.ToCamel(c.Name), converter)
		return resolver(ctx, meta, resource, c)
	}
}

func resolvePathAsSmth(path string, converter Converter) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		value := funk.Get(resource.Item, path, funk.WithAllowZero())
		if value == nil {
			meta.(*Client).Logger().Trace().Str("column", c.Name).Msg("no column value specified")
			return resource.Set(c.Name, nil)
		}

		res, err := converter(value)
		if err != nil {
			return fmt.Errorf("error converting path %s: %w", path, err)
		}

		meta.(*Client).Logger().Trace().Str("column", c.Name).Interface("value", res).Msg("setting column value")
		return resource.Set(c.Name, res)
	}
}

var (
	ResolveAsDict schema.ColumnResolver = resolveAsSmth(dictConverter)
	ResolveAsTime schema.ColumnResolver = resolveAsSmth(timeConverter)
)

func ResolvePathAsDict(path string) schema.ColumnResolver {
	return resolvePathAsSmth(path, dictConverter)
}

func ResolvePathAsTime(path string) schema.ColumnResolver {
	return resolvePathAsSmth(path, timeConverter)
}

func EnumPathResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		if stringer, ok := funk.Get(r.Item, path, funk.WithAllowZero()).(fmt.Stringer); ok {
			return r.Set(c.Name, stringer.String())
		} else {
			return nil
		}
	}
}
