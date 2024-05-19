package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/corentings/email-tracker/app/views/page"
	"github.com/corentings/email-tracker/domain"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, _ int, t templ.Component) error {
	c.Response().Writer.Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(c.Request().Context(), c.Response())
}

func Redirect(c echo.Context, path string, statusCode int) error {
	c.Response().Header().Set("HX-Redirect", path)
	return c.NoContent(statusCode)
}

func GetNonce(c echo.Context) domain.Nonce {
	return c.Get("nonce").(domain.Nonce)
}

func RedirectToErrorPage(c echo.Context, errorCode int) error {
	pageToReturn := page.InternalServerError()
	switch errorCode {
	case http.StatusUnauthorized:
		pageToReturn = page.NotAuthorized()
	case http.StatusNotFound:
		pageToReturn = page.NotFound()
	case http.StatusInternalServerError:
		pageToReturn = page.InternalServerError()
	case http.StatusBadRequest:
		pageToReturn = page.BadRequest()
	}

	errorPage := page.ErrorPage("email-tracker", true, GetNonce(c), pageToReturn)

	return Render(c, errorCode, errorPage)
}
