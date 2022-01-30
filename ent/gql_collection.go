// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CompetitionQuery) CollectFields(ctx context.Context, satisfies ...string) *CompetitionQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return c
}

func (c *CompetitionQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *CompetitionQuery {
	return c
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pr *PrefectureQuery) CollectFields(ctx context.Context, satisfies ...string) *PrefectureQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		pr = pr.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return pr
}

func (pr *PrefectureQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *PrefectureQuery {
	return pr
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}
