// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Competition struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Prefecture struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Recruitment struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Content     *string      `json:"content"`
	Type        Type         `json:"type"`
	Level       Level        `json:"level"`
	Place       *string      `json:"place"`
	StartAt     *time.Time   `json:"startAt"`
	LocationLat *float64     `json:"locationLat"`
	LocationLng *float64     `json:"locationLng"`
	IsPublished bool         `json:"isPublished"`
	Capacity    *int         `json:"capacity"`
	ClosingAt   *time.Time   `json:"closingAt"`
	Competition *Competition `json:"competition"`
	Prefecture  *Prefecture  `json:"prefecture"`
	User        *User        `json:"user"`
}

type User struct {
	ID                      string                  `json:"id"`
	Name                    string                  `json:"name"`
	Email                   string                  `json:"email"`
	Role                    Role                    `json:"role"`
	Avatar                  string                  `json:"avatar"`
	Introduction            *string                 `json:"introduction"`
	EmailVerificationStatus EmailVerificationStatus `json:"emailVerificationStatus"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RecruitmentInput struct {
	Title         string     `json:"title"`
	Content       *string    `json:"content"`
	Type          Type       `json:"type"`
	Level         Level      `json:"level"`
	Place         *string    `json:"place"`
	StartAt       *time.Time `json:"startAt"`
	LocationLat   *float64   `json:"locationLat"`
	LocationLng   *float64   `json:"locationLng"`
	Capacity      *int       `json:"capacity"`
	IsPublished   bool       `json:"isPublished"`
	ClosingAt     *time.Time `json:"closingAt"`
	CompetitionID *string    `json:"competitionId"`
	PrefectureID  *string    `json:"prefectureId"`
}

type EmailVerificationStatus string

const (
	EmailVerificationStatusUnnecessary EmailVerificationStatus = "UNNECESSARY"
	EmailVerificationStatusPending     EmailVerificationStatus = "PENDING"
	EmailVerificationStatusVerified    EmailVerificationStatus = "VERIFIED"
)

var AllEmailVerificationStatus = []EmailVerificationStatus{
	EmailVerificationStatusUnnecessary,
	EmailVerificationStatusPending,
	EmailVerificationStatusVerified,
}

func (e EmailVerificationStatus) IsValid() bool {
	switch e {
	case EmailVerificationStatusUnnecessary, EmailVerificationStatusPending, EmailVerificationStatusVerified:
		return true
	}
	return false
}

func (e EmailVerificationStatus) String() string {
	return string(e)
}

func (e *EmailVerificationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EmailVerificationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EmailVerificationStatus", str)
	}
	return nil
}

func (e EmailVerificationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Level string

const (
	LevelUnnecessary Level = "UNNECESSARY"
	LevelEnjoy       Level = "ENJOY"
	LevelBeginner    Level = "BEGINNER"
	LevelMiddle      Level = "MIDDLE"
	LevelExpert      Level = "EXPERT"
	LevelOpen        Level = "OPEN"
)

var AllLevel = []Level{
	LevelUnnecessary,
	LevelEnjoy,
	LevelBeginner,
	LevelMiddle,
	LevelExpert,
	LevelOpen,
}

func (e Level) IsValid() bool {
	switch e {
	case LevelUnnecessary, LevelEnjoy, LevelBeginner, LevelMiddle, LevelExpert, LevelOpen:
		return true
	}
	return false
}

func (e Level) String() string {
	return string(e)
}

func (e *Level) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Level(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Level", str)
	}
	return nil
}

func (e Level) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleAdmin   Role = "ADMIN"
	RoleGeneral Role = "GENERAL"
)

var AllRole = []Role{
	RoleAdmin,
	RoleGeneral,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleGeneral:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Type string

const (
	TypeUnnecessary Type = "UNNECESSARY"
	TypeOpponent    Type = "OPPONENT"
	TypeIndividual  Type = "INDIVIDUAL"
	TypeTeammate    Type = "TEAMMATE"
	TypeJoining     Type = "JOINING"
	TypeCoaching    Type = "COACHING"
	TypeOthers      Type = "OTHERS"
)

var AllType = []Type{
	TypeUnnecessary,
	TypeOpponent,
	TypeIndividual,
	TypeTeammate,
	TypeJoining,
	TypeCoaching,
	TypeOthers,
}

func (e Type) IsValid() bool {
	switch e {
	case TypeUnnecessary, TypeOpponent, TypeIndividual, TypeTeammate, TypeJoining, TypeCoaching, TypeOthers:
		return true
	}
	return false
}

func (e Type) String() string {
	return string(e)
}

func (e *Type) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Type(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}

func (e Type) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
