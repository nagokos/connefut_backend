// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type ApplyForRecruitmentError interface {
	IsApplyForRecruitmentError()
}

type ChangeUserEmailResult interface {
	IsChangeUserEmailResult()
}

type ChangeUserPasswordResult interface {
	IsChangeUserPasswordResult()
}

type Connection interface {
	IsConnection()
}

type CreateRecruitmentResult interface {
	IsCreateRecruitmentResult()
}

type CreateTagResult interface {
	IsCreateTagResult()
}

type Edge interface {
	IsEdge()
}

type Error interface {
	IsError()
}

type LoginUserResult interface {
	IsLoginUserResult()
}

type Node interface {
	IsNode()
}

type RegisterUserResult interface {
	IsRegisterUserResult()
}

type ResetUserPasswordResult interface {
	IsResetUserPasswordResult()
}

type SendResetPasswordEmailToUserResult interface {
	IsSendResetPasswordEmailToUserResult()
}

type UpdateRecruitmentResult interface {
	IsUpdateRecruitmentResult()
}

type UpdateUserResult interface {
	IsUpdateUserResult()
}

type VerifyUserEmailResult interface {
	IsVerifyUserEmailResult()
}

type AddStockResult struct {
	FeedbackStock   *FeedbackStock   `json:"feedbackStock"`
	RecruitmentEdge *RecruitmentEdge `json:"recruitmentEdge"`
}

type Applicant struct {
	ID          string       `json:"id"`
	DatabaseID  int          `json:"databaseId"`
	Message     string       `json:"message"`
	CreatedAt   time.Time    `json:"createdAt"`
	Recruitment *Recruitment `json:"recruitment"`
}

func (Applicant) IsNode() {}

type ApplyForRecruitmentAuthorizationError struct {
	Message string `json:"message"`
}

func (ApplyForRecruitmentAuthorizationError) IsApplyForRecruitmentError() {}
func (ApplyForRecruitmentAuthorizationError) IsError()                    {}

type ApplyForRecruitmentInput struct {
	Message string `json:"message"`
}

type ApplyForRecruitmentInvalidInputError struct {
	Message string                               `json:"message"`
	Field   ApplyForRecruitmentInvalidInputField `json:"field"`
}

func (ApplyForRecruitmentInvalidInputError) IsApplyForRecruitmentError() {}
func (ApplyForRecruitmentInvalidInputError) IsError()                    {}

type ApplyForRecruitmentPayload struct {
	Feedback *FeedbackApplicant         `json:"feedback"`
	Errors   []ApplyForRecruitmentError `json:"errors"`
}

type ApplyForRecruitmentSelfGeneratedError struct {
	Message string `json:"message"`
}

func (ApplyForRecruitmentSelfGeneratedError) IsApplyForRecruitmentError() {}
func (ApplyForRecruitmentSelfGeneratedError) IsError()                    {}

type ChangeUserEmailInput struct {
	NewEmail string `json:"newEmail"`
}

type ChangeUserEmailInvalidInputError struct {
	Message string                           `json:"message"`
	Field   ChangeUserEmailInvalidInputField `json:"field"`
}

func (ChangeUserEmailInvalidInputError) IsChangeUserEmailResult() {}
func (ChangeUserEmailInvalidInputError) IsError()                 {}

type ChangeUserEmailSuccess struct {
	Viewer *Viewer `json:"viewer"`
}

func (ChangeUserEmailSuccess) IsChangeUserEmailResult() {}

type ChangeUserPasswordAuthenticationError struct {
	Message string `json:"message"`
}

func (ChangeUserPasswordAuthenticationError) IsChangeUserPasswordResult() {}
func (ChangeUserPasswordAuthenticationError) IsError()                    {}

type ChangeUserPasswordInput struct {
	CurrentPassword         string `json:"currentPassword"`
	NewPassword             string `json:"newPassword"`
	NewPasswordConfirmation string `json:"newPasswordConfirmation"`
}

