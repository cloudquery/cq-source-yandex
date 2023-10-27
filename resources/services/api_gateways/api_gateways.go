package api_gateways

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func ApiGateways() *schema.Table {
	return &schema.Table{
		Name:      "yandex_api_gateways_api_gateways",
		Resolver:  fetchApiGateways,
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
				Name:     "status",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Domain"),
			},
			{
				Name:     "log_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogGroupId"),
			},
			{
				Name:     "attached_domains",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AttachedDomains"),
			},
			{
				Name:     "connectivity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Connectivity"),
			},
			{
				Name:     "log_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogOptions"),
			},
		},
	}
}
