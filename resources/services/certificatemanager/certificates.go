package certificatemanager

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:      "yandex_certificatemanager_certificates",
		Resolver:  fetchCertificates,
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
				Name:     "type",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "domains",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Domains"),
			},
			{
				Name:     "status",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "issuer",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Issuer"),
			},
			{
				Name:     "subject",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject"),
			},
			{
				Name:     "serial",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Serial"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdatedAt"),
			},
			{
				Name:     "issued_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("IssuedAt"),
			},
			{
				Name:     "not_after",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("NotAfter"),
			},
			{
				Name:     "not_before",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("NotBefore"),
			},
			{
				Name:     "challenges",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Challenges"),
			},
			{
				Name:     "deletion_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DeletionProtection"),
			},
		},
	}
}
