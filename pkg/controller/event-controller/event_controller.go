package event_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	event_service "storyly/pkg/service/event-service"
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
// @Success 200 {object} meal_package.Package
// @Router /event/{token} [get]
func (c EventController) PostEvent(ctx echo.Context) error {
	//var (
	//	err      error
	//)

	//stdContext := tracing.CreateContextFromEcho(ctx)
	//token := ctx.Param("token")

	//response, err = c.packageService.Get(stdContext, packageId)
	//if err != nil {
	//	errorResponse := err.(errors.ErrorResponse)
	//	c.logFactory.For(stdContext).Error("Get package by Id is failed for packageId: %s, with error: %+v", packageId, errorResponse)
	//	new_relic.NoticeError(stdContext, errorResponse)
	//	return ctx.JSON(errorResponse.StatusCode, errorResponse)
	//}
	return ctx.JSON(http.StatusOK, nil)
}
