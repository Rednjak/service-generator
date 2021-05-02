package service

import (
	"context"
	"fmt"
	"module_placeholder/pkg/adapters"
	"module_placeholder/pkg/app"
	"module_placeholder/pkg/app/command"
	"module_placeholder/pkg/app/query"
	"module_placeholder/pkg/shared/config"
)

func NewApplication(context context.Context, cfg *config.Configuration) app.Application {
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseIP, cfg.DatabaseScheme)
	db, err := adapters.InitMySQLConn(cfg.DatabaseDriver, connection)
	if err != nil {
		panic(err)
	}

	resourceRepo := adapters.NewMySqlResourceRepository(db)

	return app.Application{
		Commands: app.Commands{
			CancelResource: command.NewCancelResourceHandler(resourceRepo),
			CreateResource: command.NewCreateResourceHandler(resourceRepo),
		},
		Queries: app.Queries{
			AllResources: query.NewAllResourceHandler(resourceRepo),
			GetResource:  query.NewGetResourceHandler(resourceRepo),
		},
	}
}
