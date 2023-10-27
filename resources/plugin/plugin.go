package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

var (
	Name    = "yandex-cloud"
	Kind    = "source"
	Team    = "yandex"
	Version = "Development"
)

func Plugin() *plugin.P {
	return plugin.NewPlugin(
		Name,
		Version,
		configure,
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
	)
	allTables := Tables()
	// here you can append custom non-generated tables
	return plugin.NewSourcePlugin(
		"yandex",
		Version,
		allTables,
		client.Configure,
	)
}
