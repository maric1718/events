package port

import (
	"events_app/internal/core/domain"
)

type MarketService interface {
	Get() ([]domain.Market, error)
}

type MarketRepository interface {
	Get() ([]domain.Market, error)
}
