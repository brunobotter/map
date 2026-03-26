package providers

import (
	"github.com/brunobotter/map/application/service"
	"github.com/brunobotter/map/application/usecase"
	"github.com/brunobotter/map/main/container"
)

type UseCaseProvider struct{}

func NewUseCaseProvider() *UseCaseProvider {
	return &UseCaseProvider{}
}
func (p *UseCaseProvider) Register(c container.Container) {
	c.Singleton(func(mapService service.MapService) usecase.MapUsecase {
		return usecase.NewMapUsecase(mapService)
	})
}
