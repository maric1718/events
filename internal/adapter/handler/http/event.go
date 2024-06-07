package http

import (
	"events_app/internal/core/domain"
	"events_app/internal/core/port"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type EventHandler struct {
	svc port.EventService
}

func NewEventHandler(svc port.EventService) *EventHandler {
	return &EventHandler{
		svc,
	}
}

func (eh *EventHandler) GetEvents(c *gin.Context) {
	// filters
	filters := domain.EventFilters{}

	if err := c.ShouldBindWith(&filters, binding.Form); err != nil {
		return
	}

	data, err := eh.svc.Get(filters)
	if err != nil {
		ThrowStatusInternalServerError(err.Error(), c)
		return
	}

	ThrowStatusOk(data, c)
}
