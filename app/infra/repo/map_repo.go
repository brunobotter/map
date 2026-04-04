package repo

import (
	"github.com/brunobotter/map/application/domain"
	appRepo "github.com/brunobotter/map/application/repo"
)

type mapRepository struct{}

func NewMapRepository() appRepo.MapRepository {
	return &mapRepository{}
}

func (r *mapRepository) GetTraffic() ([]domain.Traffic, error) {
	return []domain.Traffic{
		{Road: "Avenida Central", Level: "moderate", Status: "slow"},
		{Road: "Rua das Flores", Level: "light", Status: "flowing"},
	}, nil
}

func (r *mapRepository) GetEvents() ([]domain.MapEvent, error) {
	return []domain.MapEvent{
		{Title: "Feira de Saúde", Location: "Praça Principal", StartAt: "2026-03-30T09:00:00Z"},
		{Title: "Mutirão de Vacinação", Location: "Clínica Central", StartAt: "2026-04-02T08:00:00Z"},
	}, nil
}
