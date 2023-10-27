package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestHelper(t *testing.T, table *schema.Table, createService func() (*Services, error), options TestOptions) {
	version := "vDev"
	t.Helper()

	table.IgnoreInTests = false

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger) (schema.ClientMeta, error) {
		svc, err := createService()
		if err != nil {
			return nil, fmt.Errorf("failed to createService: %w", err)
		}
		var ycSpec Spec
		if err := spec.UnmarshalSpec(&ycSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal yc spec: %w", err)
		}
		c := NewYandexClient(logger, nil, nil,
			[]string{"test-folder-id"},
			[]string{"test-cloud-id"},
			[]string{"test-organization-id"},
			svc,
		)

		return c, nil
	}
	p := plugin.NewSourcePlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	plugin.TestSourcePluginSync(t, p)
}
