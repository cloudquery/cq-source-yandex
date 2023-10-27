package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func MultiplexByOrg(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	res := make([]schema.ClientMeta, len(client.orgs))
	for i, id := range client.orgs {
		res[i] = client.withOrgID(id)
	}
	return res
}

func MultiplexByCloud(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	res := make([]schema.ClientMeta, len(client.clouds))
	for i, id := range client.clouds {
		res[i] = client.withCloudID(id)
	}
	return res
}

func MultiplexByFolder(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	res := make([]schema.ClientMeta, len(client.clouds))
	for i, id := range client.folders {
		res[i] = client.withFolderID(id)
	}
	return res
}
