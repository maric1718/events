package repository

import (
	"encoding/json"
	"events_app/internal/core/domain"
	"os"
)

var CurrentMarketsState []domain.Market

type MarketAdapter struct{}

func NewMarketRepository() *MarketAdapter {
	return &MarketAdapter{}
}

// Init takes data from a file and populates CurrentMarketsState
func (mr *MarketAdapter) Init() error {
	file := filesPath + "markets.json"

	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, &CurrentMarketsState); err != nil {
		return err
	}

	return nil
}

func (mr *MarketAdapter) Get() ([]domain.Market, error) {
	payload := MarketFilterInactive(CurrentMarketsState)

	return payload, nil
}

func MarketFilterInactive(markets []domain.Market) []domain.Market {
	activeMarkets := []domain.Market{}

	for _, market := range markets {
		if market.Status == 1 {
			activeOutcomes := []domain.MarketOutcome{}

			for _, outcome := range market.Outcomes {
				if outcome.Status == 1 {
					activeOutcomes = append(activeOutcomes, outcome)
				}
			}

			market.Outcomes = activeOutcomes
			activeMarkets = append(activeMarkets, market)
		}
	}

	return activeMarkets
}
