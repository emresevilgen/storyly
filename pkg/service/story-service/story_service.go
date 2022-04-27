package story_service

import (
	"context"
	"encoding/json"
	"reflect"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/model/domain"
	"storyly/pkg/model/responses/story-responses"
	"storyly/pkg/repository/postgresql-repository"
	inmemory_cache_service "storyly/pkg/service/inmemory-cache-service"
	"strconv"
	"time"
)

const storiesCacheExpDuration = time.Duration(1) * time.Second

type StoryService interface {
	GetStories(ctx context.Context, appId int64) (story_responses.StoryListResponse, error)
	ValidateStoryAndApp(ctx context.Context, appId int64, storyId int64) (bool, error)
}

type storyService struct {
	logFactory      log_factory.Factory
	storyRepository postgresql_repository.StoryRepository
	cacheService    inmemory_cache_service.Service
}

func NewStoryService(storyRepository postgresql_repository.StoryRepository, cacheService inmemory_cache_service.Service) *storyService {
	return &storyService{
		logFactory:      log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(storyService{}), nil)),
		storyRepository: storyRepository,
		cacheService:    cacheService,
	}
}

func (s *storyService) GetStories(ctx context.Context, appId int64) (story_responses.StoryListResponse, error) {
	appIdStr := strconv.FormatInt(appId, 10)
	if cacheRes, found := s.cacheService.Get(appIdStr); found {
		return prepareStoryListResponse(appId, cacheRes.([]story_responses.StoryResponse)), nil
	}

	stories, err := s.storyRepository.GetStories(ctx, appId)
	if err != nil {
		return story_responses.StoryListResponse{}, err
	}

	storiesResp := make([]story_responses.StoryResponse, 0)

	for _, story := range stories {
		var metadata domain.Metadata
		err = json.Unmarshal([]byte(story.Metadata), &metadata)
		if err != nil {
			return story_responses.StoryListResponse{}, err
		}

		storiesResp = append(storiesResp, story_responses.StoryResponse{
			Id:       story.StoryId,
			Metadata: metadata.Image,
		})
	}

	resp := prepareStoryListResponse(appId, storiesResp)

	s.cacheService.Set(appIdStr, storiesResp, storiesCacheExpDuration)

	return resp, nil
}

func (s *storyService) ValidateStoryAndApp(ctx context.Context, appId int64, storyId int64) (bool, error) {
	story, err := s.storyRepository.GetStory(ctx, storyId)
	if err != nil {
		return false, err
	}

	return appId == story.AppId, nil
}

func prepareStoryListResponse(appId int64, storiesResp []story_responses.StoryResponse) story_responses.StoryListResponse {
	return story_responses.StoryListResponse{
		AppId:     appId,
		Timestamp: time.Now().Unix(),
		Metadata:  storiesResp,
	}
}
