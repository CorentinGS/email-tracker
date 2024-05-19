//go:build wireinject
// +build wireinject

package services

import (
	"github.com/corentings/email-tracker/app/handlers"
	"github.com/corentings/email-tracker/pkg/postgres"
	"github.com/corentings/email-tracker/services/email"
	"github.com/google/wire"
)

func InitializeEmailHandler() handlers.EmailController {
	wire.Build(handlers.NewEmailController, email.NewUseCase, postgres.GetPool)
	return handlers.EmailController{}
}

func InitializeJwtMiddleware() handlers.JwtMiddleware {
	wire.Build(handlers.NewJwtMiddleware)
	return handlers.JwtMiddleware{}
}
