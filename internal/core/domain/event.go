package domain

import (
	"encoding/json"
	"log/slog"
	"time"
)

const (
	TimeFormat = "2006-01-02T15:04:05"
)

type Event struct {
	ID       string        `json:"ID"`
	Name     string        `json:"name"`
	StartsAt time.Time     `json:"startsAt"`
	Status   int           `json:"status"`
	Markets  []EventMarket `json:"markets"`
}

type EventMarket struct {
	ID       string               `json:"id"`
	MarketID string               `json:"marketId"`
	Status   int                  `json:"status"`
	Outcomes []EventMarketOutcome `json:"outcomes"`
}

type EventMarketOutcome struct {
	ID        string  `json:"id"`
	OutcomeID string  `json:"outcomeId"`
	Status    int     `json:"status"`
	Odds      float32 `json:"odds"`
}

type EventFilters struct {
	From string `form:"from"`
	To   string `form:"to"`
}

func (e *Event) UnmarshalJSON(data []byte) (err error) {
	type Alias Event

	aux := &struct {
		StartsAt string `json:"startsAt"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.StartsAt != "" {

		s, err := time.Parse(TimeFormat, aux.StartsAt)
		if err != nil {
			slog.Error("Error parsing time: %v", "error", err)
			return err
		}

		e.StartsAt = s
	}

	return nil
}
