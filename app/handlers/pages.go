package handlers

import (
	"net/http"

	"github.com/corentings/email-tracker/app/views/page"
	"github.com/labstack/echo/v4"
)

type PageController struct{}

func NewPageController() *PageController {
	return &PageController{}
}

func (p *PageController) GetIndex(c echo.Context) error {
	hero := page.Index()

	index := page.IndexPage("email-tracker", false, GetNonce(c), hero)

	return Render(c, http.StatusOK, index)
}

func (p *PageController) GetAdmin(c echo.Context) error {
	hero := page.Admin()

	admin := page.AdminPage("email-tracker", true, GetNonce(c), hero)

	return Render(c, http.StatusOK, admin)
}
