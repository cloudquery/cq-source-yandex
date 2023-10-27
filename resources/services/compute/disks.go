package compute

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Disks() *schema.Table {
	return &schema.Table{
		Name:      "yandex_compute_disks",
		Resolver:  fetchDisks,
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
				Name:     "type_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TypeId"),
			},
			{
				Name:     "zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ZoneId"),
			},
			{
				Name:     "size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "block_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BlockSize"),
			},
			{
				Name:     "product_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ProductIds"),
			},
			{
				Name:     "status",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "instance_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("InstanceIds"),
			},
			{
				Name:     "disk_placement_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DiskPlacementPolicy"),
			},
		},
	}
}
