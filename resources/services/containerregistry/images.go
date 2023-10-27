package containerregistry

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:      "yandex_containerregistry_images",
		Resolver:  fetchImages,
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
				Name:        "folder_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveMultiplexedResourceID,
				Description: `Folder ID`,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "digest",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Digest"),
			},
			{
				Name:     "compressed_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CompressedSize"),
			},
			{
				Name:     "config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Config"),
			},
			{
				Name:     "layers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Layers"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreatedAt"),
			},
		},
	}
}
