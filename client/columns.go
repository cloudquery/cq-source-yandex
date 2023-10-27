package client

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var (
	CloudIDColumn = schema.Column{
		Name:       "cloud_id",
		Type:       arrow.BinaryTypes.String,
		Resolver:   ResolveMultiplexedResourceID,
		PrimaryKey: true,
	}
)
