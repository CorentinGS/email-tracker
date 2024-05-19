package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/corentings/email-tracker/app/views/page"
	"github.com/corentings/email-tracker/config"
	db "github.com/corentings/email-tracker/db/sqlc"
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
		return c.File("assets/img/favicon.ico")
	}

	// send a message to discord using a webhook
	err = SendDiscordMessage(c.Request().Context(), email, ip)
	if err != nil {
		slog.Error("Signature: error sending discord message", slog.String("error", err.Error()))
		return c.File("assets/img/favicon.ico")
	}

	// return the image
	return c.File("assets/img/favicon.ico")
}

type webhook struct {
	Content string `json:"content"`
}

const discordTimeout = 10

func SendDiscordMessage(ctx context.Context, email db.Email, ip string) error {
	// create a new discord message
	message := "```New tracking event```\n" +
		"Recipient:** " + email.Recipient + "**\n" +
		"Subject: **" + email.Subject + "**\n" +
		"> IP: " + ip

	webhookContent := webhook{
		Content: message,
	}

	jsonWebhook, _ := json.Marshal(webhookContent)

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(ctx, discordTimeout*time.Second)
	// Cancel the context when we are done to release resources
	defer cancel()

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.WebhookURL, bytes.NewBuffer(jsonWebhook))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return err
}

func (u *EmailController) PostEmail(c echo.Context) error {
	// get form values
	recipient := c.FormValue("recipient")
	subject := c.FormValue("subject")

	// if recipient or subject is empty
	if recipient == "" || subject == "" {
		return RedirectToErrorPage(c, http.StatusBadRequest)
	}

	// create a new email
	email, err := u.useCase.CreateEmail(c.Request().Context(), recipient, subject)
	if err != nil {
		slog.Error("CreateEmail: error creating email", slog.String("error", err.Error()))
		return RedirectToErrorPage(c, http.StatusInternalServerError)
	}

	adminErrorComponent := page.AdminError("Email link: https://tracker.corentings.dev/image/" + email.Uuid.String())

	return Render(c, http.StatusOK, adminErrorComponent)
}

func (u *EmailController) GetEmails(c echo.Context) error {
	// get page and limit values
	pageParam, limitParam := getPageLimitValues(c)

	if pageParam < 1 {
		pageParam = 1
	}

	if limitParam < 1 {
		limitParam = 10
	}

	// get emails from the database
	emails, err := u.useCase.GetEmailsWithPagination(c.Request().Context(), limitParam, (pageParam-1)*limitParam)
	if err != nil {
		slog.Error("GetEmails: error getting emails", slog.String("error", err.Error()))
		return RedirectToErrorPage(c, http.StatusInternalServerError)
	}

	if len(emails) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	slog.Debug("GetEmails: emails", slog.Int("count", len(emails)))
	// render the emails
	adminEmailsComponent := page.ListEmails(emails, pageParam)

	return Render(c, http.StatusOK, adminEmailsComponent)
}

func (u *EmailController) GetTrackers(c echo.Context) error {
	// get page and limit values
	pageParam, limitParam := getPageLimitValues(c)

	if pageParam < 1 {
		pageParam = 1
	}

	if limitParam < 1 {
		limitParam = 10
	}

	// get trackers from the database
	trackers, err := u.useCase.GetTrackersWithPagination(c.Request().Context(), limitParam, (pageParam-1)*limitParam)
	if err != nil {
		slog.Error("GetTrackers: error getting trackers", slog.String("error", err.Error()))
		return RedirectToErrorPage(c, http.StatusInternalServerError)
	}

	if len(trackers) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	slog.Debug("GetTrackers: trackers", slog.Int("count", len(trackers)))
	// render the trackers
	adminTrackersComponent := page.ListTrackers(trackers, pageParam)

	return Render(c, http.StatusOK, adminTrackersComponent)
}

func getPageLimitValues(c echo.Context) (int, int) {
	pageParam := c.QueryParam("page")
	limit := c.QueryParam("limit")

	// Convert the page number and limit to integers
	pageParamInt, err := strconv.Atoi(pageParam)
	if err != nil {
		return 0, 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return 0, 0
	}

	return pageParamInt, limitInt
}
