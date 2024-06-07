package service

import (
	"events_app/internal/core/domain"
	"events_app/internal/core/port"
)

type MarketService struct {
	repo port.MarketRepository
}

func NewMarketService(repo port.MarketRepository) *MarketService {
	return &MarketService{
		repo,
	}
}

func (ms *MarketService) Get() ([]domain.Market, error) {
	markets, err := ms.repo.Get()
	if err != nil {
		return markets, err
	}

	return markets, nil
}

func (ms *MarketService) Send() {}
