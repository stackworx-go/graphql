package main

import (
	"fmt"
	"os"

	"github.com/stackworx-go/graphql-client/internal/integration/graph/generated"
	"github.com/vektah/gqlparser/v2/formatter"
)

func main() {
	es := generated.NewExecutableSchema(generated.Config{})

	f, err := os.OpenFile("./schema.graphql", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	formatter := formatter.NewFormatter(f)
	formatter.FormatSchema(es.Schema())
}
