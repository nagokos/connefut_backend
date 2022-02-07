// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/nagokos/connefut_backend/ent/competition"
)

// Competition is the model entity for the Competition schema.
type Competition struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompetitionQuery when eager-loading is set.
	Edges CompetitionEdges `json:"edges"`
}

// CompetitionEdges holds the relations/edges for other nodes in the graph.
type CompetitionEdges struct {
	// Recruitments holds the value of the recruitments edge.
	Recruitments []*Recruitment `json:"recruitments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RecruitmentsOrErr returns the Recruitments value or an error if the edge
// was not loaded in eager-loading.
func (e CompetitionEdges) RecruitmentsOrErr() ([]*Recruitment, error) {
	if e.loadedTypes[0] {
		return e.Recruitments, nil
	}
	return nil, &NotLoadedError{edge: "recruitments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Competition) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case competition.FieldID, competition.FieldName:
			values[i] = new(sql.NullString)
		case competition.FieldCreatedAt, competition.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Competition", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Competition fields.
func (c *Competition) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case competition.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case competition.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case competition.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case competition.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		}
	}
	return nil
}

// QueryRecruitments queries the "recruitments" edge of the Competition entity.
func (c *Competition) QueryRecruitments() *RecruitmentQuery {
	return (&CompetitionClient{config: c.config}).QueryRecruitments(c)
}

// Update returns a builder for updating this Competition.
// Note that you need to call Competition.Unwrap() before calling this method if this Competition
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Competition) Update() *CompetitionUpdateOne {
	return (&CompetitionClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Competition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Competition) Unwrap() *Competition {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Competition is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Competition) String() string {
	var builder strings.Builder
	builder.WriteString("Competition(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Competitions is a parsable slice of Competition.
type Competitions []*Competition

func (c Competitions) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}