// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/nagokos/connefut_backend/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Role holds the value of the "role" field.
	Role user.Role `json:"role,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Introduction holds the value of the "introduction" field.
	Introduction string `json:"introduction,omitempty"`
	// EmailVerificationStatus holds the value of the "email_verification_status" field.
	EmailVerificationStatus bool `json:"email_verification_status,omitempty"`
	// EmailVerificationToken holds the value of the "email_verification_token" field.
	EmailVerificationToken string `json:"email_verification_token,omitempty"`
	// EmailVerificationTokenExpiresAt holds the value of the "email_verification_token_expires_at" field.
	EmailVerificationTokenExpiresAt time.Time `json:"email_verification_token_expires_at,omitempty"`
	// PasswordDigest holds the value of the "password_digest" field.
	PasswordDigest string `json:"password_digest,omitempty"`
	// LastSignInAt holds the value of the "last_sign_in_at" field.
	LastSignInAt time.Time `json:"last_sign_in_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldEmailVerificationStatus:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldName, user.FieldEmail, user.FieldRole, user.FieldAvatar, user.FieldIntroduction, user.FieldEmailVerificationToken, user.FieldPasswordDigest:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldEmailVerificationTokenExpiresAt, user.FieldLastSignInAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = user.Role(value.String)
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldIntroduction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field introduction", values[i])
			} else if value.Valid {
				u.Introduction = value.String
			}
		case user.FieldEmailVerificationStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_verification_status", values[i])
			} else if value.Valid {
				u.EmailVerificationStatus = value.Bool
			}
		case user.FieldEmailVerificationToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email_verification_token", values[i])
			} else if value.Valid {
				u.EmailVerificationToken = value.String
			}
		case user.FieldEmailVerificationTokenExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field email_verification_token_expires_at", values[i])
			} else if value.Valid {
				u.EmailVerificationTokenExpiresAt = value.Time
			}
		case user.FieldPasswordDigest:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_digest", values[i])
			} else if value.Valid {
				u.PasswordDigest = value.String
			}
		case user.FieldLastSignInAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_sign_in_at", values[i])
			} else if value.Valid {
				u.LastSignInAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteString(", avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", introduction=")
	builder.WriteString(u.Introduction)
	builder.WriteString(", email_verification_status=")
	builder.WriteString(fmt.Sprintf("%v", u.EmailVerificationStatus))
	builder.WriteString(", email_verification_token=")
	builder.WriteString(u.EmailVerificationToken)
	builder.WriteString(", email_verification_token_expires_at=")
	builder.WriteString(u.EmailVerificationTokenExpiresAt.Format(time.ANSIC))
	builder.WriteString(", password_digest=")
	builder.WriteString(u.PasswordDigest)
	builder.WriteString(", last_sign_in_at=")
	builder.WriteString(u.LastSignInAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
