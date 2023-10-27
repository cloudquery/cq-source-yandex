package kms

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func SymmetricKeys() *schema.Table {
	return &schema.Table{
		Name:      "yandex_kms_symmetric_keys",
		Resolver:  fetchSymmetricKeys,
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
				Name:     "primary_version",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrimaryVersion"),
			},
			{
				Name:     "default_algorithm",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DefaultAlgorithm"),
			},
			{
				Name:     "rotated_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("RotatedAt"),
			},
			{
				Name:     "rotation_period",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RotationPeriod"),
			},
			{
				Name:     "deletion_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DeletionProtection"),
			},
		},
	}
}