type ChangeUserPasswordInvalidInputError struct {
	Message string                              `json:"message"`
	Field   ChangeUserPasswordInvalidInputField `json:"field"`
}

func (ChangeUserPasswordInvalidInputError) IsError() {}

type ChangeUserPasswordInvalidInputErrors struct {
	InvalidInputs []*ChangeUserPasswordInvalidInputError `json:"invalidInputs"`
}

func (ChangeUserPasswordInvalidInputErrors) IsChangeUserPasswordResult() {}

type ChangeUserPasswordSuccess struct {
	IsChangedPassword bool `json:"isChangedPassword"`
}

func (ChangeUserPasswordSuccess) IsChangeUserPasswordResult() {}

type CreateRecruitmentInvalidInputError struct {
	Field   RecruitmentInvalidInputField `json:"field"`
	Message string                       `json:"message"`
}

type CreateRecruitmentInvalidInputErrors struct {
	InvalidInputs []*CreateRecruitmentInvalidInputError `json:"invalidInputs"`
}

func (CreateRecruitmentInvalidInputErrors) IsCreateRecruitmentResult() {}

type CreateRecruitmentSuccess struct {
	RecruitmentEdge *RecruitmentEdge `json:"recruitmentEdge"`
}

func (CreateRecruitmentSuccess) IsCreateRecruitmentResult() {}

type CreateTagInput struct {
	Name string `json:"name"`
}

type CreateTagInvalidInputError struct {
	Field   CreateTagInvalidInputField `json:"field"`
	Message string                     `json:"message"`
}

type CreateTagInvalidInputErrors struct {
	InvalidInputs []*CreateTagInvalidInputError `json:"invalidInputs"`
}

func (CreateTagInvalidInputErrors) IsCreateTagResult() {}

type CreateTagSuccess struct {
	TagEdge *TagEdge `json:"tagEdge"`
}

func (CreateTagSuccess) IsCreateTagResult() {}

type DeleteRecruitmentResult struct {
	DeletedRecruitmentID string `json:"deletedRecruitmentId"`
}

type Entrie struct {
	User *User `json:"user"`
}

type FeedbackApplicant struct {
	ID           string `json:"id"`
	IsAppliedFor bool   `json:"isAppliedFor"`
}

func (FeedbackApplicant) IsNode() {}

type FollowConnection struct {
	PageInfo *PageInfo     `json:"pageInfo"`
	Edges    []*FollowEdge `json:"edges"`
}

func (FollowConnection) IsConnection() {}

type FollowEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

func (FollowEdge) IsEdge() {}

type FollowResult struct {
	FeedbackFollow *FeedbackFollow `json:"feedbackFollow"`
	Viewer         *Viewer         `json:"viewer"`
}

type LoginUserAuthenticationError struct {
	Message string `json:"message"`
}

func (LoginUserAuthenticationError) IsLoginUserResult() {}
func (LoginUserAuthenticationError) IsError()           {}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserInvalidInputError struct {
	Message string                     `json:"message"`
	Field   LoginUserInvalidInputField `json:"field"`
}

func (LoginUserInvalidInputError) IsError() {}

type LoginUserInvalidInputErrors struct {
	InvalidInputs []*LoginUserInvalidInputError `json:"invalidInputs"`
}

func (LoginUserInvalidInputErrors) IsLoginUserResult() {}

type LoginUserNotFoundError struct {
	Message string `json:"message"`
}

func (LoginUserNotFoundError) IsLoginUserResult() {}
func (LoginUserNotFoundError) IsError()           {}

type LoginUserSuccess struct {
	Viewer *Viewer `json:"viewer"`
}

func (LoginUserSuccess) IsLoginUserResult() {}

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
	ID         string `json:"id"`
	DatabaseID int    `json:"databaseId"`
	Name       string `json:"name"`
}

func (Prefecture) IsNode() {}

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

