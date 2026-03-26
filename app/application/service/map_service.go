package service

import (
	"context"

	"github.com/brunobotter/map/application/domain"
)

type MapService interface {
	GetMapData(ctx context.Context) (*domain.MapData, error)
}

type mapService struct{}

func NewMapService() MapService {
	return &mapService{}
}

func (s *mapService) GetMapData(ctx context.Context) (*domain.MapData, error) {
	_ = ctx

	return &domain.MapData{
		Weather: domain.Weather{
			Status:      "sunny",
			Temperature: 27.5,
			Unit:        "C",
		},
		Traffic: []domain.Traffic{
			{Road: "Avenida Central", Level: "moderate", Status: "slow"},
			{Road: "Rua das Flores", Level: "light", Status: "flowing"},
		},
		Events: []domain.MapEvent{
			{Title: "Feira de Saúde", Location: "Praça Principal", StartAt: "2026-03-30T09:00:00Z"},
			{Title: "Mutirão de Vacinação", Location: "Clínica Central", StartAt: "2026-04-02T08:00:00Z"},
		},
	}, nil
}
