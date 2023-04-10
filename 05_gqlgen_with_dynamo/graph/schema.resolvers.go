package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"github.com/lambda-gqgen-presentation/graph/model"
	"github.com/lambda-gqgen-presentation/models"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	return r.todoService.AddTodo(models.Todo{
		Text:   input.Text,
		UserID: input.UserID,
		Done:   false,
	})
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, userID string) ([]*models.Todo, error) {
	return r.todoService.GetTodosByUserId(userID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