type RecruitmentInput struct {
	Title        string     `json:"title"`
	SportID      string     `json:"sportId"`
	Type         Type       `json:"type"`
	Detail       *string    `json:"detail"`
	PrefectureID string     `json:"prefectureId"`
	Venue        *string    `json:"venue"`
	StartAt      *time.Time `json:"startAt"`
	ClosingAt    *time.Time `json:"closingAt"`
	LocationLat  *float64   `json:"locationLat"`
	LocationLng  *float64   `json:"locationLng"`
	Status       Status     `json:"status"`
	TagIds       []string   `json:"tagIds"`
}

type RegisterUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserInvalidInputError struct {
	Message string                        `json:"message"`
	Field   RegisterUserInvalidInputField `json:"field"`
}

func (RegisterUserInvalidInputError) IsError() {}

type RegisterUserInvalidInputErrors struct {
	InvalidInputs []*RegisterUserInvalidInputError `json:"invalidInputs"`
}

func (RegisterUserInvalidInputErrors) IsRegisterUserResult() {}

type RegisterUserSuccess struct {
	Viewer *Viewer `json:"viewer"`
}

func (RegisterUserSuccess) IsRegisterUserResult() {}

type RemoveStockResult struct {
	FeedbackStock        *FeedbackStock `json:"feedbackStock"`
	RemovedRecruitmentID string         `json:"removedRecruitmentId"`
}

type ResetUserPasswordInput struct {
	NewPassword             string `json:"newPassword"`
	NewPasswordConfirmation string `json:"newPasswordConfirmation"`
}

type ResetUserPasswordInvalidInputError struct {
	Field   ResetUserPasswordInvalidInputField `json:"field"`
	Message string                             `json:"message"`
}

type ResetUserPasswordInvalidInputErrors struct {
	InvalidInputs []*ResetUserPasswordInvalidInputError `json:"invalidInputs"`
}

func (ResetUserPasswordInvalidInputErrors) IsResetUserPasswordResult() {}

type ResetUserPasswordInvalidTokenError struct {
	Message string `json:"message"`
}

func (ResetUserPasswordInvalidTokenError) IsResetUserPasswordResult() {}

type ResetUserPasswordSuccess struct {
	Viewer *Viewer `json:"viewer"`
}

func (ResetUserPasswordSuccess) IsResetUserPasswordResult() {}

type Room struct {
	ID     string  `json:"id"`
	Entrie *Entrie `json:"entrie"`
}

type SendResetPasswordEmailToUserInput struct {
	Email string `json:"email"`
}

type SendResetPasswordEmailToUserInvalidInputError struct {
	Field   SendResetPasswordEmailToUserInvalidInputField `json:"field"`
	Message string                                        `json:"message"`
}

func (SendResetPasswordEmailToUserInvalidInputError) IsSendResetPasswordEmailToUserResult() {}

type SendResetPasswordEmailToUserNotFoundError struct {
	Message string `json:"message"`
}

func (SendResetPasswordEmailToUserNotFoundError) IsSendResetPasswordEmailToUserResult() {}

type SendResetPasswordEmailToUserSuccess struct {
	IsSentEmail bool `json:"isSentEmail"`
}

func (SendResetPasswordEmailToUserSuccess) IsSendResetPasswordEmailToUserResult() {}

type Sport struct {
	ID         string `json:"id"`
	DatabaseID int    `json:"databaseId"`
	Name       string `json:"name"`
}

func (Sport) IsNode() {}

type Tag struct {
	ID         string `json:"id"`
	DatabaseID int    `json:"databaseId"`
	Name       string `json:"name"`
}

func (Tag) IsNode() {}

type TagConnection struct {
	PageInfo *PageInfo  `json:"pageInfo"`
	Edges    []*TagEdge `json:"edges"`
}

func (TagConnection) IsConnection() {}

type TagEdge struct {
	Cursor string `json:"cursor"`
	Node   *Tag   `json:"node"`
}

