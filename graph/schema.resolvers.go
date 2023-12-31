package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating, and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"github.com/punkestu/hello-graphql/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{
		ID:   "todo2",
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "The User",
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	dummyTodo := model.Todo{
		ID:   "todo1",
		Text: "test todo",
		Done: false,
		User: &model.User{
			ID:   "user1",
			Name: "admin",
		},
	}
	todos = append(todos, &dummyTodo)
	return todos, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	dummyUser := model.User{
		ID:   "user3",
		Name: "User 3",
	}
	users = append(users, &dummyUser)
	return users, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
