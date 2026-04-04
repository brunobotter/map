package service

import (
	"context"

	"github.com/brunobotter/map/application/domain"
	"github.com/brunobotter/map/application/integration"
)

type EventsService interface {
	GetEvents(ctx context.Context, lat, lng float64) ([]domain.Event, error)
}

type eventsService struct {
	eventsIntegration integration.EventsIntegration
}

func NewEventsService(eventsIntegration integration.EventsIntegration) EventsService {
	return &eventsService{eventsIntegration: eventsIntegration}
}

func (s *eventsService) GetEvents(ctx context.Context, lat, lng float64) ([]domain.Event, error) {
	return s.eventsIntegration.GetEvents(ctx, lat, lng)
}
