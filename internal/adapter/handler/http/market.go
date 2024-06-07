package http

import (
	"events_app/internal/core/port"

	"github.com/gin-gonic/gin"
)

type MarketHandler struct {
	svc port.MarketService
}

func NewMarketHandler(svc port.MarketService) *MarketHandler {
	return &MarketHandler{
		svc,
	}
}

func (mh *MarketHandler) GetMarkets(c *gin.Context) {
	data, err := mh.svc.Get()
	if err != nil {
		ThrowStatusInternalServerError(err.Error(), c)
		return
	}

	ThrowStatusOk(data, c)
}
