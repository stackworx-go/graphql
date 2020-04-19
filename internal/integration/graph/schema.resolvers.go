package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/stackworx-go/gqlgen-relay/internal/integration/graph/generated"
	"github.com/stackworx-go/gqlgen-relay/internal/integration/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return r.Resolver.CreateTodo.Data, r.Resolver.CreateTodo.Error
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.Resolver.Todos.Data, r.Resolver.Todos.Error
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
