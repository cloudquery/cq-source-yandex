package recipes

import (
	"github.com/cloudquery/plugin-sdk/v4/codegen"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1"
)

func KMS() []*Resource {
	return []*Resource{
		{
			Service:      "kms",
			SubService:   "symmetric_keys",
			Struct:       new(kms.SymmetricKey),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
	}
}
