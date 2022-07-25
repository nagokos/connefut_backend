package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nagokos/connefut_backend/auth"
	"github.com/nagokos/connefut_backend/graph/generated"
	"github.com/nagokos/connefut_backend/graph/model"
	"github.com/nagokos/connefut_backend/graph/models/recruitment"
	"github.com/nagokos/connefut_backend/graph/models/relationship"
	"github.com/nagokos/connefut_backend/graph/models/search"
	"github.com/nagokos/connefut_backend/graph/models/user"
	"github.com/nagokos/connefut_backend/graph/utils"
	"github.com/nagokos/connefut_backend/logger"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUserInput) (*model.RegisterUserPayload, error) {
	u := user.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	err := u.CreateUserValidate()

	if err != nil {
		logger.NewLogger().Error(err.Error())
		errs := err.(validation.Errors)

		var payload model.RegisterUserPayload

		for k, errMessage := range errs {
			payload.UserErrors = append(payload.UserErrors, &model.RegisterUserInvalidInputError{
				Message: errMessage.Error(),
				Field:   model.RegisterUserInvalidInputField(strings.ToLower(k)),
			})
		}

		return &payload, nil
	}

	payload, err := u.RegisterUser(ctx, r.dbPool)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return payload, nil
	}

	token, _ := user.CreateToken(payload.Viewer.DatabaseID)
	auth.SetAuthCookie(ctx, token)

	return payload, nil
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (*model.LoginUserPayload, error) {
	u := user.User{
		Email:    input.Email,
		Password: input.Password,
	}

	err := u.AuthenticateUserValidate()
	if err != nil {
		logger.NewLogger().Error(err.Error())
		errs := err.(validation.Errors)

		var payload model.LoginUserPayload

		for k, errMessage := range errs {
			payload.UserErrors = append(payload.UserErrors, model.LoginUserInvalidInputError{
				Message: errMessage.Error(),
				Field:   model.LoginUserInvalidInputField(strings.ToLower(k)),
			})
		}

		return &payload, nil
	}

	payload, err := u.LoginUser(ctx, r.dbPool)
	if err != nil {
		return payload, nil
	}

	token, _ := user.CreateToken(payload.Viewer.DatabaseID)
	auth.SetAuthCookie(ctx, token)

	return payload, nil
}

// LogoutUser is the resolver for the logoutUser field.
func (r *mutationResolver) LogoutUser(ctx context.Context) (bool, error) {
	auth.RemoveAuthCookie(ctx)
	return true, nil
}

// Viewer is the resolver for the viewer field.
func (r *queryResolver) Viewer(ctx context.Context) (*model.User, error) {
	user := auth.ForContext(ctx)
	return user, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := user.GetUser(ctx, r.dbPool, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return utils.GenerateUniqueID("User", obj.DatabaseID), nil
}

// Recruitments is the resolver for the recruitments field.
func (r *userResolver) Recruitments(ctx context.Context, obj *model.User, first *int, after *string) (*model.RecruitmentConnection, error) {
	params, err := search.NewSearchParams(first, after, nil, nil)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}
	connection, err := recruitment.GetUserRecruitments(ctx, r.dbPool, obj.DatabaseID, params)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}
	return connection, err
}

// Followings is the resolver for the followings field.
func (r *userResolver) Followings(ctx context.Context, obj *model.User, first *int, after *string) (*model.FollowConnection, error) {
	params, err := search.NewSearchParams(first, after, nil, nil)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}
	connection, err := relationship.GetFollowings(ctx, r.dbPool, obj.DatabaseID, params)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}
	return connection, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