func (TagEdge) IsEdge() {}

type UnFollowResult struct {
	FeedbackFollow *FeedbackFollow `json:"feedbackFollow"`
	Viewer         *Viewer         `json:"viewer"`
}

type UpdateRecruitmentInvalidInputError struct {
	Field   RecruitmentInvalidInputField `json:"field"`
	Message string                       `json:"message"`
}

type UpdateRecruitmentInvalidInputErrors struct {
	InvalidInputs []*UpdateRecruitmentInvalidInputError `json:"invalidInputs"`
}

func (UpdateRecruitmentInvalidInputErrors) IsUpdateRecruitmentResult() {}

type UpdateRecruitmentSuccess struct {
	Recruitment *Recruitment `json:"recruitment"`
}

func (UpdateRecruitmentSuccess) IsUpdateRecruitmentResult() {}

type UpdateUserInput struct {
	Name          string   `json:"name"`
	Introduction  string   `json:"introduction"`
	SportIds      []string `json:"sportIds"`
	PrefectureIds []string `json:"prefectureIds"`
	WebsiteURL    string   `json:"websiteURL"`
}

type UpdateUserInvalidInputError struct {
	Field   UpdateUserInvalidInputField `json:"field"`
	Message string                      `json:"message"`
}

type UpdateUserInvalidInputErrors struct {
	InvalidInputs []*UpdateUserInvalidInputError `json:"invalidInputs"`
}

func (UpdateUserInvalidInputErrors) IsUpdateUserResult() {}

type UpdateUserSuccess struct {
	Viewer *Viewer `json:"viewer"`
}

func (UpdateUserSuccess) IsUpdateUserResult() {}

type VerifyUserEmailAuthenticationError struct {
	Message string `json:"message"`
}

func (VerifyUserEmailAuthenticationError) IsVerifyUserEmailResult() {}
func (VerifyUserEmailAuthenticationError) IsError()                 {}

type VerifyUserEmailCodeExpiredError struct {
	Message string `json:"message"`
}

func (VerifyUserEmailCodeExpiredError) IsVerifyUserEmailResult() {}
func (VerifyUserEmailCodeExpiredError) IsError()                 {}

type VerifyUserEmailInput struct {
	Code string `json:"code"`
}

type VerifyUserEmailInvalidInputError struct {
	Message string                           `json:"message"`
	Field   VerifyUserEmailInvalidInputField `json:"field"`
}

func (VerifyUserEmailInvalidInputError) IsVerifyUserEmailResult() {}
func (VerifyUserEmailInvalidInputError) IsError()                 {}

type VerifyUserEmailSuccess struct {
	Viewer *Viewer `json:"viewer"`
}

func (VerifyUserEmailSuccess) IsVerifyUserEmailResult() {}

type Viewer struct {
	AccountUser *User `json:"accountUser"`
}

type ApplicantInput struct {
	Message string `json:"message"`
}

type CreateMessageInput struct {
	Content string `json:"content"`
}

type ApplyForRecruitmentInvalidInputField string

const (
	ApplyForRecruitmentInvalidInputFieldMessage ApplyForRecruitmentInvalidInputField = "MESSAGE"
)

var AllApplyForRecruitmentInvalidInputField = []ApplyForRecruitmentInvalidInputField{
	ApplyForRecruitmentInvalidInputFieldMessage,
}

func (e ApplyForRecruitmentInvalidInputField) IsValid() bool {
	switch e {
	case ApplyForRecruitmentInvalidInputFieldMessage:
		return true
	}
	return false
}

func (e ApplyForRecruitmentInvalidInputField) String() string {
	return string(e)
}

func (e *ApplyForRecruitmentInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ApplyForRecruitmentInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ApplyForRecruitmentInvalidInputField", str)
	}
	return nil
}

