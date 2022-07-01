// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Connection interface {
	IsConnection()
}

type Edge interface {
	IsEdge()
}

type Error interface {
	IsError()
}

type Node interface {
	IsNode()
}

type UserLoginError interface {
	IsUserLoginError()
}

type Applicant struct {
	Message     string       `json:"message"`
	CreatedAt   time.Time    `json:"createdAt"`
	Recruitment *Recruitment `json:"recruitment"`
}

type Competition struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Competition) IsNode() {}

type Entrie struct {
	User *User `json:"user"`
}

type Message struct {
	Content   *string    `json:"content"`
	Applicant *Applicant `json:"applicant"`
	User      *User      `json:"user"`
	CreatedAt time.Time  `json:"createdAt"`
}

type PageInfo struct {
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
}

type Prefecture struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Prefecture) IsNode() {}

type Recruitment struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Detail      *string      `json:"detail"`
	Type        Type         `json:"type"`
	Place       *string      `json:"place"`
	StartAt     *time.Time   `json:"startAt"`
	LocationLat *float64     `json:"locationLat"`
	LocationLng *float64     `json:"locationLng"`
	Status      Status       `json:"status"`
	ClosingAt   *time.Time   `json:"closingAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	CreatedAt   time.Time    `json:"createdAt"`
	PublishedAt *time.Time   `json:"publishedAt"`
	Competition *Competition `json:"competition"`
	Prefecture  *Prefecture  `json:"prefecture"`
	User        *User        `json:"user"`
	Tags        []*Tag       `json:"tags"`
	Applicant   *Applicant   `json:"applicant"`
}

func (Recruitment) IsNode() {}

type RecruitmentConnection struct {
	PageInfo *PageInfo          `json:"pageInfo"`
	Edges    []*RecruitmentEdge `json:"edges"`
}

func (RecruitmentConnection) IsConnection() {}

type RecruitmentEdge struct {
	Cursor string       `json:"cursor"`
	Node   *Recruitment `json:"node"`
}

func (RecruitmentEdge) IsEdge() {}

type Room struct {
	ID     string  `json:"id"`
	Entrie *Entrie `json:"entrie"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID                      string                  `json:"id"`
	DatabaseID              *int                    `json:"databaseId"`
	Name                    string                  `json:"name"`
	Email                   string                  `json:"email"`
	Role                    Role                    `json:"role"`
	Avatar                  string                  `json:"avatar"`
	Introduction            *string                 `json:"introduction"`
	EmailVerificationStatus EmailVerificationStatus `json:"emailVerificationStatus"`
}

func (User) IsNode() {}

type UserLoginAuthenticationError struct {
	Message string `json:"message"`
}

func (UserLoginAuthenticationError) IsUserLoginError() {}
func (UserLoginAuthenticationError) IsError()          {}

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginInvalidInputError struct {
	Message string                     `json:"message"`
	Field   UserLoginInvalidInputField `json:"field"`
}

func (UserLoginInvalidInputError) IsUserLoginError() {}
func (UserLoginInvalidInputError) IsError()          {}

type UserLoginPayload struct {
	User       *User            `json:"user"`
	UserErrors []UserLoginError `json:"userErrors"`
}

type UserRegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterInvalidInputError struct {
	Message string                        `json:"message"`
	Field   UserRegisterInvalidInputField `json:"field"`
}

func (UserRegisterInvalidInputError) IsError() {}

type UserRegisterPayload struct {
	User       *User                            `json:"user"`
	UserErrors []*UserRegisterInvalidInputError `json:"userErrors"`
}

type ApplicantInput struct {
	Message string `json:"message"`
}

type CreateMessageInput struct {
	Content string `json:"content"`
}

type CreateTagInput struct {
	Name string `json:"name"`
}

type PaginationInput struct {
	First  *int    `json:"first"`
	After  *string `json:"after"`
	Last   *int    `json:"last"`
	Before *string `json:"before"`
}

