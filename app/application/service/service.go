package service

import (
	"context"

	"github.com/brunobotter/map/application/domain"
	"github.com/brunobotter/map/application/integration"
)

type WeatherService interface {
	GetWeather(ctx context.Context, lat, lng float64) (domain.Weather, error)
}

type weatherService struct {
	weatherIntegration integration.WeatherIntegration
}

func NewWeatherService(weatherIntegration integration.WeatherIntegration) WeatherService {
	return &weatherService{weatherIntegration: weatherIntegration}
}

func (s *weatherService) GetWeather(ctx context.Context, lat, lng float64) (domain.Weather, error) {
	return s.weatherIntegration.GetWeather(ctx, lat, lng)
}
