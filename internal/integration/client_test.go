package integration

import (
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stackworx-go/graphql-client/internal/integration/graph"
	"github.com/stackworx-go/graphql-client/internal/integration/graph/generated"
	"github.com/stackworx-go/graphql-client/internal/integration/graph/model"

	"github.com/stretchr/testify/assert"
)

func String(s string) *string {
	return &s
}

func TestTodosQuery(t *testing.T) {
	// given
	resolvers := &graph.Resolver{}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
	resolvers.Todos.Data = []*model.Todo{{
		ID:   "1",
		Text: "Buy Groceries",
		Done: false,
		User: &model.User{
			ID:   "1",
			Name: "John",
		},
	}}

	ts := httptest.NewServer(srv)
	defer ts.Close()

	client := Client{
		Url: ts.URL,
	}

	// when
	data, err := client.TodosQuery()

	// then
	assert.NoError(t, err)
	assert.Equal(t, resolvers.Todos.Args, graph.TodosArgs{UserID: nil})
	assert.Equal(t, data, &TodosQueryPayload{
		Todos: []TodosQueryPayloadTodos{
			{
				Id:   "1",
				Text: "Buy Groceries",
				User: TodosQueryPayloadTodosUser{
					Id:   "1",
					Name: "John",
				},
			},
		},
	})
}

func TestTodosQueryWithVariables(t *testing.T) {
	// given
	resolvers := &graph.Resolver{}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
	resolvers.Todos.Data = []*model.Todo{{
		ID:   "1",
		Text: "Buy Groceries",
		Done: false,
		User: &model.User{
			ID:   "1",
			Name: "John",
		},
	}}

	ts := httptest.NewServer(srv)
	defer ts.Close()

	client := Client{
		Url: ts.URL,
	}

	// when
	data, err := client.TodosQueryWithVariables("4")

	// then
	assert.NoError(t, err)
	assert.Equal(t, resolvers.Todos.Args, graph.TodosArgs{UserID: String("4")})
	assert.Equal(t, data, &TodosQueryWithVariablesPayload{
		Todos: []TodosQueryWithVariablesPayloadTodos{
			{
				Id:   "1",
				Text: "Buy Groceries",
				User: TodosQueryWithVariablesPayloadTodosUser{
					Id:   "1",
					Name: "John",
				},
			},
		},
	})
}

func TestCreateTodoMutation(t *testing.T) {
	// given
	resolvers := &graph.Resolver{}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
	resolvers.CreateTodo.Data = &model.CreateTodoPayload{
		Todo: &model.Todo{
			ID:   "1",
			Text: "Buy Groceries",
			Done: false,
			User: &model.User{
				ID:   "1",
				Name: "John",
			},
		},
	}

	ts := httptest.NewServer(srv)
	defer ts.Close()

	client := Client{
		Url: ts.URL,
	}

	// when
	data, err := client.CreateTodoMutation(CreateTodoInput{
		Text:   "Create TODO",
		UserId: "4",
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, resolvers.CreateTodo.Args, model.CreateTodoInput{UserID: "4", Text: "Create TODO"})
	assert.Equal(t, data, &CreateTodoMutationPayload{
		CreateTodo: CreateTodoMutationPayloadCreateTodo{
			Todo: CreateTodoMutationPayloadCreateTodoTodo{
				Id:   "1",
				Text: "Buy Groceries",
				User: CreateTodoMutationPayloadCreateTodoTodoUser{
					Id:   "1",
					Name: "John",
				},
			},
		},
	})
}

func TestNodeUserQuery(t *testing.T) {
	// given
	resolvers := &graph.Resolver{}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
	resolvers.Node.Data = &model.User{
		ID:   "1",
		Name: "John",
	}

	ts := httptest.NewServer(srv)
	defer ts.Close()

	client := Client{
		Url: ts.URL,
	}

	// when
	data, err := client.NodeQuery("1")

	// then
	assert.NoError(t, err)
	// assert.Equal(t, resolvers.Node.Args, "1")
	assert.Equal(t, data, &NodeQueryPayload{
		Node: &NodeQueryPayloadNode{
			// Id: "1",
			UserFragment: &NodeQueryPayloadNodeUserFragment{
				Name: "John",
			},
		},
	})
}

func TestNodeTodoQuery(t *testing.T) {
	// given
	resolvers := &graph.Resolver{}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
	resolvers.Node.Data = &model.Todo{
		ID:   "1",
		Text: "TODO",
	}

	ts := httptest.NewServer(srv)
	defer ts.Close()

	client := Client{
		Url: ts.URL,
	}

	// when
	data, err := client.NodeQuery("1")

	// then
	assert.NoError(t, err)
	// assert.Equal(t, resolvers.Node.Args, "1")
	assert.Equal(t, data, &NodeQueryPayload{
		Node: &NodeQueryPayloadNode{
			// Id: "1",
			TodoFragment: &NodeQueryPayloadNodeTodoFragment{
				Text: "TODO",
			},
		},
	})
}
