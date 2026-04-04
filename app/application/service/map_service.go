package service

import (
	"context"

	"github.com/brunobotter/map/application/domain"
)

type MapService interface {
	GetMapData(ctx context.Context) (*domain.MapData, error)
}

type mapService struct {
	weatherService WeatherService
	trafficService TrafficService
	eventsService  EventsService
}

func NewMapService(weatherService WeatherService, trafficService TrafficService, eventsService EventsService) MapService {
	return &mapService{weatherService: weatherService, trafficService: trafficService, eventsService: eventsService}
}

func (s *mapService) GetMapData(ctx context.Context) (*domain.MapData, error) {
	const (
		lat = -23.55052
		lng = -46.633308
	)
	weather, err := s.weatherService.GetWeather(ctx, lat, lng)
	if err != nil {
		weather = domain.Weather{Status: "unknown", Temperature: 0, Unit: "C"}
	}

	traffic, err := s.trafficService.GetIncidents(ctx, -23.55052, -46.633308)
	if err != nil {
		traffic = []domain.Traffic{}
	}

	events, err := s.eventsService.GetEvents(ctx, lat, lng)
	if err != nil {
		events = []domain.Event{}
	}

	return &domain.MapData{
		Weather: weather,
		Traffic: traffic,
		Events:  events,
	}, nil
}
