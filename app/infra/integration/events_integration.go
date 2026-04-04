package integration

import (
	"context"

	"github.com/brunobotter/map/application/domain"
	appIntegration "github.com/brunobotter/map/application/integration"
)

type eventsIntegration struct{}

func NewEventsIntegration() appIntegration.EventsIntegration {
	return &eventsIntegration{}
}

func (s *eventsIntegration) GetEvents(ctx context.Context, lat, lng float64) ([]domain.Event, error) {
	_ = ctx

	return []domain.Event{
		{Name: "Feira de Saúde", Lat: lat + 0.0101, Lng: lng + 0.0072},
		{Name: "Mutirão de Vacinação", Lat: lat - 0.0064, Lng: lng - 0.0089},
	}, nil
}
