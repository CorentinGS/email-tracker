package app

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/corentings/email-tracker/config"
	"github.com/corentings/email-tracker/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho(cfg *config.Config) *echo.Echo {
	echo := echo.New()
	registerMiddlewares(echo, cfg)
	registerStaticRoutes(echo)
	registerRoutes(echo)
	return echo
}

func registerMiddlewares(e *echo.Echo, cfg *config.Config) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(CSPMiddleware)

	e.Use(middleware.Recover())

	e.Use(middleware.Secure())

	csrfConfig := middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "cookie:_csrf",
		CookiePath:     "/",
		CookieDomain:   cfg.HTTP.Host,
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	})

	e.Use(csrfConfig)
}

func generateSecureNonce() (string, error) {
	nonce := make([]byte, config.NonceLength)
	_, err := rand.Read(nonce)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(nonce), err
}

func CSPMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		htmxNonce, _ := generateSecureNonce()
		hyperscriptNonce, _ := generateSecureNonce()
		picoCSSNonce, _ := generateSecureNonce()
		preloadNonce, _ := generateSecureNonce()
		umamiNonce, _ := generateSecureNonce()
		cssScopeInlineNonce, _ := generateSecureNonce()

		_ = "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg="

		cspHeader := fmt.Sprintf(
			"default-src 'self'; connect-src 'self' ; script-src 'nonce-%s' 'nonce-%s' 'nonce-%s' 'nonce-%s' 'nonce-%s'; style-src 'self' 'unsafe-inline' https://fonts.gstatic.com fonts.googleapis.com; font-src 'self' https://fonts.gstatic.com",
			htmxNonce, cssScopeInlineNonce, hyperscriptNonce, preloadNonce, umamiNonce)

		c.Response().Header().Set("Content-Security-Policy", cspHeader)

		c.Set("nonce", domain.Nonce{
			HtmxNonce:           htmxNonce,
			HyperscriptNonce:    hyperscriptNonce,
			PicoCSSNonce:        picoCSSNonce,
			PreloadNonce:        preloadNonce,
			UmamiNonce:          umamiNonce,
			CSSScopeInlineNonce: cssScopeInlineNonce,
		})

		return next(c)
	}
}

func StaticAssetsCacheControlMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
		return next(c)
	}
}

func StaticPageCacheControlMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Set Cache-Control header
		c.Response().Header().Set("Cache-Control", "private, max-age=60")

		return next(c)
	}
}