func (e ApplyForRecruitmentInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ChangeUserEmailInvalidInputField string

const (
	ChangeUserEmailInvalidInputFieldNewEmail ChangeUserEmailInvalidInputField = "NEW_EMAIL"
)

var AllChangeUserEmailInvalidInputField = []ChangeUserEmailInvalidInputField{
	ChangeUserEmailInvalidInputFieldNewEmail,
}

func (e ChangeUserEmailInvalidInputField) IsValid() bool {
	switch e {
	case ChangeUserEmailInvalidInputFieldNewEmail:
		return true
	}
	return false
}

func (e ChangeUserEmailInvalidInputField) String() string {
	return string(e)
}

func (e *ChangeUserEmailInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ChangeUserEmailInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ChangeUserEmailInvalidInputField", str)
	}
	return nil
}

func (e ChangeUserEmailInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ChangeUserPasswordInvalidInputField string

const (
	ChangeUserPasswordInvalidInputFieldCurrentPassword         ChangeUserPasswordInvalidInputField = "CURRENT_PASSWORD"
	ChangeUserPasswordInvalidInputFieldNewPassword             ChangeUserPasswordInvalidInputField = "NEW_PASSWORD"
	ChangeUserPasswordInvalidInputFieldNewPasswordConfirmation ChangeUserPasswordInvalidInputField = "NEW_PASSWORD_CONFIRMATION"
)

var AllChangeUserPasswordInvalidInputField = []ChangeUserPasswordInvalidInputField{
	ChangeUserPasswordInvalidInputFieldCurrentPassword,
	ChangeUserPasswordInvalidInputFieldNewPassword,
	ChangeUserPasswordInvalidInputFieldNewPasswordConfirmation,
}

func (e ChangeUserPasswordInvalidInputField) IsValid() bool {
	switch e {
	case ChangeUserPasswordInvalidInputFieldCurrentPassword, ChangeUserPasswordInvalidInputFieldNewPassword, ChangeUserPasswordInvalidInputFieldNewPasswordConfirmation:
		return true
	}
	return false
}

func (e ChangeUserPasswordInvalidInputField) String() string {
	return string(e)
}

func (e *ChangeUserPasswordInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ChangeUserPasswordInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ChangeUserPasswordInvalidInputField", str)
	}
	return nil
}

func (e ChangeUserPasswordInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CreateTagInvalidInputField string

const (
	CreateTagInvalidInputFieldName CreateTagInvalidInputField = "NAME"
)

var AllCreateTagInvalidInputField = []CreateTagInvalidInputField{
	CreateTagInvalidInputFieldName,
}

func (e CreateTagInvalidInputField) IsValid() bool {
	switch e {
	case CreateTagInvalidInputFieldName:
		return true
	}
	return false
}

func (e CreateTagInvalidInputField) String() string {
	return string(e)
}

func (e *CreateTagInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CreateTagInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CreateTagInvalidInputField", str)
	}
	return nil
}

func (e CreateTagInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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

type LoginUserInvalidInputField string

const (
	LoginUserInvalidInputFieldEmail    LoginUserInvalidInputField = "EMAIL"
	LoginUserInvalidInputFieldPassword LoginUserInvalidInputField = "PASSWORD"
)

var AllLoginUserInvalidInputField = []LoginUserInvalidInputField{
	LoginUserInvalidInputFieldEmail,
	LoginUserInvalidInputFieldPassword,
}

func (e LoginUserInvalidInputField) IsValid() bool {
	switch e {
	case LoginUserInvalidInputFieldEmail, LoginUserInvalidInputFieldPassword:
		return true
	}
	return false
}

func (e LoginUserInvalidInputField) String() string {
	return string(e)
}

func (e *LoginUserInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LoginUserInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LoginUserInvalidInputField", str)
	}
	return nil
}

