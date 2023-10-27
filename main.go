package main

import (
	"github.com/cloudquery/plugin-sdk/v4/serve"
	"github.com/yandex-cloud/cq-provider-yandex/resources/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
