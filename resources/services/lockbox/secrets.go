package lockbox

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Secrets() *schema.Table {
	return &schema.Table{
		Name:      "yandex_lockbox_secrets",
		Resolver:  fetchSecrets,
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
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "status",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "current_version",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CurrentVersion"),
			},
			{
				Name:     "deletion_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DeletionProtection"),
			},
		},
	}
}
