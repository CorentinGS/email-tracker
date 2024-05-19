package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/corentings/email-tracker/config"
	"github.com/corentings/email-tracker/pkg/jwt"
	"github.com/corentings/email-tracker/services/email"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EmailController struct {
	useCase email.IUseCase
}

const (
	SessionTokenCookieKey = "session_token"
	ExpiresDuration       = 92 * time.Hour
)

func NewEmailController(email email.IUseCase) EmailController {
	return EmailController{useCase: email}
}

func (u *EmailController) Login(c echo.Context) error {
	// get form values
	token := c.FormValue("token")

	// if token is empty
	if token == "" {
		return RedirectToErrorPage(c, http.StatusBadRequest)
	}

	if token != config.AccessToken {
		return RedirectToErrorPage(c, http.StatusUnauthorized)
	}

	// create a new jwt
	jwt, err := jwt.GetJwtInstance().GetJwt().GenerateToken(c.Request().Context(), 1)
	if err != nil {
		slog.Error("Login: error generating jwt", slog.String("error", err.Error()))
		return RedirectToErrorPage(c, http.StatusInternalServerError)
	}

	// set the jwt in a cookie
	c.SetCookie(&http.Cookie{
		Name:     SessionTokenCookieKey,
		Value:    jwt,
		Expires:  time.Now().Add(ExpiresDuration),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})

	return Redirect(c, "/admin", http.StatusSeeOther)
}

func (u *EmailController) GetImage(c echo.Context) error {
	// get information from the request
	ip := c.RealIP()

	// get uuid from the request path
	uuidString := c.Param("uuid")

	// get uuid from uuidString
	imageUUID, err := uuid.Parse(uuidString)
	if err != nil {
		return c.File("assets/img/favicon.ico")
	}

	// get the email from the database
	email, err := u.useCase.GetEmail(c.Request().Context(), imageUUID)
	if err != nil {
		return c.File("assets/img/favicon.ico")
	}

	// add a tracking event
	err = u.useCase.AddTracking(c.Request().Context(), email, ip)
	if err != nil {
		slog.Error("Signature: error adding tracking event", slog.String("error", err.Error()))
		return RedirectToErrorPage(c, http.StatusInternalServerError)
	}

	// return the image
	return c.File("assets/img/favicon.ico")
}
