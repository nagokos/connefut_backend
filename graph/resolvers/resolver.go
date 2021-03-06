package resolvers

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nagokos/connefut_backend/graph/generated"
	"github.com/nagokos/connefut_backend/graph/model"
	"github.com/nagokos/connefut_backend/graph/models/user"
	"github.com/nagokos/connefut_backend/logger"
)

// This file will not be regenerated automatically.
//go:generate go run github.com/99designs/gqlgen

// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	dbPool *pgxpool.Pool
}

func NewSchema(dbPool *pgxpool.Pool) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{dbPool: dbPool},
		Directives: generated.DirectiveRoot{
			HasLoggedIn: func(ctx context.Context, _ interface{}, next graphql.Resolver) (res interface{}, err error) {
				viewer := user.GetViewer(ctx)
				if viewer == nil {
					logger.NewLogger().Error("user not loggedIn")
					return nil, nil
				}
				return next(ctx)
			},
			EmailVerified: func(ctx context.Context, _ interface{}, next graphql.Resolver, status model.EmailVerificationStatus) (res interface{}, err error) {
				viewer := user.GetViewer(ctx)
				if viewer.EmailVerificationStatus != status {
					logger.NewLogger().Error("access denined")
					return nil, fmt.Errorf("access denined")
				}
				return next(ctx)
			},
		},
	})
}