func (e LoginUserInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RecruitmentInvalidInputField string

const (
	RecruitmentInvalidInputFieldTitle        RecruitmentInvalidInputField = "TITLE"
	RecruitmentInvalidInputFieldSportID      RecruitmentInvalidInputField = "SPORT_ID"
	RecruitmentInvalidInputFieldType         RecruitmentInvalidInputField = "TYPE"
	RecruitmentInvalidInputFieldDetail       RecruitmentInvalidInputField = "DETAIL"
	RecruitmentInvalidInputFieldPrefectureID RecruitmentInvalidInputField = "PREFECTURE_ID"
	RecruitmentInvalidInputFieldVenue        RecruitmentInvalidInputField = "VENUE"
	RecruitmentInvalidInputFieldStartAt      RecruitmentInvalidInputField = "START_AT"
	RecruitmentInvalidInputFieldClosingAt    RecruitmentInvalidInputField = "CLOSING_AT"
)

var AllRecruitmentInvalidInputField = []RecruitmentInvalidInputField{
	RecruitmentInvalidInputFieldTitle,
	RecruitmentInvalidInputFieldSportID,
	RecruitmentInvalidInputFieldType,
	RecruitmentInvalidInputFieldDetail,
	RecruitmentInvalidInputFieldPrefectureID,
	RecruitmentInvalidInputFieldVenue,
	RecruitmentInvalidInputFieldStartAt,
	RecruitmentInvalidInputFieldClosingAt,
}

func (e RecruitmentInvalidInputField) IsValid() bool {
	switch e {
	case RecruitmentInvalidInputFieldTitle, RecruitmentInvalidInputFieldSportID, RecruitmentInvalidInputFieldType, RecruitmentInvalidInputFieldDetail, RecruitmentInvalidInputFieldPrefectureID, RecruitmentInvalidInputFieldVenue, RecruitmentInvalidInputFieldStartAt, RecruitmentInvalidInputFieldClosingAt:
		return true
	}
	return false
}

func (e RecruitmentInvalidInputField) String() string {
	return string(e)
}

func (e *RecruitmentInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RecruitmentInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RecruitmentInvalidInputField", str)
	}
	return nil
}

func (e RecruitmentInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RegisterUserInvalidInputField string

const (
	RegisterUserInvalidInputFieldName     RegisterUserInvalidInputField = "NAME"
	RegisterUserInvalidInputFieldEmail    RegisterUserInvalidInputField = "EMAIL"
	RegisterUserInvalidInputFieldPassword RegisterUserInvalidInputField = "PASSWORD"
)

var AllRegisterUserInvalidInputField = []RegisterUserInvalidInputField{
	RegisterUserInvalidInputFieldName,
	RegisterUserInvalidInputFieldEmail,
	RegisterUserInvalidInputFieldPassword,
}

func (e RegisterUserInvalidInputField) IsValid() bool {
	switch e {
	case RegisterUserInvalidInputFieldName, RegisterUserInvalidInputFieldEmail, RegisterUserInvalidInputFieldPassword:
		return true
	}
	return false
}

func (e RegisterUserInvalidInputField) String() string {
	return string(e)
}

func (e *RegisterUserInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RegisterUserInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RegisterUserInvalidInputField", str)
	}
	return nil
}

func (e RegisterUserInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ResetUserPasswordInvalidInputField string

const (
	ResetUserPasswordInvalidInputFieldNewPassword             ResetUserPasswordInvalidInputField = "NEW_PASSWORD"
	ResetUserPasswordInvalidInputFieldNewPasswordConfirmation ResetUserPasswordInvalidInputField = "NEW_PASSWORD_CONFIRMATION"
)

var AllResetUserPasswordInvalidInputField = []ResetUserPasswordInvalidInputField{
	ResetUserPasswordInvalidInputFieldNewPassword,
	ResetUserPasswordInvalidInputFieldNewPasswordConfirmation,
}

func (e ResetUserPasswordInvalidInputField) IsValid() bool {
	switch e {
	case ResetUserPasswordInvalidInputFieldNewPassword, ResetUserPasswordInvalidInputFieldNewPasswordConfirmation:
		return true
	}
	return false
}

func (e ResetUserPasswordInvalidInputField) String() string {
	return string(e)
}

func (e *ResetUserPasswordInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ResetUserPasswordInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ResetUserPasswordInvalidInputField", str)
	}
	return nil
}

