package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/VitorEmanoel/books-loan/common"
	"github.com/vektah/gqlparser/v2/ast"
)

func GetFields(ctx context.Context, allowed ...string) []string {
	var fieldContext = graphql.GetFieldContext(ctx)
	var fields []string
	for _, selection := range fieldContext.Field.SelectionSet {
		field, ok := selection.(*ast.Field)
		if ok {
			for _, allow := range allowed {
				if field.Name == allow {
					fields = append(fields, common.ToSnakeCase(field.Name))
					break
				}
			}
		}
	}
	return fields
}
