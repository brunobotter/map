package server

import (
	"github.com/brunobotter/map/api/controllers"
	"github.com/brunobotter/map/main/server/router"
)

func (s *Server) setupApiRouter(healthController *controllers.HealthHandler, mapHandler *controllers.MapHandler) {
	var routs router.Router
	s.container.NamedResolve(&routs, "Routes")

	s.echo.GET("/health", healthController.Health)
	routs.GET("/map", mapHandler.MapData)
}
