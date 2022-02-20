package recruitment

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nagokos/connefut_backend/auth"
	"github.com/nagokos/connefut_backend/ent"
	"github.com/nagokos/connefut_backend/ent/recruitment"
	"github.com/nagokos/connefut_backend/graph/model"
	"github.com/nagokos/connefut_backend/logger"
)

type Recruitment struct {
	Title         string
	Type          model.Type
	Level         model.Level
	Place         *string
	StartAt       *time.Time
	Content       *string
	LocationLat   *float64
	LocationLng   *float64
	IsPublished   bool
	Capacity      *int
	ClosingAt     *time.Time
	CompetitionID *string
	PrefectureID  *string
}

func requiredIfUnnecessaryType() validation.RuleFunc {
	return func(v interface{}) error {
		if v == model.TypeOpponent {
			return errors.New("募集タイプを選択してください")
		}
		return nil
	}
}

func requiredIfUnnecessaryLevel() validation.RuleFunc {
	return func(v interface{}) error {
		if v == model.LevelUnnecessary {
			return errors.New("レベルを選択してください")
		}
		return nil
	}
}

func checkWithinTheDeadline(start time.Time) validation.RuleFunc {
	return func(v interface{}) error {
		var err error
		switch s := v.(type) {
		case *time.Time:
			difference := start.Sub(*s).Minutes()
			fmt.Println(difference)
			if difference < 60 {
				err = errors.New("募集期限は開催日時の1時間以上前に設定してください")
			} else {
				err = nil
			}
		}
		return err
	}
}

func (r Recruitment) CreateRecruitmentValidate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Title,
			validation.Required.Error("タイトルを入力してください"),
			validation.RuneLength(1, 60).Error("タイトルは1文字以上60文字以内で入力してください"),
		),
		validation.Field(
			&r.CompetitionID,
			validation.When(r.IsPublished,
				validation.Required.Error("募集競技を選択してください"),
			),
		),
		validation.Field(
			&r.Type,
			validation.In(
				model.TypeUnnecessary,
				model.TypeOpponent,
				model.TypeIndividual,
				model.TypeTeammate,
				model.TypeJoining,
				model.TypeCoaching,
				model.TypeOthers,
			),
			validation.When(r.IsPublished,
				validation.By(requiredIfUnnecessaryType()),
			),
		),
		validation.Field(
			&r.Content,
			validation.When(r.IsPublished,
				validation.Required.Error("募集の詳細を入力してください"),
				validation.RuneLength(1, 10000).Error("募集の詳細は10000文字以内で入力してください"),
			).Else(validation.RuneLength(0, 10000).Error("募集の詳細は10000文字以内で入力してください")),
		),
		validation.Field(
			&r.PrefectureID,
			validation.When(r.IsPublished,
				validation.Required.Error("募集エリアを選択してください"),
			),
		),
		validation.Field(
			&r.Place,
			validation.When(r.IsPublished,
				validation.When(r.Type == model.TypeOpponent || r.Type == model.TypeIndividual,
					validation.Required.Error("会場名を入力してください"),
				),
			),
		),
		validation.Field(
			&r.Level,
			validation.Required.Error("レベルを選択してください"),
			validation.In(
				model.LevelUnnecessary,
				model.LevelEnjoy,
				model.LevelBeginner,
				model.LevelMiddle,
				model.LevelExpert,
				model.LevelOpen,
			).Error("選択肢の中から選んでください"),
			validation.When(r.IsPublished,
				validation.When(
					r.Type == model.TypeOpponent ||
						r.Type == model.TypeIndividual ||
						r.Type == model.TypeTeammate ||
						r.Type == model.TypeJoining ||
						r.Type == model.TypeCoaching,
					validation.By(requiredIfUnnecessaryLevel()),
				),
			),
		),
		validation.Field(
			&r.Capacity,
			validation.When(r.IsPublished,
				validation.Required.Error("募集人数は1名以上にしてください"),
				validation.Min(1).Error("募集人数は1名以上にしてください"),
			),
		),
		validation.Field(
			&r.StartAt,
			validation.When(r.IsPublished,
				validation.When(r.Type == model.TypeOpponent || r.Type == model.TypeIndividual,
					validation.Required.Error("開催日時を設定してください"),
				),
			),
		),
		validation.Field(
			&r.ClosingAt,
			validation.When(r.IsPublished,
				validation.Required.Error("募集期限を設定してください"),
				validation.By(checkWithinTheDeadline(*r.StartAt)),
			),
		),
	)
}

func (r *Recruitment) CreateRecruitment(ctx context.Context, client *ent.RecruitmentClient) (*model.Recruitment, error) {
	currentUser := auth.ForContext(ctx)

	res, err := client.
		Create().
		SetTitle(r.Title).
		SetType(recruitment.Type(strings.ToLower(string(r.Type)))).
		SetLevel(recruitment.Level(strings.ToLower(string(r.Level)))).
		SetNillableCapacity(r.Capacity).
		SetNillableStartAt(r.StartAt).
		SetNillableContent(r.Content).
		SetNillablePlace(r.Place).
		SetNillableLocationLat(r.LocationLat).
		SetNillableLocationLng(r.LocationLng).
		SetIsPublished(r.IsPublished).
		SetNillableClosingAt(r.ClosingAt).
		SetNillableCompetitionID(r.CompetitionID).
		SetNillablePrefectureID(r.PrefectureID).
		SetUserID(currentUser.ID).
		Save(ctx)
	if err != nil {
		logger.Log.Error().Msg(fmt.Sprintln("recruitment create errors:", err.Error()))
		return &model.Recruitment{}, err
	}

	resRecruitment := &model.Recruitment{
		ID:          res.ID,
		Title:       res.Title,
		Type:        model.Type(res.Type),
		Level:       model.Level(res.Level),
		Place:       &res.Place,
		StartAt:     &res.StartAt,
		Content:     &res.Content,
		LocationLat: &res.LocationLat,
		LocationLng: &res.LocationLng,
		IsPublished: res.IsPublished,
		Capacity:    &res.Capacity,
		ClosingAt:   &res.ClosingAt,
		User:        currentUser,
	}

	return resRecruitment, nil
}
