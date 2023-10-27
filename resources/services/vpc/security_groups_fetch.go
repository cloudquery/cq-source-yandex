package vpc

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func fetchSecurityGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.Services.VPC.SecurityGroup().SecurityGroupIterator(ctx, &vpc.ListSecurityGroupsRequest{FolderId: c.MultiplexedResourceId})
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
