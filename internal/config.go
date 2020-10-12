package internal

type ScalarMapping struct {
	PackageName string `yaml:"packageName"`
	Name        string `yaml:"name"`
}

type Config struct {
	Queries         []string `yaml:"queries"`
	SchemaPath      string   `yaml:"schema"`
	DestinationPath string   `yaml:"destination"`
	PackageName     string   `yaml:"packageName"`

	ScalarMappings map[string]ScalarMapping `yaml:"scalars,omitempty"`
}

// DefaultConfig creates a copy of the default config
func DefaultConfig() *Config {
	return &Config{
		Queries:         []string{"./**/*.graphql"},
		SchemaPath:      "schema.graphqls",
		DestinationPath: "./client.go",
		ScalarMappings: map[string]ScalarMapping{
			"Upload": {
				PackageName: "github.com/99designs/gqlgen/graphql",
				Name:        "Upload",
			},
		},
	}
}

// TODO: validate package is provided
