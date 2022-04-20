package story_service

import (
	"context"
	"reflect"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/model/responses/story-responses"
	"storyly/pkg/repository/postgresql-repository"
	"time"
)

type StoryService interface {
	GetStories(ctx context.Context, appId int64) (story_responses.StoryListResponse, error)
}

type storyService struct {
	logFactory      log_factory.Factory
	storyRepository postgresql_repository.StoryRepository
}

func NewStoryService(storyRepository postgresql_repository.StoryRepository) *storyService {
	return &storyService{
		logFactory:      log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(storyService{}), nil)),
		storyRepository: storyRepository,
	}
}

func (r *storyService) GetStories(ctx context.Context, appId int64) (story_responses.StoryListResponse, error) {
	stories, err := r.storyRepository.GetStories(ctx, appId)
	if err != nil {
		return story_responses.StoryListResponse{}, err
	}

	storiesResp := make([]story_responses.StoryResponse, 0)

	for _, story := range stories {
		storiesResp = append(storiesResp, story_responses.StoryResponse{
			Id:       story.Id,
			Metadata: story.Metadata.Image,
		})
	}

	return story_responses.StoryListResponse{
		AppId:     appId,
		Timestamp: time.Now().UnixMilli(),
		Metadata:  storiesResp,
	}, nil
}
