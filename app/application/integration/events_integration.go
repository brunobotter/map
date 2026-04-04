package integration

import (
	"context"

	"github.com/brunobotter/map/application/domain"
)

type EventsIntegration interface {
	GetEvents(ctx context.Context, lat, lng float64) ([]domain.Event, error)
}
