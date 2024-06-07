package util

import "events_app/internal/core/domain"

func RemoveEventFromSlice(s []domain.Event, i int) []domain.Event {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func RemoveMarketFromSlice(s []domain.Market, i int) []domain.Market {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