func (e ResetUserPasswordInvalidInputField) MarshalGQL(w io.Writer) {
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

type SendResetPasswordEmailToUserInvalidInputField string

const (
	SendResetPasswordEmailToUserInvalidInputFieldEmail SendResetPasswordEmailToUserInvalidInputField = "EMAIL"
)

var AllSendResetPasswordEmailToUserInvalidInputField = []SendResetPasswordEmailToUserInvalidInputField{
	SendResetPasswordEmailToUserInvalidInputFieldEmail,
}

func (e SendResetPasswordEmailToUserInvalidInputField) IsValid() bool {
	switch e {
	case SendResetPasswordEmailToUserInvalidInputFieldEmail:
		return true
	}
	return false
}

func (e SendResetPasswordEmailToUserInvalidInputField) String() string {
	return string(e)
}

func (e *SendResetPasswordEmailToUserInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SendResetPasswordEmailToUserInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SendResetPasswordEmailToUserInvalidInputField", str)
	}
	return nil
}

func (e SendResetPasswordEmailToUserInvalidInputField) MarshalGQL(w io.Writer) {
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
	TypeOpponent Type = "OPPONENT"
	TypePersonal Type = "PERSONAL"
	TypeMember   Type = "MEMBER"
	TypeJoin     Type = "JOIN"
	TypeOther    Type = "OTHER"
)

var AllType = []Type{
	TypeOpponent,
	TypePersonal,
	TypeMember,
	TypeJoin,
	TypeOther,
}

func (e Type) IsValid() bool {
	switch e {
	case TypeOpponent, TypePersonal, TypeMember, TypeJoin, TypeOther:
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

type UpdateUserInvalidInputField string

const (
	UpdateUserInvalidInputFieldName         UpdateUserInvalidInputField = "NAME"
	UpdateUserInvalidInputFieldIntroduction UpdateUserInvalidInputField = "INTRODUCTION"
	UpdateUserInvalidInputFieldWebsiteURL   UpdateUserInvalidInputField = "WEBSITE_URL"
)

var AllUpdateUserInvalidInputField = []UpdateUserInvalidInputField{
	UpdateUserInvalidInputFieldName,
	UpdateUserInvalidInputFieldIntroduction,
	UpdateUserInvalidInputFieldWebsiteURL,
}

func (e UpdateUserInvalidInputField) IsValid() bool {
	switch e {
	case UpdateUserInvalidInputFieldName, UpdateUserInvalidInputFieldIntroduction, UpdateUserInvalidInputFieldWebsiteURL:
		return true
	}
	return false
}

func (e UpdateUserInvalidInputField) String() string {
	return string(e)
}

func (e *UpdateUserInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UpdateUserInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UpdateUserInvalidInputField", str)
	}
	return nil
}

func (e UpdateUserInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VerifyUserEmailInvalidInputField string

const (
	VerifyUserEmailInvalidInputFieldCode VerifyUserEmailInvalidInputField = "CODE"
)

var AllVerifyUserEmailInvalidInputField = []VerifyUserEmailInvalidInputField{
	VerifyUserEmailInvalidInputFieldCode,
}

func (e VerifyUserEmailInvalidInputField) IsValid() bool {
	switch e {
	case VerifyUserEmailInvalidInputFieldCode:
		return true
	}
	return false
}

func (e VerifyUserEmailInvalidInputField) String() string {
	return string(e)
}

func (e *VerifyUserEmailInvalidInputField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VerifyUserEmailInvalidInputField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VerifyUserEmailInvalidInputField", str)
	}
	return nil
}

func (e VerifyUserEmailInvalidInputField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
