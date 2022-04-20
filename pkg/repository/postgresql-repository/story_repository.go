package postgresql_repository

import (
	"context"
	"reflect"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/model/documents/story_documents"
)

type StoryRepository interface {
	GetStories(ctx context.Context, appId int64) ([]story_documents.StoryDocument, error)
}

type storyRepository struct {
	logFactory log_factory.Factory
}

func NewStoryRepository() *storyRepository {
	return &storyRepository{
		logFactory: log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(storyRepository{}), nil)),
	}
}

func (r *storyRepository) GetStories(ctx context.Context, appId int64) ([]story_documents.StoryDocument, error) {
	return []story_documents.StoryDocument{
		{
			AppId:    appId,
			Id:       2,
			Metadata: story_documents.MetadataDocument{Image: "image1.png"},
		},
	}, nil
}
