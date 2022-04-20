package event_service

import (
	"context"
	"reflect"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
)

type EventService interface {
	PostEvent(ctx context.Context, appId int64) error
}

type eventService struct {
	logFactory log_factory.Factory
}

func NewEventService() *eventService {
	return &eventService{
		logFactory: log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(eventService{}), nil)),
	}
}

func (r *eventService) PostEvent(ctx context.Context, appId int64) error {
	return nil
}
