package main

import (
	"log"

	"github.com/cloudquery/plugin-sdk/v4/serve"
	"github.com/yandex-cloud/cq-provider-yandex/resources/plugin"
)

func main() {
	p := serve.Plugin(plugin.Plugin())
	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
