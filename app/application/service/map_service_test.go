package service

import (
	"context"

	"github.com/brunobotter/map/application/domain"
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

type eventsServiceStub struct {
	result []domain.Event
	err    error
}

func (e *eventsServiceStub) GetEvents(ctx context.Context, lat, lng float64) ([]domain.Event, error) {
	_ = ctx
	_ = lat
	_ = lng
	if e.err != nil {
		return nil, e.err
	}
	return e.result, nil
}
