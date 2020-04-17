package main

import (
	"log"
	"os"

	"github.com/stackworx-go/gqlgen-relay/internal"
	"github.com/urfave/cli/v2"
)

type Config struct {
	QueryFilenames string
}

func main() {
	var queries string
	var schemaFile string

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
		},
		Action: func(c *cli.Context) error {
			return internal.Generate(queries, schemaFile)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
