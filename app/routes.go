package app

import (
	"net/http"

	"github.com/corentings/email-tracker/app/handlers"
	"github.com/corentings/email-tracker/assets"
	"github.com/corentings/email-tracker/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func registerStaticRoutes(e *echo.Echo) {
	g := e.Group("/static", StaticAssetsCacheControlMiddleware)

	g.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       ".",
		Browse:     false,
		Filesystem: assets.Assets(),
	}))
}

func registerRoutes(e *echo.Echo) {
	serviceContainer := services.DefaultServiceContainer()
	pageController := handlers.NewPageController()
	jwtMiddleware := serviceContainer.JwtMiddleware()
	emailController := serviceContainer.EmailHandler()

	// Page routes
	e.GET("/", pageController.GetIndex)

	// JWT routes
	e.POST("/login", emailController.Login)

	connectedGroup := e.Group("/admin")
	connectedGroup.Use(jwtMiddleware.AuthorizeUser)
	connectedGroup.GET("", pageController.GetAdmin)

	// Email routes
	e.GET("/image", emailController.GetImage)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}
