package postgresql_repository

import (
	"context"
	"database/sql"
	"reflect"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/model/documents/event_documents"
)

type EventRepository interface {
	InsertEvent(ctx context.Context, event event_documents.EventDocument) error
}

type eventRepository struct {
	logFactory log_factory.Factory
	dbCluster  *sql.DB
}

func NewEventRepository(dbCluster *sql.DB) *eventRepository {
	return &eventRepository{
		logFactory: log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(eventRepository{}), nil)),
		dbCluster:  dbCluster,
	}
}

func (r *eventRepository) InsertEvent(ctx context.Context, event event_documents.EventDocument) error {
	return nil
}
