package event_service

import (
	"context"
	"reflect"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/model/documents/event_documents"
	event_requests "storyly/pkg/model/requests/event-requests"
	postgresql_repository "storyly/pkg/repository/postgresql-repository"
)

type EventService interface {
	PostEvent(ctx context.Context, appId int64, request event_requests.EventRequest) error
}

type eventService struct {
	logFactory      log_factory.Factory
	eventRepository postgresql_repository.EventRepository
}

func NewEventService(eventRepository postgresql_repository.EventRepository) *eventService {
	return &eventService{
		logFactory:      log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(eventService{}), nil)),
		eventRepository: eventRepository,
	}
}

func (s *eventService) PostEvent(ctx context.Context, appId int64, request event_requests.EventRequest) error {
	return s.eventRepository.InsertEvent(ctx, event_documents.EventDocument{
		AppId:     appId,
		StoryId:   request.StoryId,
		EventType: string(request.EventType),
		UserId:    request.UserId,
	})
}
