package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

func GetFields(ctx context.Context) []string {
	var fieldContext = graphql.GetFieldContext(ctx)
	var fields []string
	for _, selection := range fieldContext.Field.SelectionSet {
		field, ok := selection.(*ast.Field)
		if ok {
			fields = append(fields, field.Name)
		}
	}
	return fields
}
