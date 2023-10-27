package recipes

import (
	"github.com/cloudquery/plugin-sdk/v4/codegen"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

func ResourceManager() []*Resource {
	return []*Resource{
		{
			Service:      "resourcemanager",
			SubService:   "clouds",
			Struct:       new(resourcemanager.Cloud),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexCloud,
		},
		{
			Service:      "resourcemanager",
			SubService:   "folders",
			Struct:       new(resourcemanager.Folder),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
	}
}
