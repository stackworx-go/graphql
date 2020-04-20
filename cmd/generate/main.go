package main

import (
	"log"
	"os"

	"github.com/stackworx-go/graphql-client/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	var queries string
	var schemaFile string
	var destination string
	var packageName string

	app := &cli.App{
		Name:  "generate",
		Usage: "Generate structs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "queries",
				Value:       "integration/**/*.graphql",
				Usage:       "golang glob for queries",
				Required:    true,
				Destination: &queries,
			},
			&cli.StringFlag{
				Name:        "schema",
				Value:       "schema.graphqls",
				Usage:       "golang glob for queries",
				Required:    true,
				Destination: &schemaFile,
			},
			&cli.StringFlag{
				Name:        "destination",
				Value:       "client.go",
				Usage:       "Destination for generated client",
				Required:    true,
				Destination: &destination,
			},
			&cli.StringFlag{
				Name:        "packageName",
				Usage:       "Client Package Name",
				Required:    true,
				Destination: &packageName,
			},
		},
		Action: func(c *cli.Context) error {
			return internal.Generate(queries, schemaFile, destination, packageName)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
