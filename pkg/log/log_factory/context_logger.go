package log_factory

import (
	"context"
	"github.com/google/uuid"
	eventLogger "storyly/pkg/log"
)

const (
	correlationIdHeaderKey = "x-correlationid"
	agentNameHeaderKey     = "x-agentname"
	unknownAgent           = "unknown-agent"
	executorUserHeaderKey  = "x-executor-user"
	userIpKey              = "x-ty-ip"
)

type ctxLogger struct {
	logger *eventLogger.Logger
	ctx    context.Context
}

func (cl ctxLogger) Info(msg string, args ...interface{}) {
	fields := cl.getFields(nil)
	cl.logger.InfoWithFields(fields, msg, args...)
}

func (cl ctxLogger) InfoWithFields(fields map[string]interface{}, msg string, args ...interface{}) {
	finalFields := cl.getFields(fields)
	cl.logger.InfoWithFields(finalFields, msg, args...)
}

func (cl ctxLogger) Error(msg string, args ...interface{}) {
	fields := cl.getFields(nil)
	cl.logger.ErrorWithFields(fields, msg, args...)
}

func (cl ctxLogger) ErrorWithFields(fields map[string]interface{}, msg string, args ...interface{}) {
	finalFields := cl.getFields(fields)
	cl.logger.ErrorWithFields(finalFields, msg, args...)
}

func (cl ctxLogger) Warn(msg string, args ...interface{}) {
	fields := cl.getFields(nil)
	cl.logger.WarnWithFields(fields, msg, args...)
}

func (cl ctxLogger) WarnWithFields(fields map[string]interface{}, msg string, args ...interface{}) {
	finalFields := cl.getFields(fields)
	cl.logger.WarnWithFields(finalFields, msg, args...)
}

func (cl ctxLogger) Debug(msg string, args ...interface{}) {
	fields := cl.getFields(nil)
	cl.logger.DebugWithFields(fields, msg, args...)
}

func (cl ctxLogger) DebugWithFields(fields map[string]interface{}, msg string, args ...interface{}) {
	finalFields := cl.getFields(fields)
	cl.logger.DebugWithFields(finalFields, msg, args...)
}

func (cl ctxLogger) getFields(fields map[string]interface{}) map[string]interface{} {
	finalFields := map[string]interface{}{
		correlationIdHeaderKey: cl.findCorrelationId(),
		agentNameHeaderKey:     cl.findAgentName(),
		executorUserHeaderKey:  cl.findExecutorUser(),
		userIpKey:              cl.findUserIp(),
	}

	for k, v := range fields {
		finalFields[k] = v
	}

	return finalFields
}

func (cl ctxLogger) findCorrelationId() string {
	correlationId := cl.ctx.Value(correlationIdHeaderKey)
	if correlationId == nil {
		return uuid.New().String()
	}
	return correlationId.(string)
}

func (cl ctxLogger) findAgentName() string {
	agentName := cl.ctx.Value(agentNameHeaderKey)
	if agentName == nil {
		return unknownAgent
	}
	return agentName.(string)
}

func (cl ctxLogger) findExecutorUser() string {
	executorUser := cl.ctx.Value(executorUserHeaderKey)
	if executorUser == nil {
		return ""
	}
	return executorUser.(string)
}

func (cl ctxLogger) findUserIp() string {
	userIp := cl.ctx.Value(userIpKey)
	if userIp == nil {
		return ""
	}
	return userIp.(string)
}
