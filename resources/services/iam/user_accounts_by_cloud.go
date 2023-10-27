package iam

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func UserAccountsByCloud() *schema.Table {
	return &schema.Table{
		Name:      "yandex_iam_user_accounts_by_cloud",
		Resolver:  fetchUserAccountsByCloud,
		Multiplex: client.MultiplexByCloud,
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
				Name:     "cloud_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveMultiplexedResourceID,
			},
			{
				Name:     "saml_user_account",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamlUserAccount"),
			},
			{
				Name:     "yandex_passport_user_account",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("YandexPassportUserAccount"),
			},
		},
	}
}
