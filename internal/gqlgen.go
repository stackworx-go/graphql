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

func loadQuery(schema *ast.Schema, queryFile string) (*ast.QueryDocument, error) {
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
