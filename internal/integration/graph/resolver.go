package graph

import "github.com/stackworx-go/graphql-client/internal/integration/graph/model"

// TodosArgs TodosArgs export
type TodosArgs struct {
	UserID *string
}

// Resolver Resolver export
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
	Node struct {
		Data  model.Node
		Error error
		Args  string
	}
}
