package graphql

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/stackworx-go/graphql/internal"
	"github.com/stackworx-go/graphql/internal/generation"
	"github.com/vektah/gqlparser/v2/ast"
)

func loadClientTemplate() *template.Template {

	t := template.Must(template.New("client").
		Parse(string(internal.MustAsset("template/client.tmpl"))),
	)

	template.Must(t.Parse(string(internal.MustAsset("template/struct.tmpl"))))
	return t
}

func loadRequestTemplate() *template.Template {
	t := template.Must(template.New("request").
		Parse(string(internal.MustAsset("template/request.tmpl"))),
	)

	template.Must(t.Parse(string(internal.MustAsset("template/struct.tmpl"))))
	return t
}

// Generate Generate export
func Generate(cfg *internal.Config) error {
	schema, err := internal.LoadSchema(cfg.SchemaPath)

	if err != nil {
		return fmt.Errorf("failed to load schema: %w", err)
	}

	return GenerateWithSchema(cfg.Queries, cfg.DestinationPath, cfg.PackageName, cfg.Scalar.Upload, schema)
}

// GenerateWithSchema GenerateWithSchema export
func GenerateWithSchema(queriesGlob []string, destination, packageName, scalarUpload string, schema *ast.Schema) error {
	clientTemplate := loadClientTemplate()
	requestTemplate := loadRequestTemplate()

	if len(queriesGlob) > 1 {
		return fmt.Errorf("only single queries matcher currently supported")
	}

	// TODO: handle list
	files, err := internal.Glob(queriesGlob[0])

	if err != nil {
		return fmt.Errorf("failed to find query files: %s", err)
	}

	var queries []*ast.QueryDocument
	operationNames := make(map[string]interface{})

	if len(files) == 0 {
		return fmt.Errorf("no queries found. Glob: %s", queriesGlob)
	}

	for _, file := range files {
		query, err := internal.LoadQuery(schema, file)

		if err != nil {
			return fmt.Errorf("failed to process query %s: %w", file, err)
		}

		if len(query.Operations) > 1 {
			return fmt.Errorf("each query should contain a single operation. File: %s", file)
		}

		if err = validateQuery(file, query); err != nil {
			return fmt.Errorf("failed to validate query %s: %w", file, err)
		}

		queries = append(queries, query)

		for _, operation := range query.Operations {
			if _, ok := operationNames[operation.Name]; ok {
				return fmt.Errorf("operation Name %s is not unique", operation.Name)
			}

			// Check for file uploads
			//if operation.Operation == ast.Mutation {
			//	operation.VariableDefinitions
			//}

			operationNames[operation.Name] = nil
		}
	}

	var generatedQueries []generation.Query

	for _, query := range queries {
		newQueries, err := generation.GenerateQueries(query, schema, scalarUpload)
		if err != nil {
			return err
		}
		generatedQueries = append(generatedQueries, newQueries...)

	}

	inputStructs, err := generation.GenerateInputStructs(schema)

	if err != nil {
		return fmt.Errorf("failed to generate input structs: %w", err)
	}

	var buf bytes.Buffer

	err = clientTemplate.Execute(&buf, struct {
		PackageName      string
		InputStructs     generation.Structs
		ScalarFileUpload string
	}{
		PackageName:      packageName,
		InputStructs:     inputStructs,
		ScalarFileUpload: scalarUpload,
	})

	if err != nil {
		return fmt.Errorf("failed execute client template: %w", err)
	}

	for _, q := range generatedQueries {
		err = requestTemplate.Execute(&buf, q)

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
		return fmt.Errorf("fragments are not allowed")
	}

	for _, op := range query.Operations {
		if op.Name == "" {
			return fmt.Errorf("anonymous query found in %s", file)
		}

		switch op.Operation {
		case ast.Query:
			if !strings.HasSuffix(op.Name, "Query") {
				return fmt.Errorf("expected Query Operation Name to end in Query, got: %s, file: %s", op.Name, file)
			}
		case ast.Mutation:
			if !strings.HasSuffix(op.Name, "Mutation") {
				return fmt.Errorf("expected Query Operation Name to end in Query, got: %s, file: %s", op.Name, file)
			}
		case ast.Subscription:
			return fmt.Errorf("subscriptions are currently not supported %s in %s", op.Name, file)
		}

		expectedName := op.Name + ".graphql"
		_, filename := filepath.Split(file)

		if expectedName != filename {
			return fmt.Errorf("expected filename to be %s, got: %s", expectedName, file)
		}

	}

	return nil
}
