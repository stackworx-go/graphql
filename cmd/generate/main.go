package main

import (
	"fmt"
	"github.com/stackworx-go/graphql/internal"
	"io/ioutil"
	"log"
	"os"

	"github.com/stackworx-go/graphql"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

func main() {
	var configPath string

	app := &cli.App{
		Name:  "generate",
		Usage: "Generate structs",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "config",
				Required:    true,
				Usage:       "Yaml Configuration file path",
				Destination: &configPath,
			},
		},
		Action: func(c *cli.Context) error {
			cfg := internal.DefaultConfig()

			b, err := ioutil.ReadFile(c.Path("config"))
			if err != nil {
				return fmt.Errorf("unable to read config: %w", err)
			}

			if err := yaml.UnmarshalStrict(b, cfg); err != nil {
				return fmt.Errorf("unable to parse config: %w", err)
			}

			return graphql.Generate(cfg)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
