package postgresql_repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"reflect"
	commonErrors "storyly/pkg/errors"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/model/documents/events"
)

type EventRepository interface {
	InsertEvent(ctx context.Context, event events.Event) error
	UpdateEventCount(ctx context.Context, event events.Event) error
	GetEvent(ctx context.Context, event events.Event) (events.Event, error)
	GetCountByAppIdAndDate(ctx context.Context, appId int64, date string) (int64, error)
}

type eventRepository struct {
	logFactory log_factory.Factory
	dbCluster  *gorm.DB
}

func NewEventRepository(dbCluster *gorm.DB) *eventRepository {
	return &eventRepository{
		logFactory: log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(eventRepository{}), nil)),
		dbCluster:  dbCluster,
	}
}

func (r *eventRepository) InsertEvent(ctx context.Context, event events.Event) error {
	r.logFactory.For(ctx).Info("Starting to insert event")

	result := r.dbCluster.Create(&event)
	if result.Error != nil {
		return commonErrors.CreateDbError(result.Error)
	}

	r.logFactory.For(ctx).Info("Finished inserting event")
	return nil
}

func (r *eventRepository) UpdateEventCount(ctx context.Context, event events.Event) error {
	r.logFactory.For(ctx).Info("Starting to update event")

	result := r.dbCluster.First(&events.Event{}, "story_id = ? AND app_id = ? AND event_type = ? AND user_id = ? AND date = ?",
		event.StoryId, event.AppId, event.EventType, event.UserId, event.Date).Updates(&event)
	if result.Error != nil {
		return commonErrors.CreateDbError(result.Error)
	}

	r.logFactory.For(ctx).Info("Finished updating event")
	return nil
}

func (r *eventRepository) GetEvent(ctx context.Context, reqEvent events.Event) (events.Event, error) {
	r.logFactory.For(ctx).Info("Starting to query reqEvent")

	var event events.Event
	result := r.dbCluster.First(&event, "story_id = ? AND app_id = ? AND event_type = ? AND user_id = ? AND date = ?",
		reqEvent.StoryId, reqEvent.AppId, reqEvent.EventType, reqEvent.UserId, reqEvent.Date)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return events.Event{}, nil
		}

		return events.Event{}, commonErrors.CreateDbError(result.Error)
	}

	r.logFactory.For(ctx).Info("Finished gathering reqEvent")
	return event, nil
}

func (r *eventRepository) GetCountByAppIdAndDate(ctx context.Context, appId int64, date string) (int64, error) {
	r.logFactory.For(ctx).Info("Starting to query event count by appId and date")

	var count int64
	result := r.dbCluster.Model(&events.Event{}).Where("app_id=? AND date=?", appId, date).Distinct("user_id").Count(&count)
	if result.Error != nil {
		return 0, commonErrors.CreateDbError(result.Error)
	}

	r.logFactory.For(ctx).Info("Finished gathering event count by appId and date")
	return count, nil
}
