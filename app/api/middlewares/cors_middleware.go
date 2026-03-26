package middlewares

import (
	"net/http"

	"github.com/brunobotter/map/main/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CORSMiddleware struct {
	middlewareFunc any
}

func (m *CORSMiddleware) GetMiddleware() any {
	return m.middlewareFunc
}

func NewCORSMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions},
		AllowHeaders:     []string{"content-type", "x-itau-apikey", "x-itau-correlationid", "x-charon-params", "x-charon-session", "x-itau-recaptcha-token", "x-journey-token", "x-journey-id", "user-agent"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           600,
	})
}
