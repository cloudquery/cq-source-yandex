// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:      "yandex_compute_images",
		Resolver:  fetchImages,
		Multiplex: client.MultiplexBy(client.Folders),
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
				Name:     "family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Family"),
			},
			{
				Name:     "storage_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("StorageSize"),
			},
			{
				Name:     "min_disk_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinDiskSize"),
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
				Name:     "os",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Os"),
			},
			{
				Name:     "pooled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Pooled"),
			},
		},
	}
}
