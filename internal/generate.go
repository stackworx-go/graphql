package internal

import (
	"fmt"
	"os"
	"text/template"

	"github.com/stackworx-go/graphql-client/internal/generation"
	"github.com/vektah/gqlparser/v2/ast"
)

func loadClientTemplate() (*template.Template, error) {
	t := template.Must(template.New("client").
		Parse(string(MustAsset("template/client.tmpl"))))

	return t, nil
}

func loadRequestTemplate() (*template.Template, error) {
	t := template.Must(template.New("client").
		Parse(string(MustAsset("template/request.tmpl"))))

	return t, nil
}

// Generate Generate export
func Generate(queriesGlob, schemaFile, destination, packageName string) error {
	clientTemplate, err := loadClientTemplate()

	if err != nil {
		return err
	}

	requestTemplate, err := loadRequestTemplate()

	if err != nil {
		return err
	}

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

	f, err := os.OpenFile(destination, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("failed to open destination file: %s for writting: %w", destination, err)
	}

	defer f.Close()

	err = clientTemplate.Execute(f, struct {
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

		err = requestTemplate.Execute(f, t)

		if err != nil {
			return fmt.Errorf("failed execute request template: %w", err)
		}

		_, _ = f.WriteString("\n")
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
