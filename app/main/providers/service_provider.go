package providers

import (
	httpContract "github.com/brunobotter/map/application/http"
	"github.com/brunobotter/map/application/integration"
	"github.com/brunobotter/map/application/repo"
	"github.com/brunobotter/map/application/service"
	infraHttp "github.com/brunobotter/map/infra/http"
	infraIntegration "github.com/brunobotter/map/infra/integration"
	"github.com/brunobotter/map/infra/logger"
	"github.com/brunobotter/map/main/config"
	"github.com/brunobotter/map/main/container"
)

type ServiceProvider struct{}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
func (p *ServiceProvider) Register(c container.Container) {
	c.Singleton(func(cfg *config.Config, appLogger logger.Logger) httpContract.Client {
		return infraHttp.NewHttpClient(cfg, appLogger)
	})

	c.Singleton(func(httpClient httpContract.Client, cfg *config.Config) integration.WeatherIntegration {
		return infraIntegration.NewOpenWeatherIntegration(httpClient, cfg)
	})

	c.Singleton(func(weatherIntegration integration.WeatherIntegration) service.WeatherService {
		return service.NewWeatherService(weatherIntegration)
	})

	c.Singleton(func() service.TrafficService {
		return service.NewTrafficService()
	})

	c.Singleton(func(weatherService service.WeatherService, trafficService service.TrafficService, mapRepository repo.MapRepository) service.MapService {
		return service.NewMapService(weatherService, trafficService, mapRepository)
	})
}
