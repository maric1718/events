package repository

import (
	"encoding/json"
	"events_app/internal/core/domain"
	"os"
	"time"
)

var CurrentEventsState []domain.Event

type EventAdapter struct{}

func NewEventRepository() *EventAdapter {
	return &EventAdapter{}
}

// Init takes data from a file and populates CurrentEventsState
func (ea *EventAdapter) Init() error {
	file := filesPath + "events.json"

	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, &CurrentEventsState); err != nil {
		return err
	}

	return nil
}

func (ea *EventAdapter) Get(filters domain.EventFilters) ([]domain.Event, error) {
	payload := EventFilterInactive(CurrentEventsState)

	payload, err := FilterEvents(payload, filters)
	if err != nil {
		return payload, err
	}

	return payload, nil
}

func FilterEvents(payload []domain.Event, filters domain.EventFilters) ([]domain.Event, error) {
	filteredPayload := []domain.Event{}

	var from, to time.Time
	var err error

	if filters.From != "" {
		from, err = time.Parse(domain.TimeFormat, filters.From)
		if err != nil {
			return payload, err
		}
	}

	if filters.To != "" {
		to, err = time.Parse(domain.TimeFormat, filters.To)
		if err != nil {
			return payload, err
		}
	}

	if filters.From != "" && filters.To != "" {
		for _, event := range CurrentEventsState {
			if event.StartsAt.After(from) && event.StartsAt.Before(to) {
				filteredPayload = append(filteredPayload, event)
			}
		}

		payload = filteredPayload

	} else if filters.From != "" {

		for _, event := range CurrentEventsState {
			if event.StartsAt.After(from) {
				filteredPayload = append(filteredPayload, event)
			}
		}

		payload = filteredPayload

	} else if filters.To != "" {
		for _, event := range CurrentEventsState {
			if event.StartsAt.Before(to) {
				filteredPayload = append(filteredPayload, event)
			}
		}

		payload = filteredPayload
	}

	return payload, nil
}

func EventFilterInactive(events []domain.Event) []domain.Event {
	activeEvents := []domain.Event{}

	for _, event := range events {
		if event.Status == 1 {
			activeMarkets := []domain.EventMarket{}

			for _, market := range event.Markets {
				if market.Status == 1 {
					activeOutcomes := []domain.EventMarketOutcome{}

					for _, outcome := range market.Outcomes {
						if outcome.Status == 1 {
							activeOutcomes = append(activeOutcomes, outcome)
						}
					}

					market.Outcomes = activeOutcomes
					activeMarkets = append(activeMarkets, market)
				}
			}

			event.Markets = activeMarkets
			activeEvents = append(activeEvents, event)
		}
	}

	return activeEvents
}
