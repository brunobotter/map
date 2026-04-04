package repo

import "github.com/brunobotter/map/application/domain"

type MapRepository interface {
	GetEvents() ([]domain.MapEvent, error)
}
