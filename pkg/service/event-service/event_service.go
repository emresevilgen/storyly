package event_service

import (
	"context"
	"reflect"
	commonErrors "storyly/pkg/errors"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/model/documents/events"
	"storyly/pkg/model/requests/event-requests"
	"storyly/pkg/model/responses/event-responses"
	"storyly/pkg/repository/postgresql-repository"
	"storyly/pkg/service/story-service"
	"time"
)

type EventService interface {
	PostEvent(ctx context.Context, appId int64, request event_requests.EventRequest) error
	GetEventMetrics(ctx context.Context, appId int64, date string) (event_responses.EventMetricsResponse, error)
}
type eventService struct {
	logFactory      log_factory.Factory
	eventRepository postgresql_repository.EventRepository
	storyService    story_service.StoryService
}

func NewEventService(eventRepository postgresql_repository.EventRepository, storyService story_service.StoryService) *eventService {
	return &eventService{
		logFactory:      log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(eventService{}), nil)),
		eventRepository: eventRepository,
		storyService:    storyService,
	}
}

func (s *eventService) PostEvent(ctx context.Context, appId int64, request event_requests.EventRequest) error {
	isValid, err := s.storyService.ValidateStoryAndApp(ctx, appId, request.StoryId)
	if err != nil {
		return err
	}

	if !isValid {
		return commonErrors.CreateAuthenticationError()
	}

	event := events.Event{
		AppId:     appId,
		StoryId:   request.StoryId,
		EventType: string(request.EventType),
		UserId:    request.UserId,
		Date:      time.Now().Format("2006-01-02"),
	}

	eventFromRepo, err := s.eventRepository.GetEvent(ctx, event)
	if err != nil {
		return err
	}

	if eventFromRepo.IsMissing() {
		event.Count = 1
		return s.eventRepository.InsertEvent(ctx, event)
	}

	eventFromRepo.Count += 1
	return s.eventRepository.UpdateEventCount(ctx, eventFromRepo)
}

func (s *eventService) GetEventMetrics(ctx context.Context, appId int64, date string) (event_responses.EventMetricsResponse, error) {
	count, err := s.eventRepository.GetCountByAppIdAndDate(ctx, appId, date)
	if err != nil {
		return event_responses.EventMetricsResponse{}, err
	}

	return event_responses.EventMetricsResponse{
		AppId:            appId,
		DailyActiveUsers: count,
	}, nil
}
