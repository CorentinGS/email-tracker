package services

import (
	"sync"

	"github.com/corentings/email-tracker/app/handlers"
)

type ServiceContainer struct {
	jwtHandler   handlers.JwtMiddleware
	emailHandler handlers.EmailController
}

var (
	container *ServiceContainer //nolint:gochecknoglobals // Singleton
	once      sync.Once         //nolint:gochecknoglobals // Singleton
)

func DefaultServiceContainer() *ServiceContainer {
	emailHandler := InitializeEmailHandler()
	jwtHandler := InitializeJwtMiddleware()

	return NewServiceContainer(emailHandler, jwtHandler)
}

func NewServiceContainer(emailHandler handlers.EmailController, jwtHandler handlers.JwtMiddleware,
) *ServiceContainer {
	once.Do(func() {
		container = &ServiceContainer{
			emailHandler: emailHandler,
			jwtHandler:   jwtHandler,
		}
	})
	return container
}

func (sc *ServiceContainer) EmailHandler() handlers.EmailController {
	return sc.emailHandler
}

func (sc *ServiceContainer) JwtMiddleware() handlers.JwtMiddleware {
	return sc.jwtHandler
}
