package resolver

import (
	"context"
	"fmt"
	"meetup-graphql/core"
	"meetup-graphql/graphql/exec"

	"github.com/google/uuid"
)

func (r *mutationResolver) AddUser(ctx context.Context, input *core.NewUser) (*core.User, error) {
	user := core.User{
		UserId:   uuid.New(),
		Username: input.Username,
		Email:    input.Email,
	}

	r.users = append(r.users, user)
	return &user, nil
}

func (r *queryResolver) Meetups(ctx context.Context) ([]core.Meetup, error) {
	panic(fmt.Errorf("not implemented: Meetups - meetups"))
}

func (r *queryResolver) Users(ctx context.Context) ([]core.User, error) {
	return r.users, nil
}

func (r *userResolver) UserID(ctx context.Context, obj *core.User) (string, error) {
	return obj.UserId.String(), nil
}

func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() exec.QueryResolver { return &queryResolver{r} }

func (r *Resolver) User() exec.UserResolver { return &userResolver{r} }

type mutationResolver struct {
	*Resolver
}
type queryResolver struct {
	*Resolver
}
type userResolver struct {
	*Resolver
}
