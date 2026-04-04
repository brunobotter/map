package providers

import (
	"github.com/brunobotter/map/application/repo"
	infraRepo "github.com/brunobotter/map/infra/repo"
	"github.com/brunobotter/map/main/container"
)

type RepositoryProvider struct{}

func NewRepositoryProvider() *RepositoryProvider {
	return &RepositoryProvider{}
}
func (p *RepositoryProvider) Register(c container.Container) {
	c.Singleton(func() repo.MapRepository {
		return infraRepo.NewMapRepository()
	})
}
