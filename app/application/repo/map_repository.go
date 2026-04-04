package repo

import "github.com/brunobotter/map/application/domain"

type MapRepository interface {
	GetTraffic() ([]domain.Traffic, error)
	GetEvents() ([]domain.MapEvent, error)
}
