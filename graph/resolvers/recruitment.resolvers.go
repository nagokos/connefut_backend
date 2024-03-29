package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nagokos/connefut_backend/graph/generated"
	"github.com/nagokos/connefut_backend/graph/loader"
	"github.com/nagokos/connefut_backend/graph/model"
	"github.com/nagokos/connefut_backend/graph/models/recruitment"
	"github.com/nagokos/connefut_backend/graph/models/search"
	"github.com/nagokos/connefut_backend/graph/models/stock"
	"github.com/nagokos/connefut_backend/graph/models/user"
	"github.com/nagokos/connefut_backend/graph/utils"
	"github.com/nagokos/connefut_backend/logger"
)

// CreateRecruitment is the resolver for the createRecruitment field.
func (r *mutationResolver) CreateRecruitment(ctx context.Context, input model.RecruitmentInput) (model.CreateRecruitmentResult, error) {
	i := recruitment.RecruitmentInput{
		Title:        input.Title,
		Type:         input.Type,
		Detail:       input.Detail,
		StartAt:      input.StartAt,
		Venue:        input.Venue,
		LocationLat:  input.LocationLat,
		LocationLng:  input.LocationLng,
		Status:       input.Status,
		ClosingAt:    input.ClosingAt,
		SportID:      utils.DecodeUniqueIDIdentifierOnly(input.SportID),
		PrefectureID: utils.DecodeUniqueIDIdentifierOnly(input.PrefectureID),
		TagIDs:       utils.DecodeUniqueIDs(input.TagIds),
	}

	if err := i.RecruitmentValidate(); err != nil {
		logger.NewLogger().Error(err.Error())
		errs := err.(validation.Errors)

		var result model.CreateRecruitmentInvalidInputErrors
		for field, errMessage := range errs {
			result.InvalidInputs = append(result.InvalidInputs, &model.CreateRecruitmentInvalidInputError{
				Field:   model.RecruitmentInvalidInputField(strings.ToUpper(field)),
				Message: errMessage.Error(),
			})
		}
		return result, nil
	}

	result, err := i.CreateRecruitment(ctx, r.dbPool)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateRecruitment is the resolver for the updateRecruitment field.
func (r *mutationResolver) UpdateRecruitment(ctx context.Context, id string, input model.RecruitmentInput) (model.UpdateRecruitmentResult, error) {
	i := recruitment.RecruitmentInput{
		Title:        input.Title,
		Type:         input.Type,
		Detail:       input.Detail,
		StartAt:      input.StartAt,
		Venue:        input.Venue,
		LocationLat:  input.LocationLat,
		LocationLng:  input.LocationLng,
		Status:       input.Status,
		ClosingAt:    input.ClosingAt,
		SportID:      utils.DecodeUniqueIDIdentifierOnly(input.SportID),
		PrefectureID: utils.DecodeUniqueIDIdentifierOnly(input.PrefectureID),
		TagIDs:       utils.DecodeUniqueIDs(input.TagIds),
	}

	if err := i.RecruitmentValidate(); err != nil {
		logger.NewLogger().Error(err.Error())
		errs := err.(validation.Errors)

		var result model.UpdateRecruitmentInvalidInputErrors
		for field, message := range errs {
			result.InvalidInputs = append(result.InvalidInputs, &model.UpdateRecruitmentInvalidInputError{
				Field:   model.RecruitmentInvalidInputField(strings.ToUpper(field)),
				Message: message.Error(),
			})
		}
		return result, nil
	}

	result, err := i.UpdateRecruitment(ctx, r.dbPool, utils.DecodeUniqueIDIdentifierOnly(id))
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}

	return result, nil
}

// DeleteRecruitment is the resolver for the deleteRecruitment field.
func (r *mutationResolver) DeleteRecruitment(ctx context.Context, id string) (*model.DeleteRecruitmentResult, error) {
	success, err := recruitment.DeleteRecruitment(ctx, r.dbPool, utils.DecodeUniqueIDIdentifierOnly(id))
	if err != nil {
		return nil, err
	}
	return success, nil
}

