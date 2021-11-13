package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cpunion/graph/app/graph/generated"
	gmodel "github.com/cpunion/graph/app/graph/model"
	"github.com/icza/gox/gox"
)

func (r *mutationResolver) SignUp(ctx context.Context, input gmodel.SignupInput) (*gmodel.SignPayload, error) {
	user := gmodel.User{
		Email: "a@b.com",
	}
	return &gmodel.SignPayload{User: &user}, nil
}

func (r *mutationResolver) SignIn(ctx context.Context, input gmodel.SigninInput) (*gmodel.SignPayload, error) {
	user := gmodel.User{
		Email: "a@b.com",
	}
	return &gmodel.SignPayload{User: &user}, nil
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*gmodel.User, error) {
	result := gmodel.User{
		Email: "a@b.com",
	}
	return &result, nil
}

func (r *userResolver) Servers(ctx context.Context, obj *gmodel.User) ([]*gmodel.Server, error) {
	return []*gmodel.Server{
		{
			ID:   "1",
			Name: gox.NewString("Server 1"),
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
