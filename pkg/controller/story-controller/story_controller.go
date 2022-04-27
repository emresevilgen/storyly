package story_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"storyly/pkg/errors"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/model/responses/story-responses"
	"storyly/pkg/service/story-service"
	"storyly/pkg/tracing"
)

type StoryController struct {
	logFactory   log_factory.Factory
	storyService story_service.StoryService
}

func NewStoryController(storyService story_service.StoryService) *StoryController {
	return &StoryController{
		logFactory:   log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(StoryController{}), nil)),
		storyService: storyService,
	}
}

// GetStories godoc
// @Summary Retrieve stories by token
// @Description Retrieve stories by token
// @ID get-stories-by-token
// @Accept json
// @Produce json
// @Param token path string true "Token"
// @Param correlationId header string true "Correlation Id"
// @Param agentName header string true "Agent Name"
// @Param executorUser header string true "Executor User"
// @Success 200 {object} story_responses.StoryListResponse
// @Router /stories/{token} [get]
func (c StoryController) GetStories(ctx echo.Context) error {
	var (
		response story_responses.StoryListResponse
		err      error
	)

	appId := ctx.Get("AppId").(int64)
	stdContext := tracing.CreateContextFromEcho(ctx)

	response, err = c.storyService.GetStories(stdContext, appId)
	if err != nil {
		errorResponse := err.(errors.ErrorResponse)
		c.logFactory.For(stdContext).Error("Get stories failed for appId: %s, with error: %+v", appId, errorResponse)
		return ctx.JSON(errorResponse.StatusCode, errorResponse)
	}

	return ctx.JSON(http.StatusOK, response)
}
