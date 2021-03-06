package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"io/ioutil"

	"github.com/stackworx-go/graphql/internal/integration/graph/generated"
	"github.com/stackworx-go/graphql/internal/integration/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.CreateTodoPayload, error) {
	r.Resolver.CreateTodo.Args = input
	return r.Resolver.CreateTodo.Data, r.Resolver.CreateTodo.Error
}

func (r *mutationResolver) UploadFile(ctx context.Context, input model.UploadFileInput) (*model.UploadFilePayload, error) {
	data, err := ioutil.ReadAll(input.File.File)

	if err != nil {
		return nil, err
	}

	return &model.UploadFilePayload{
		ID: input.ID,
		File: &model.File{
			Name:    input.File.Filename,
			Content: string(data),
		},
	}, nil
}

func (r *mutationResolver) UploadFiles(ctx context.Context, input model.UploadFilesInput) (*model.UploadFilesPayload, error) {
	files := make([]*model.File, 0)

	for _, f := range input.Files {
		data, err := ioutil.ReadAll(f.File)

		if err != nil {
			return nil, err
		}

		files = append(files, &model.File{
			Name:    f.Filename,
			Content: string(data),
		})
	}

	return &model.UploadFilesPayload{
		Files: files,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context, userID *string) ([]*model.Todo, error) {
	r.Resolver.Todos.Args = TodosArgs{
		UserID: userID,
	}
	return r.Resolver.Todos.Data, r.Resolver.Todos.Error
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	r.Resolver.Node.Args = id
	return r.Resolver.Node.Data, r.Resolver.Node.Error
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
