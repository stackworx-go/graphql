package graphqlclient

//go:generate go run github.com/99designs/gqlgen generate

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/stackworx-go/graphql-client/internal"
	"github.com/stackworx-go/graphql-client/internal/generation"
	"github.com/vektah/gqlparser/v2/ast"
)

func loadClientTemplate() (*template.Template, error) {
	t := template.Must(template.New("client").
		Parse(string(internal.MustAsset("template/client.tmpl"))))

	return t, nil
}

func loadRequestTemplate() (*template.Template, error) {
	t := template.Must(template.New("client").
		Parse(string(internal.MustAsset("template/request.tmpl"))))

	return t, nil
}

// Generate Generate export
func Generate(queriesGlob, schemaFile, destination, packageName string) error {
	schema, err := internal.LoadSchema(schemaFile)

	if err != nil {
		return fmt.Errorf("failed to load schema: %w", err)
	}

	return GenerateWithSchema(queriesGlob, destination, packageName, schema)
}

// GenerateWithSchema GenerateWithSchema export
func GenerateWithSchema(queriesGlob, destination, packageName string, schema *ast.Schema) error {
	clientTemplate, err := loadClientTemplate()

	if err != nil {
		return err
	}

	requestTemplate, err := loadRequestTemplate()

	if err != nil {
		return err
	}

	files, err := internal.Glob(queriesGlob)

	if err != nil {
		return fmt.Errorf("failed to find query files: %s", err)
	}

	var queries []*ast.QueryDocument
	operationNames := make(map[string]interface{})

	if len(files) == 0 {
		return fmt.Errorf("No queries found. Glob: %s", queriesGlob)
	}

	for _, file := range files {
		query, err := internal.LoadQuery(schema, file)

		if err != nil {
			return fmt.Errorf("Failed to process query %s: %w", file, err)
		}

		if len(query.Operations) > 1 {
			return fmt.Errorf("Each query should contain a single operation. File: %s", file)
		}

		if err = validateQuery(file, query); err != nil {
			return fmt.Errorf("Failed to validate query %s: %w", file, err)
		}

		queries = append(queries, query)

		for _, operation := range query.Operations {
			if _, ok := operationNames[operation.Name]; ok {
				return fmt.Errorf("Operation Name %s is not unique", operation.Name)
			}

			operationNames[operation.Name] = nil
		}
	}

	var generatedQueries []generation.Query

	for _, query := range queries {
		newQueries, err := generation.GenerateStruct(query)
		if err != nil {
			return err
		}
		generatedQueries = append(generatedQueries, newQueries...)
	}

	inputStructs, err := generation.GenerateInputStructs(schema)

	if err != nil {
		return fmt.Errorf("Failed to generate input structs: %w", err)
	}

	var buf bytes.Buffer

	err = clientTemplate.Execute(&buf, struct {
		PackageName  string
		InputStructs string
	}{
		PackageName:  packageName,
		InputStructs: inputStructs.Print(),
	})

	if err != nil {
		return fmt.Errorf("failed execute client template: %w", err)
	}

	for _, q := range generatedQueries {
		type tmpl struct {
			Name      string
			Arguments generation.Arguments
			HasInput  bool
			Payload   string
			Query     string
		}

		t := tmpl{
			Name:    q.Name,
			Query:   q.Query,
			Payload: q.Payload.Print(),
		}

		if q.Arguments != nil {
			t.Arguments = q.Arguments
			t.HasInput = true
		}

		err = requestTemplate.Execute(&buf, t)

		if err != nil {
			return fmt.Errorf("failed execute request template: %w", err)
		}
	}

	data, err := format.Source(buf.Bytes())

	if err != nil {
		return fmt.Errorf("failed to format code: %w", err)
	}

	f, err := os.OpenFile(destination, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("failed to open destination file: %s for writting: %w", destination, err)
	}

	defer f.Close()
	_, err = f.Write(data)

	if err != nil {
		return fmt.Errorf("failed to write client file: %w", err)
	}

	return nil
}

func validateQuery(file string, query *ast.QueryDocument) error {
	if len(query.Fragments) > 0 {
		return fmt.Errorf("Fragments are not allowed")
	}

	for _, op := range query.Operations {
		if op.Name == "" {
			return fmt.Errorf("Anonymous query found in %s", file)
		}

		switch op.Operation {
		case ast.Query:
			if !strings.HasSuffix(op.Name, "Query") {
				return fmt.Errorf("Expected Query Operation Name to end in Query, got: %s, file: %s", op.Name, file)
			}
		case ast.Mutation:
			if !strings.HasSuffix(op.Name, "Mutation") {
				return fmt.Errorf("Expected Query Operation Name to end in Query, got: %s, file: %s", op.Name, file)
			}
		case ast.Subscription:
			return fmt.Errorf("Subscriptions are currently not supported %s in %s", op.Name, file)
		}

		expectedName := op.Name + ".graphql"
		_, filename := filepath.Split(file)

		if expectedName != filename {
			return fmt.Errorf("Expected filename to be %s, got: %s", expectedName, file)
		}

	}

	return nil
}
