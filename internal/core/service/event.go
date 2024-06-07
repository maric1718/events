package service

import (
	"events_app/internal/core/domain"
	"events_app/internal/core/port"
)

type EventService struct {
	repo port.EventRepository
}

func NewEventService(repo port.EventRepository) *EventService {
	return &EventService{
		repo,
	}
}

func (ms *EventService) Get(filters domain.EventFilters) ([]domain.Event, error) {
	events, err := ms.repo.Get(filters)
	if err != nil {
		return events, err
	}

	return events, nil
}
