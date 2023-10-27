package recipes

import (
	"github.com/cloudquery/plugin-sdk/v4/codegen"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	id = "Id"
)

var (
	idCol = codegen.ColumnDefinition{
		Name:        "id",
		Type:        schema.TypeString,
		Resolver:    `schema.PathResolver("Id")`,
		Description: `Resource ID`,
		Options:     schema.ColumnCreationOptions{PrimaryKey: true},
	}
)
