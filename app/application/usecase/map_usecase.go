package usecase

import (
	"context"

	"github.com/brunobotter/map/application/domain"
	"github.com/brunobotter/map/application/service"
)

type MapUsecase interface {
	Execute(ctx context.Context) (*domain.MapData, error)
}

type mapUsecase struct {
	mapService service.MapService
}

func NewMapUsecase(mapService service.MapService) MapUsecase {
	return &mapUsecase{mapService: mapService}
}

func (u *mapUsecase) Execute(ctx context.Context) (*domain.MapData, error) {
	return u.mapService.GetMapData(ctx)
}
