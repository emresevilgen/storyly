package tracing_middleware

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"storyly/pkg/errors"
	"storyly/pkg/tracing"
	stringUtil "storyly/pkg/utils/string-util"
)

type tracingMiddleware struct {
	validateCorrelationId bool
	validateAgentName     bool
	validateExecutorUser  bool
}

func NewTracingMiddleware(validateCorrelationId, validateAgentName, validateExecutorUser bool) *tracingMiddleware {
	return &tracingMiddleware{
		validateCorrelationId,
		validateAgentName,
		validateExecutorUser,
	}
}

func (m *tracingMiddleware) GetTracingInformationFromEcho(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		correlationId := ctx.Request().Header.Get(tracing.CorrelationId)
		if stringUtil.IsBlank(correlationId) {
			if m.validateCorrelationId {
				return ctx.JSON(http.StatusBadRequest, errors.CreateError(http.StatusBadRequest, getRequiredError(tracing.CorrelationId)))
			}
			correlationId = uuid.New().String()
		}

		agentName := ctx.Request().Header.Get(tracing.AgentName)
		if m.validateAgentName && stringUtil.IsBlank(agentName) {
			return ctx.JSON(http.StatusBadRequest, errors.CreateError(http.StatusBadRequest, getRequiredError(tracing.AgentName)))
		}

		executorUser := ctx.Request().Header.Get(tracing.ExecutorUser)
		if m.validateExecutorUser && stringUtil.IsBlank(executorUser) {
			return ctx.JSON(http.StatusBadRequest, errors.CreateError(http.StatusBadRequest, getRequiredError(tracing.ExecutorUser)))
		}

		ctx.Set(tracing.AgentName, agentName)
		ctx.Set(tracing.CorrelationId, correlationId)
		ctx.Set(tracing.ExecutorUser, executorUser)
		return next(ctx)
	}
}

func getRequiredError(headerKey string) string {
	return fmt.Sprintf("%s header is required", headerKey)
}
