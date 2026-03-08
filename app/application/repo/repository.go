package repo

import "github.com/brunobotter/map/application/domain"

type EventRepository interface {
	Create(event *domain.Event) error

	FindEventsInArea(
		north float64,
		south float64,
		east float64,
		west float64,
	) ([]domain.Event, error)
}