type RecruitmentInput struct {
	Title         string                 `json:"title"`
	CompetitionID string                 `json:"competitionId"`
	Type          Type                   `json:"type"`
	Detail        *string                `json:"detail"`
	PrefectureID  *string                `json:"prefectureId"`
	Place         *string                `json:"place"`
	StartAt       *time.Time             `json:"startAt"`
	ClosingAt     *time.Time             `json:"closingAt"`
	LocationLat   *float64               `json:"locationLat"`
	LocationLng   *float64               `json:"locationLng"`
	Status        Status                 `json:"status"`
	Tags          []*RecruitmentTagInput `json:"tags"`
}

type RecruitmentTagInput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	IsNew bool   `json:"isNew"`
}

type SearchRecruitmentInput struct {
	CompetitionID *string    `json:"competitionId"`
	PrefectureID  *string    `json:"prefectureId"`
	Type          *string    `json:"type"`
	StartAt       *time.Time `json:"startAt"`
}

type EmailVerificationStatus string

const (
	EmailVerificationStatusPending  EmailVerificationStatus = "PENDING"
	EmailVerificationStatusVerified EmailVerificationStatus = "VERIFIED"
)

var AllEmailVerificationStatus = []EmailVerificationStatus{
	EmailVerificationStatusPending,
	EmailVerificationStatusVerified,
}

func (e EmailVerificationStatus) IsValid() bool {
	switch e {
	case EmailVerificationStatusPending, EmailVerificationStatusVerified:
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

type Status string

const (
	StatusDraft     Status = "DRAFT"
	StatusPublished Status = "PUBLISHED"
	StatusClosed    Status = "CLOSED"
)

var AllStatus = []Status{
	StatusDraft,
	StatusPublished,
	StatusClosed,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusDraft, StatusPublished, StatusClosed:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Type string

const (
	TypeOpponent   Type = "OPPONENT"
	TypeIndividual Type = "INDIVIDUAL"
	TypeMember     Type = "MEMBER"
	TypeJoining    Type = "JOINING"
	TypeOthers     Type = "OTHERS"
)

var AllType = []Type{
	TypeOpponent,
	TypeIndividual,
	TypeMember,
	TypeJoining,
	TypeOthers,
}

func (e Type) IsValid() bool {
	switch e {
	case TypeOpponent, TypeIndividual, TypeMember, TypeJoining, TypeOthers:
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

type UserLoginInvalidInputField string

const (
	UserLoginInvalidInputFieldEmail    UserLoginInvalidInputField = "EMAIL"
	UserLoginInvalidInputFieldPassword UserLoginInvalidInputField = "PASSWORD"
)

var AllUserLoginInvalidInputField = []UserLoginInvalidInputField{
	UserLoginInvalidInputFieldEmail,
	UserLoginInvalidInputFieldPassword,
}

func (e UserLoginInvalidInputField) IsValid() bool {
	switch e {
	case UserLoginInvalidInputFieldEmail, UserLoginInvalidInputFieldPassword:
		return true
	}
	return false
}

func (e UserLoginInvalidInputField) String() string {
	return string(e)
}

func (e *UserLoginInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserLoginInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserLoginInvalidInputField", str)
	}
	return nil
}

func (e UserLoginInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRegisterInvalidInputField string

const (
	UserRegisterInvalidInputFieldName     UserRegisterInvalidInputField = "NAME"
	UserRegisterInvalidInputFieldEmail    UserRegisterInvalidInputField = "EMAIL"
	UserRegisterInvalidInputFieldPassword UserRegisterInvalidInputField = "PASSWORD"
)

var AllUserRegisterInvalidInputField = []UserRegisterInvalidInputField{
	UserRegisterInvalidInputFieldName,
	UserRegisterInvalidInputFieldEmail,
	UserRegisterInvalidInputFieldPassword,
}

func (e UserRegisterInvalidInputField) IsValid() bool {
	switch e {
	case UserRegisterInvalidInputFieldName, UserRegisterInvalidInputFieldEmail, UserRegisterInvalidInputFieldPassword:
		return true
	}
	return false
}

func (e UserRegisterInvalidInputField) String() string {
	return string(e)
}

func (e *UserRegisterInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRegisterInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRegisterInvalidInputField", str)
	}
	return nil
}

func (e UserRegisterInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