// Recruitments is the resolver for the recruitments field.
func (r *queryResolver) Recruitments(ctx context.Context, first *int, after *string) (*model.RecruitmentConnection, error) {
	sp, err := search.NewSearchParams(first, after, nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := recruitment.GetRecruitments(ctx, r.dbPool, sp)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ViewerRecruitments is the resolver for the viewerRecruitments field.
func (r *queryResolver) ViewerRecruitments(ctx context.Context, first *int, after *string) (*model.RecruitmentConnection, error) {
	params, err := search.NewSearchParams(first, after, nil, nil)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}

	res, err := recruitment.GetViewerRecruitments(ctx, r.dbPool, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Recruitment is the resolver for the recruitment field.
func (r *queryResolver) Recruitment(ctx context.Context, id string) (*model.Recruitment, error) {
	res, err := recruitment.GetRecruitment(ctx, r.dbPool, utils.DecodeUniqueIDIdentifierOnly(id))
	if err != nil {
		return res, err
	}
	return res, nil
}

// StockedRecruitments is the resolver for the stockedRecruitments field.
func (r *queryResolver) StockedRecruitments(ctx context.Context, first *int, after *string) (*model.RecruitmentConnection, error) {
	params, err := search.NewSearchParams(first, after, nil, nil)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}

	res, err := recruitment.GetStockedRecruitments(ctx, r.dbPool, params)
	if err != nil {
		return nil, err
	}

	return res, err
}

// AppliedRecruitments is the resolver for the appliedRecruitments field.
func (r *queryResolver) AppliedRecruitments(ctx context.Context) ([]*model.Recruitment, error) {
	res, err := recruitment.GetAppliedRecruitments(ctx, r.dbPool)
	if err != nil {
		return res, err
	}
	return res, err
}

// ID is the resolver for the id field.
func (r *recruitmentResolver) ID(ctx context.Context, obj *model.Recruitment) (string, error) {
	return utils.GenerateUniqueID("Recruitment", obj.DatabaseID), nil
}

// Type is the resolver for the type field.
func (r *recruitmentResolver) Type(ctx context.Context, obj *model.Recruitment) (model.Type, error) {
	return model.Type(strings.ToUpper(obj.Type.String())), nil
}

// Status is the resolver for the status field.
func (r *recruitmentResolver) Status(ctx context.Context, obj *model.Recruitment) (model.Status, error) {
	return model.Status(strings.ToUpper(obj.Status.String())), nil
}

// Sport is the resolver for the sport field.
func (r *recruitmentResolver) Sport(ctx context.Context, obj *model.Recruitment) (*model.Sport, error) {
	sport, err := loader.GetSport(ctx, obj.SportID)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}
	return sport, nil
}

// Prefecture is the resolver for the prefecture field.
func (r *recruitmentResolver) Prefecture(ctx context.Context, obj *model.Recruitment) (*model.Prefecture, error) {
	prefecture, err := loader.GetPrefecture(ctx, obj.PrefectureID)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}
	return prefecture, nil
}

// User is the resolver for the user field.
func (r *recruitmentResolver) User(ctx context.Context, obj *model.Recruitment) (*model.User, error) {
	user, err := loader.GetUser(ctx, obj.UserID)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}
	return user, nil
}

// Tags is the resolver for the tags field.
func (r *recruitmentResolver) Tags(ctx context.Context, obj *model.Recruitment) ([]*model.Tag, error) {
	tags, err := loader.LoadTagsByRecruitmentID(ctx, obj.DatabaseID)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}
	return tags, nil
}

// Applicant is the resolver for the applicant field.
func (r *recruitmentResolver) Applicant(ctx context.Context, obj *model.Recruitment) (*model.Applicant, error) {
	panic(fmt.Errorf("not implemented"))
}

// FeedbackStock is the resolver for the FeedbackStock field.
func (r *recruitmentResolver) FeedbackStock(ctx context.Context, obj *model.Recruitment) (*model.FeedbackStock, error) {
	viewer := user.GetViewer(ctx)
	//* ログインしているときは処理を進めない
	if viewer == nil {
		return &model.FeedbackStock{RecruitmentID: obj.DatabaseID}, nil
	}
	feedback, err := stock.GetFeedbackStock(ctx, r.dbPool, obj.DatabaseID)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}
	return feedback, nil
}

// Cursor is the resolver for the cursor field.
func (r *recruitmentEdgeResolver) Cursor(ctx context.Context, obj *model.RecruitmentEdge) (string, error) {
	return utils.GenerateUniqueID("Recruitment", obj.Node.DatabaseID), nil
}

// Recruitment returns generated.RecruitmentResolver implementation.
func (r *Resolver) Recruitment() generated.RecruitmentResolver { return &recruitmentResolver{r} }

// RecruitmentEdge returns generated.RecruitmentEdgeResolver implementation.
func (r *Resolver) RecruitmentEdge() generated.RecruitmentEdgeResolver {
	return &recruitmentEdgeResolver{r}
}

type recruitmentResolver struct{ *Resolver }
type recruitmentEdgeResolver struct{ *Resolver }
