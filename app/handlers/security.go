package handlers

import (
	"log/slog"
	"net/http"

	"github.com/corentings/email-tracker/domain"
	"github.com/corentings/email-tracker/pkg/jwt"
	"github.com/labstack/echo/v4"
)

// JwtMiddleware is the controller for the jwt routes.
type JwtMiddleware struct{}

// NewJwtController creates a new jwt controller.
func NewJwtMiddleware() JwtMiddleware {
	return JwtMiddleware{}
}

func (j *JwtMiddleware) AuthorizeUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return j.IsConnectedMiddleware(domain.PermissionUser, next)(c)
	}
}

func (j *JwtMiddleware) IsConnectedMiddleware(_ domain.Permission, next echo.HandlerFunc) func(c echo.Context) error {
	return func(c echo.Context) error {
		// get the token from the cookie
		cookie, err := c.Cookie("session_token")
		if err != nil {
			slog.Error("IsConnectedMiddleware: error getting cookie", slog.String("error", err.Error()))
			_ = RedirectToErrorPage(c, http.StatusUnauthorized)
		}

		token := cookie.Value

		_, err = jwt.GetJwtInstance().GetConnectedUserID(c.Request().Context(), token)
		if err != nil {
			slog.Error("IsConnectedMiddleware: error getting connected user id", slog.String("error", err.Error()))
			_ = RedirectToErrorPage(c, http.StatusUnauthorized)
		}

		return next(c)
	}
}
