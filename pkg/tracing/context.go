package tracing

import (
	"context"
	"github.com/google/uuid"
)

type tracingContext struct {
	agentName     string
	executorUser  string
	correlationId string
	ctx           context.Context
}

func NewContext() *tracingContext {
	return &tracingContext{}
}

func NewContextFrom(ctx context.Context) *tracingContext {
	return &tracingContext{ctx: ctx}
}

func (c *tracingContext) SetDefaults() *tracingContext {
	c.correlationId = uuid.New().String()
	return c
}

func (c *tracingContext) WithAudit(auditInfo AuditInfo) *tracingContext {
	c.agentName = auditInfo.AgentName
	c.executorUser = auditInfo.ExecutorUser
	c.correlationId = auditInfo.CorrelationId
	return c
}

func (c *tracingContext) AgentName(agentName string) *tracingContext {
	c.agentName = agentName
	return c
}

func (c *tracingContext) ExecutorUser(executorUser string) *tracingContext {
	c.executorUser = executorUser
	return c
}

func (c *tracingContext) CorrelationId(correlationId string) *tracingContext {
	c.correlationId = correlationId
	return c
}

func (c *tracingContext) Build() context.Context {
	c.initIfNil()

	c.ctx = context.WithValue(c.ctx, AgentName, c.agentName)
	c.ctx = context.WithValue(c.ctx, CorrelationId, c.correlationId)
	c.ctx = context.WithValue(c.ctx, ExecutorUser, c.executorUser)

	return c.ctx
}

func (c *tracingContext) initIfNil() {
	if c.ctx == nil {
		c.ctx = context.Background()
	}
}
