package port

import (
	"events_app/internal/core/domain"
)

type EventService interface {
	Get(filters domain.EventFilters) ([]domain.Event, error)
}

type EventRepository interface {
	Get(filters domain.EventFilters) ([]domain.Event, error)
}
