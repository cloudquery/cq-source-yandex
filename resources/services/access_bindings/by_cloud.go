package access_bindings

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func ByCloud() *schema.Table {
	return &schema.Table{
		Name:      "yandex_access_bindings_by_cloud",
		Multiplex: client.MultiplexByCloud,
		Transform: transformers.TransformWithStruct(&access.AccessBinding{},
			transformers.WithUnwrapStructFields("Subject"),
			transformers.WithPrimaryKeys("RoleId", "Subject.Id")),
		Columns:  schema.ColumnList{client.CloudIDColumn},
		Resolver: fetchByCloud,
	}
}
