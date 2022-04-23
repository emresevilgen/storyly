package event_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"storyly/pkg/errors"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	event_requests "storyly/pkg/model/requests/event-requests"
	event_service "storyly/pkg/service/event-service"
	"storyly/pkg/tracing"
)

type EventController struct {
	eventService event_service.EventService
	logFactory   log_factory.Factory
}

func NewEventController(eventService event_service.EventService) *EventController {
	return &EventController{
		eventService: eventService,
		logFactory:   log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(EventController{}), nil)),
	}
}

// PostEvent godoc
// @Summary Post event
// @Description Post event
// @ID post-event
// @Accept json
// @Produce json
// @Param token path string true "Token"
// @Param correlationId header string true "Correlation Id"
// @Param agentName header string true "Agent Name"
// @Param executorUser header string true "Executor User"
// @Success 200
// @Router /event/{token} [get]
func (c EventController) PostEvent(ctx echo.Context) error {
	var (
		err     error
		request event_requests.EventRequest
	)

	stdContext := tracing.CreateContextFromEcho(ctx)
	appId := ctx.Get("AppId").(int64)

	if err = ctx.Bind(&request); err != nil {
		c.logFactory.For(stdContext).Error("Binding error on post event for appId: %d, request: %+v, with error: %+v", appId, request, err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err = ctx.Validate(&request); err != nil {
		c.logFactory.For(stdContext).Error("Validation error on post event for appId: %d, request: %+v, with error: %+v", appId, request, err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err = c.eventService.PostEvent(stdContext, appId, request)
	if err != nil {
		errorResponse := err.(errors.ErrorResponse)
		c.logFactory.For(stdContext).Error("Post event is failed for appId: %d, request: %s, with error: %+v", appId, request, errorResponse)
		return ctx.JSON(errorResponse.StatusCode, errorResponse)
	}
	return ctx.NoContent(http.StatusOK)
}
