package integration

import (
	"context"

	"github.com/brunobotter/map/application/domain"
)

type WeatherIntegration interface {
	GetWeather(ctx context.Context, lat, lng float64) (domain.Weather, error)
}
