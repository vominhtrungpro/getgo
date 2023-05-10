package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kytruong0712/getgo/api/graph/model"
	"github.com/kytruong0712/getgo/api/internal/controller/users"
)

// UpsertCharacter is the resolver for the upsertCharacter field.
func (r *mutationResolver) UpsertCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	panic(fmt.Errorf("not implemented: UpsertCharacter - upsertCharacter"))
}

// Createuser is the resolver for the createuser field.
func (r *mutationResolver) Createuser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	var result model.User
	var request users.CreateInput
	request.Username = input.Username
	request.Password = input.Password
	request.Email = input.Email
	request.Password = input.Password
	request.Age = int64(input.Age)
	user, err := r.UserCtrl.Create(ctx, request)
	if err != nil {
		return &result, err
	}
	s := strconv.Itoa(int(user.ID))
	result.ID = s
	result.Externalid = user.ExternalID
	result.Username = user.Username
	result.Password = user.Password
	result.Email = user.Email
	result.Age = int(user.Age)
	return &result, nil
}

// Updateuser is the resolver for the updateuser field.
func (r *mutationResolver) Updateuser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	var result model.User
	var request users.UpdateInput
	id1, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return &result, err
	}
	request.ID = id1
	request.ExternalID = input.Externalid
	request.Username = input.Username
	request.Password = input.Password
	request.Email = input.Email
	request.Password = input.Password
	request.Age = int64(input.Age)
	errr := r.UserCtrl.Update(ctx, request)
	if errr != nil {
		return &result, errr
	}
	s := strconv.Itoa(int(request.ID))
	result.ID = s
	result.Externalid = request.ExternalID
	result.Username = request.Username
	result.Password = request.Password
	result.Email = request.Email
	result.Age = int(request.Age)
	return &result, nil
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	panic(fmt.Errorf("not implemented: Character - character"))
}

// Pogues is the resolver for the pogues field.
func (r *queryResolver) Pogues(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented: Pogues - pogues"))
}

// Kooks is the resolver for the kooks field.
func (r *queryResolver) Kooks(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented: Kooks - kooks"))
}

// Getalluser is the resolver for the getalluser field.
func (r *queryResolver) Getalluser(ctx context.Context) ([]*model.User, error) {
	var result []*model.User
	users, err := r.UserCtrl.GetAll(ctx)
	if err != nil {
		return result, err
	}
	for _, element := range users {
		var user model.User
		s := strconv.Itoa(int(element.ID))
		user.ID = s
		user.Externalid = element.ExternalID
		user.Username = element.Username
		user.Password = element.Password
		user.Email = element.Email
		user.Age = int(element.Age)
		result = append(result, &user)
	}
	return result, nil
}

// Getuserbyid is the resolver for the getuserbyid field.
func (r *queryResolver) Getuserbyid(ctx context.Context, id string) (*model.User, error) {
	var result model.User
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return &result, err
	}
	user, err := r.UserCtrl.GetById(ctx, id1)
	if err != nil {
		return &result, err
	}
	s := strconv.Itoa(int(user.ID))
	result.ID = s
	result.Externalid = user.ExternalID
	result.Username = user.Username
	result.Password = user.Password
	result.Email = user.Email
	result.Age = int(user.Age)
	return &result, nil
}

// Getuserbyusername is the resolver for the getuserbyusername field.
func (r *queryResolver) Getuserbyusername(ctx context.Context, username string) (*model.User, error) {
	var result model.User
	user, err := r.UserCtrl.GetByUsername(ctx, username)
	if err != nil {
		return &result, err
	}
	s := strconv.Itoa(int(user.ID))
	result.ID = s
	result.Externalid = user.ExternalID
	result.Username = user.Username
	result.Password = user.Password
	result.Email = user.Email
	result.Age = int(user.Age)
	return &result, nil
}

// Deleteuser is the resolver for the deleteuser field.
func (r *queryResolver) Deleteuser(ctx context.Context, id string) (*model.User, error) {
	var result *model.User
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return result, err
	}
	errr := r.UserCtrl.DeleteById(ctx, id1)
	return result, errr
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Createuser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	var result *model.User
	var request users.CreateInput
	request.Username = input.Username
	request.Password = input.Password
	request.Email = input.Email
	request.Password = input.Password
	request.Age = int64(input.Age)
	user, err := r.UserCtrl.Create(ctx, request)
	if err != nil {
		return result, err
	}
	s := strconv.Itoa(int(user.ID))
	result.ID = s
	result.Externalid = user.ExternalID
	result.Username = user.Username
	result.Password = user.Password
	result.Email = user.Email
	result.Age = int(user.Age)
	return result, nil
}
func (r *queryResolver) Updateuser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	var result *model.User
	var request users.UpdateInput
	request.Username = input.Username
	request.Password = input.Password
	request.Email = input.Email
	request.Password = input.Password
	request.Age = int64(input.Age)
	err := r.UserCtrl.Update(ctx, request)
	if err != nil {
		return result, err
	}
	return result, nil
}
