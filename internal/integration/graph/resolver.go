package graph

import "github.com/stackworx-go/gqlgen-relay/internal/integration/graph/model"

type TodosArgs struct {
	UserID *string
}

type Resolver struct {
	Todos struct {
		Data  []*model.Todo
		Error error
		Args  TodosArgs
	}
	CreateTodo struct {
		Data  *model.CreateTodoPayload
		Error error
		Args  model.CreateTodoInput
	}
}
