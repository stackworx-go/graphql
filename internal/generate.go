package internal

import (
	"fmt"

	"github.com/vektah/gqlparser/v2/ast"
)

// Generate Generate export
func Generate(queriesGlob, schemaFile string) error {
	files, err := Glob(queriesGlob)

	if err != nil {
		return fmt.Errorf("failed to find query files: %s", err)
	}

	schema, err := loadSchema(schemaFile)

	if err != nil {
		return fmt.Errorf("failed to load schema: %w", err)
	}

	var queries []*ast.QueryDocument
	operationNames := make(map[string]interface{})

	// TODO: handle no files found
	for _, file := range files {
		query, err := loadQuery(schema, file)

		if err != nil {
			return fmt.Errorf("Failed to process query %s: %w", file, err)
		}

		if err = validateQuery(query); err != nil {
			return fmt.Errorf("Failed to validate query %s: %w", file, err)
		}

		queries = append(queries, query)

		for _, operation := range query.Operations {
			if _, ok := operationNames[operation.Name]; ok {
				return fmt.Errorf("Operation %s used twice", operation.Name)
			} else {
				operationNames[operation.Name] = nil
			}
		}
	}

	var generatedQueries []Query

	for _, query := range queries {
		newQueries, err := generateStruct(query)
		if err != nil {
			return err
		}
		generatedQueries = append(generatedQueries, newQueries...)
	}

	fmt.Printf("%v", generatedQueries)

	return nil
}

func validateQuery(query *ast.QueryDocument) error {
	if len(query.Fragments) > 0 {
		return fmt.Errorf("Fragments are not allowed")
	}

	// TODO: ensure operation names match filename like relay

	// TODO: reject unnamed operations

	return nil
}
