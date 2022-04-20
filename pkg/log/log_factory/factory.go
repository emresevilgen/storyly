package log_factory

import (
	"context"
	log "storyly/pkg/log"
)

// Factory is the default logging wrapper that can create
// logger instances either for a given Context or context-less.
type Factory struct {
	logger *log.Logger
}

func (f Factory) IsDebugEnabled() bool {
	return f.logger.IsDebugEnabled()
}

// Bg creates a context-unaware logger.
func (f Factory) Bg() Logger {
	return logger{logger: f.logger}
}

// NewFactory creates a new Factory.
func NewFactory(logger *log.Logger) Factory {
	return Factory{logger: logger}
}

func (f Factory) For(ctx context.Context) Logger {
	return ctxLogger{ctx: ctx, logger: f.logger}
}
