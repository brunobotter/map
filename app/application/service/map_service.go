package service

import (
	"context"

	"github.com/brunobotter/map/application/domain"
	"github.com/brunobotter/map/application/repo"
)

type MapService interface {
	GetMapData(ctx context.Context) (*domain.MapData, error)
}

type mapService struct {
	weatherService WeatherService
	trafficService TrafficService
	mapRepository  repo.MapRepository
}

func NewMapService(weatherService WeatherService, trafficService TrafficService, mapRepository repo.MapRepository) MapService {
	return &mapService{weatherService: weatherService, trafficService: trafficService, mapRepository: mapRepository}
}

func (s *mapService) GetMapData(ctx context.Context) (*domain.MapData, error) {
	weather, err := s.weatherService.GetWeather(ctx, -23.55052, -46.633308)
	if err != nil {
		weather = domain.Weather{Status: "unknown", Temperature: 0, Unit: "C"}
	}

	traffic, err := s.trafficService.GetIncidents(ctx, -23.55052, -46.633308)
	if err != nil {
		traffic = []domain.Traffic{}
	}

	events, err := s.mapRepository.GetEvents()
	if err != nil {
		events = []domain.MapEvent{}
	}

	return &domain.MapData{
		Weather: weather,
		Traffic: traffic,
		Events:  events,
	}, nil
}
