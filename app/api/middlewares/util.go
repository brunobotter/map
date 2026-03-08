package middlewares

import (
	"github.com/brunobotter/map/infra/logger"
	"github.com/brunobotter/map/main/config"
)

func CommonMiddlewares(logger logger.Logger, cfg *config.Config) []MiddlewareFunc {
	return []MiddlewareFunc{
		NewPanicMiddleware(logger),
		NewLoggerMiddleware(logger, cfg),
	}
}
