package vpc

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func SecurityGroups() *schema.Table {
	return &schema.Table{
		Name:      "yandex_vpc_security_groups",
		Resolver:  fetchSecurityGroups,
		Multiplex: client.MultiplexByFolder,
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
				Description: `Resource ID`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "folder_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FolderId"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreatedAt"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "network_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkId"),
			},
			{
				Name:     "status",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Rules"),
			},
			{
				Name:     "default_for_network",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultForNetwork"),
			},
		},
	}
}
