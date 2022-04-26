package postgresql_repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"reflect"
	commonErrors "storyly/pkg/errors"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/model/documents/stories"
)

type StoryRepository interface {
	GetStories(ctx context.Context, appId int64) ([]stories.Story, error)
	GetStory(ctx context.Context, storyId int64) (stories.Story, error)
}

type storyRepository struct {
	logFactory log_factory.Factory
	dbCluster  *gorm.DB
}

func NewStoryRepository(dbCluster *gorm.DB) *storyRepository {
	return &storyRepository{
		logFactory: log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(storyRepository{}), nil)),
		dbCluster:  dbCluster,
	}
}

func (r *storyRepository) GetStories(ctx context.Context, appId int64) ([]stories.Story, error) {
	var storyList = make([]stories.Story, 0)
	r.logFactory.For(ctx).Info("Starting to query storyList")

	result := r.dbCluster.Where(&stories.Story{AppId: appId}).Find(&storyList)
	if result.Error != nil {
		return storyList, commonErrors.CreateDbError(result.Error)
	}

	r.logFactory.For(ctx).Info("Finished gathering storyList")
	return storyList, nil
}

func (r *storyRepository) GetStory(ctx context.Context, storyId int64) (stories.Story, error) {
	var story stories.Story
	r.logFactory.For(ctx).Info("Starting to get story")

	result := r.dbCluster.First(&story, storyId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return stories.Story{}, nil
		}
		return stories.Story{}, commonErrors.CreateDbError(result.Error)
	}

	r.logFactory.For(ctx).Info("Finished get stoty")
	return story, nil
}
