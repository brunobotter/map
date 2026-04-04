package service

import (
	"context"
	"errors"
	"testing"

	"github.com/brunobotter/map/application/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type weatherIntegrationStub struct {
	result domain.Weather
	err    error
}

func (w *weatherIntegrationStub) GetWeather(ctx context.Context, lat, lng float64) (domain.Weather, error) {
	_ = ctx
	_ = lat
	_ = lng
	if w.err != nil {
		return domain.Weather{}, w.err
	}
	return w.result, nil
}

func TestWeatherServiceGetWeather(t *testing.T) {
	expected := domain.Weather{Status: "Clouds", Temperature: 22.4, Unit: "C"}
	svc := NewWeatherService(&weatherIntegrationStub{result: expected})

	result, err := svc.GetWeather(context.Background(), -23.55052, -46.633308)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestWeatherServiceGetWeatherError(t *testing.T) {
	svc := NewWeatherService(&weatherIntegrationStub{err: errors.New("api unavailable")})

	_, err := svc.GetWeather(context.Background(), -23.55052, -46.633308)
	require.Error(t, err)
}
