package providers

import (
	"github.com/brunobotter/map/application/service"
	"github.com/brunobotter/map/main/container"
)

type ServiceProvider struct{}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
func (p *ServiceProvider) Register(c container.Container) {
	c.Singleton(func() service.MapService {
		return service.NewMapService()
	})
}
