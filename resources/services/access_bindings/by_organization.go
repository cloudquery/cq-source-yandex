package access_bindings

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func ByOrganization() *schema.Table {
	return &schema.Table{
		Name:      "yandex_access_bindings_by_organization",
		Resolver:  fetchByOrganization,
		Multiplex: client.MultiplexByOrg,
		Columns: []schema.Column{
			{
				Name:     "organization_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveMultiplexedResourceID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "role_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "subject_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject.Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "subject_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject.Type"),
			},
		},
	}
}
