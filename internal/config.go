package internal

type ScalarMapping struct {
	PackageName string `yaml:"packageName"`
	Name        string `yaml:"name"`
}

type ScalarConfig struct {
	Upload   string                   `yaml:"upload"`
	Mappings map[string]ScalarMapping `yaml:"scalars,omitempty"`
}

type Config struct {
	Queries         []string `yaml:"queries"`
	SchemaPath      string   `yaml:"schema"`
	DestinationPath string   `yaml:"destination"`
	PackageName     string   `yaml:"packageName"`

	Scalar ScalarConfig `yaml:"scalar"`
}

// DefaultConfig creates a copy of the default config
func DefaultConfig() *Config {
	return &Config{
		Queries:         []string{"./**/*.graphql"},
		SchemaPath:      "schema.graphqls",
		DestinationPath: "./client.go",
		Scalar: ScalarConfig{
			Upload: "Upload",
			Mappings: map[string]ScalarMapping{
				"Upload": {
					PackageName: "github.com/99designs/gqlgen/graphql",
					Name:        "Upload",
				},
			},
		},
	}
}

// TODO: validate package is provided
