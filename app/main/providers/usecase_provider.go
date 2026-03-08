package providers

import (
	"github.com/brunobotter/map/main/container"
)

type UseCaseProvider struct{}

func NewUseCaseProvider() *UseCaseProvider {
	return &UseCaseProvider{}
}
func (p *UseCaseProvider) Register(c container.Container) {

}
