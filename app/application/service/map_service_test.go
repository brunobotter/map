package service

import (
	"context"
	"errors"
	"testing"

	"github.com/brunobotter/map/application/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type weatherServiceStub struct {
	result domain.Weather
	err    error
}

func (w *weatherServiceStub) GetWeather(ctx context.Context, lat, lng float64) (domain.Weather, error) {
	_ = ctx
	_ = lat
	_ = lng
	if w.err != nil {
		return domain.Weather{}, w.err
	}
	return w.result, nil
}

type mapRepositoryStub struct {
	traffic    []domain.Traffic
	events     []domain.MapEvent
	trafficErr error
	eventsErr  error
}

func (r *mapRepositoryStub) GetTraffic() ([]domain.Traffic, error) {
	if r.trafficErr != nil {
		return nil, r.trafficErr
	}
	return r.traffic, nil
}

func (r *mapRepositoryStub) GetEvents() ([]domain.MapEvent, error) {
	if r.eventsErr != nil {
		return nil, r.eventsErr
	}
	return r.events, nil
}

func TestMapServiceGetMapDataUsesWeatherServiceAndRepository(t *testing.T) {
	expected := domain.Weather{Status: "Rain", Temperature: 18.1, Unit: "C"}
	repo := &mapRepositoryStub{
		traffic: []domain.Traffic{{Road: "R1", Level: "light", Status: "flowing"}},
		events:  []domain.MapEvent{{Title: "E1", Location: "L1", StartAt: "2026-04-04T00:00:00Z"}},
	}
	svc := NewMapService(&weatherServiceStub{result: expected}, repo)

	result, err := svc.GetMapData(context.Background())
	require.NoError(t, err)
	assert.Equal(t, expected, result.Weather)
	assert.Equal(t, repo.traffic, result.Traffic)
	assert.Equal(t, repo.events, result.Events)
}

func TestMapServiceGetMapDataFallbacksOnError(t *testing.T) {
	repo := &mapRepositoryStub{trafficErr: errors.New("db down"), eventsErr: errors.New("db down")}
	svc := NewMapService(&weatherServiceStub{err: errors.New("openweather unavailable")}, repo)

	result, err := svc.GetMapData(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "unknown", result.Weather.Status)
	assert.Equal(t, 0.0, result.Weather.Temperature)
	assert.Equal(t, "C", result.Weather.Unit)
	assert.Empty(t, result.Traffic)
	assert.Empty(t, result.Events)
}
