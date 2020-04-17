package internal

import (
	"fmt"
	"io/ioutil"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
	"github.com/vektah/gqlparser/v2/validator"
)

func loadSchema(schemaFile string) (*ast.Schema, error) {
	dat, err := ioutil.ReadFile(schemaFile)

	if err != nil {
		return nil, fmt.Errorf("failed to read schema: %w", err)
	}

	schema, graphqlError := gqlparser.LoadSchema(&ast.Source{
		Name:  schemaFile,
		Input: string(dat),
	})

	if graphqlError != nil {
		return nil, graphqlError
	}

	return schema, nil
}

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
		query, err := processQuery(schema, file)

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

	for _, query := range queries {
		generateStruct(query)
	}

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

func generateStruct(query *ast.QueryDocument) error {

	return nil
}

func processQuery(schema *ast.Schema, queryFile string) (*ast.QueryDocument, error) {
	queryData, err := ioutil.ReadFile(queryFile)

	if err != nil {
		return nil, fmt.Errorf("failed to read file %s %w", queryFile, err)
	}

	source := &ast.Source{
		Input: string(queryData),
	}

	queryDoc, parseError := parser.ParseQuery(source)

	if parseError != nil {
		return nil, parseError
	}

	validationErrors := validator.Validate(schema, queryDoc)

	if validationErrors != nil {
		return nil, validationErrors
	}

	return queryDoc, nil
}
