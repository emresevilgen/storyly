package tracing

import (
	"context"
	"github.com/labstack/echo/v4"
)

const (
	CorrelationId = "correlationId"
	AgentName     = "agentName"
	ExecutorUser  = "executorUser"
)

var DefaultClientHeaders = []string{
	CorrelationId,
	ExecutorUser,
}

var LoggableHeaders = []string{
	"X-Forwarded-For",
	"X-Ty-Ip",
	"Ty-Client-Ip",
	"True-Client-Ip",
	"User-Agent",
	"Sec-Fetch-Mode",
	"Origin",
	"Referer",
	"Cf-Connecting-Ip",
	"Cf-Ipcountry",
	"Cf-Request-Id",
}

func CreateContextFromEcho(ctx echo.Context) context.Context {
	stdContext := context.Background()
	stdContext = context.WithValue(stdContext, CorrelationId, ctx.Get(CorrelationId))
	stdContext = context.WithValue(stdContext, AgentName, ctx.Get(AgentName))
	stdContext = context.WithValue(stdContext, ExecutorUser, ctx.Get(ExecutorUser))
	return stdContext
}
