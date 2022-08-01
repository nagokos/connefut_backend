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

// SendVerifyEmail is the resolver for the sendVerifyEmail field.
func (r *mutationResolver) SendVerifyEmail(ctx context.Context) (bool, error) {
	isSentEmail, err := user.SendVerifyEmail(ctx, r.dbPool)
	if err != nil {
		return false, err
	}
	return isSentEmail, nil
}

// SendVerifyNewEmail is the resolver for the sendVerifyNewEmail field.
func (r *mutationResolver) SendVerifyNewEmail(ctx context.Context, input model.SendVerifyNewEmailInput) (*model.SendVerifyNewEmailPayload, error) {
	u := user.User{
		Email: input.Email,
	}

	err := u.SendVerifyNewEmailValidate()
	if err != nil {
		logger.NewLogger().Error(err.Error())
		errs := err.(validation.Errors)

		var payload model.SendVerifyNewEmailPayload

		for k, errMessage := range errs {
			payload.UserErrors = append(payload.UserErrors, &model.SendVerifyNewEmailInvalidInputError{
				Message: errMessage.Error(),
				Field:   model.SendVerifyNewEmailInvalidInputField(strings.ToLower(k)),
			})
		}

		return &payload, nil
	}

	payload, err := u.SendVerifyNewEmail(ctx, r.dbPool)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (*model.ChangePasswordPayload, error) {
	i := user.ChangePasswordInput{
		CurrentPassword:         input.CurrentPassword,
		NewPassword:             input.NewPassword,
		NewPasswordConfirmation: input.NewPasswordConfirmation,
	}

	err := i.ChangePasswordValidate()
	if err != nil {
		logger.NewLogger().Error(err.Error())
		errs := err.(validation.Errors)

		var payload model.ChangePasswordPayload

		for k, errMessage := range errs {
			payload.UserErrors = append(payload.UserErrors, &model.ChangePasswordInvalidInputError{
				Message: errMessage.Error(),
				Field:   model.ChangePasswordInvalidInputField(strings.ToLower(k)),
			})
		}

		return &payload, nil
	}

	payload, err := i.ChangePassword(ctx, r.dbPool)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// VerifyEmail is the resolver for the verifyEmail field.
func (r *mutationResolver) VerifyEmail(ctx context.Context, input model.VerifyEmailInput) (*model.VerifyEmailPayload, error) {
	i := user.VerifyEmailInput{
		Code: input.Code,
	}
	if err := i.VerifyEmailValidate(); err != nil {
		logger.NewLogger().Error(err.Error())
		errs := err.(validation.Errors)
		var payload model.VerifyEmailPayload
		for k, errMessage := range errs {
			payload.UserErrors = append(payload.UserErrors, &model.VerifyEmailInvalidInputError{
				Message: errMessage.Error(),
				Field:   model.VerifyEmailInvalidInputField(strings.ToLower(k)),
			})
		}
		return &payload, nil
	}

	payload, err := i.VerifyEmail(ctx, r.dbPool)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// Viewer is the resolver for the viewer field.
func (r *queryResolver) Viewer(ctx context.Context) (*model.User, error) {
	user := user.GetViewer(ctx)
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

// EmailVerificationStatus is the resolver for the emailVerificationStatus field.
func (r *userResolver) EmailVerificationStatus(ctx context.Context, obj *model.User) (model.EmailVerificationStatus, error) {
	return model.EmailVerificationStatus(strings.ToUpper(obj.EmailVerificationStatus.String())), nil
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
