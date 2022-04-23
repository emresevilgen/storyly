package auth_service

import (
	"context"
	"reflect"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/repository/postgresql-repository"
	inmemory_cache_service "storyly/pkg/service/inmemory-cache-service"
	"time"
)

const tokenCacheExpDuration = time.Duration(15) * time.Minute

type AuthService interface {
	GetAppId(ctx context.Context, token string) (int64, error)
}

type authService struct {
	logFactory      log_factory.Factory
	tokenRepository postgresql_repository.TokenRepository
	cacheService    inmemory_cache_service.Service
}

func NewAuthService(tokenRepository postgresql_repository.TokenRepository, cacheService inmemory_cache_service.Service) *authService {
	return &authService{
		logFactory:      log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(authService{}), nil)),
		tokenRepository: tokenRepository,
		cacheService:    cacheService,
	}
}

func (s *authService) GetAppId(ctx context.Context, token string) (int64, error) {
	if cacheRes, found := s.cacheService.Get(token); found {
		return cacheRes.(int64), nil
	}

	appId := int64(12)

	s.cacheService.Set(token, appId, tokenCacheExpDuration)

	return appId, nil
}
