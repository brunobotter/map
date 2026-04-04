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
	events    []domain.MapEvent
	eventsErr error
}

type trafficServiceStub struct {
	incidents []domain.Traffic
	err       error
}

func (s *trafficServiceStub) GetIncidents(ctx context.Context, lat, lng float64) ([]domain.Traffic, error) {
	_ = ctx
	_ = lat
	_ = lng
	if s.err != nil {
		return nil, s.err
	}
	return s.incidents, nil
}

func (r *mapRepositoryStub) GetEvents() ([]domain.MapEvent, error) {
	if r.eventsErr != nil {
		return nil, r.eventsErr
	}
	return r.events, nil
}

func TestMapServiceGetMapDataUsesWeatherServiceAndRepository(t *testing.T) {
	expected := domain.Weather{Status: "Rain", Temperature: 18.1, Unit: "C"}
	traffic := []domain.Traffic{{Lat: -23.55001, Lng: -46.63399, Type: "accident"}}
	repo := &mapRepositoryStub{
		events: []domain.MapEvent{{Title: "E1", Location: "L1", StartAt: "2026-04-04T00:00:00Z"}},
	}
	svc := NewMapService(&weatherServiceStub{result: expected}, &trafficServiceStub{incidents: traffic}, repo)

	result, err := svc.GetMapData(context.Background())
	require.NoError(t, err)
	assert.Equal(t, expected, result.Weather)
	assert.Equal(t, traffic, result.Traffic)
	assert.Equal(t, repo.events, result.Events)
}

func TestMapServiceGetMapDataFallbacksOnError(t *testing.T) {
	repo := &mapRepositoryStub{eventsErr: errors.New("db down")}
	svc := NewMapService(&weatherServiceStub{err: errors.New("openweather unavailable")}, &trafficServiceStub{err: errors.New("traffic unavailable")}, repo)

	result, err := svc.GetMapData(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "unknown", result.Weather.Status)
	assert.Equal(t, 0.0, result.Weather.Temperature)
	assert.Equal(t, "C", result.Weather.Unit)
	assert.Empty(t, result.Traffic)
	assert.Empty(t, result.Events)
}
