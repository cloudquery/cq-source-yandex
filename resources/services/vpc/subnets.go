package vpc

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Subnets() *schema.Table {
	return &schema.Table{
		Name:      "yandex_vpc_subnets",
		Resolver:  fetchSubnets,
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
				Name:     "zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ZoneId"),
			},
			{
				Name:     "v4_cidr_blocks",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("V4CidrBlocks"),
			},
			{
				Name:     "v6_cidr_blocks",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("V6CidrBlocks"),
			},
			{
				Name:     "route_table_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteTableId"),
			},
			{
				Name:     "dhcp_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DhcpOptions"),
			},
		},
	}
}
