package service

import (
	"context"

	"github.com/brunobotter/map/application/domain"
)

type TrafficService interface {
	GetIncidents(ctx context.Context, lat, lng float64) ([]domain.Traffic, error)
}

type trafficService struct{}

func NewTrafficService() TrafficService {
	return &trafficService{}
}

func (s *trafficService) GetIncidents(ctx context.Context, lat, lng float64) ([]domain.Traffic, error) {
	_ = ctx
	return []domain.Traffic{
		{Lat: lat + 0.0015, Lng: lng + 0.0011, Type: "accident"},
		{Lat: lat - 0.0021, Lng: lng - 0.0014, Type: "jam"},
	}, nil
}
