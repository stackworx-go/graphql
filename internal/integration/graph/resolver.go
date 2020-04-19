package graph

import "github.com/stackworx-go/gqlgen-relay/internal/integration/graph/model"

type Resolver struct {
	Todos struct {
		Data  []*model.Todo
		Error error
	}
	CreateTodo struct {
		Data  *model.Todo
		Error error
	}
}
