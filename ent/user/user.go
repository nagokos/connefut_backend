// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldIntroduction holds the string denoting the introduction field in the database.
	FieldIntroduction = "introduction"
	// FieldEmailVerificationStatus holds the string denoting the email_verification_status field in the database.
	FieldEmailVerificationStatus = "email_verification_status"
	// FieldEmailVerificationToken holds the string denoting the email_verification_token field in the database.
	FieldEmailVerificationToken = "email_verification_token"
	// FieldEmailVerificationTokenExpiresAt holds the string denoting the email_verification_token_expires_at field in the database.
	FieldEmailVerificationTokenExpiresAt = "email_verification_token_expires_at"
	// FieldPasswordDigest holds the string denoting the password_digest field in the database.
	FieldPasswordDigest = "password_digest"
	// FieldLastSignInAt holds the string denoting the last_sign_in_at field in the database.
	FieldLastSignInAt = "last_sign_in_at"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldEmail,
	FieldRole,
	FieldAvatar,
	FieldIntroduction,
	FieldEmailVerificationStatus,
	FieldEmailVerificationToken,
	FieldEmailVerificationTokenExpiresAt,
	FieldPasswordDigest,
	FieldLastSignInAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultAvatar holds the default value on creation for the "avatar" field.
	DefaultAvatar string
	// IntroductionValidator is a validator for the "introduction" field. It is called by the builders before save.
	IntroductionValidator func(string) error
	// DefaultEmailVerificationStatus holds the default value on creation for the "email_verification_status" field.
	DefaultEmailVerificationStatus bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Role defines the type for the "role" enum field.
type Role string

// RoleGeneral is the default value of the Role enum.
const DefaultRole = RoleGeneral

// Role values.
const (
	RoleAdmin   Role = "admin"
	RoleGeneral Role = "general"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleAdmin, RoleGeneral:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (r Role) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(r.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (r *Role) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*r = Role(str)
	if err := RoleValidator(*r); err != nil {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}
