package auth_service

import (
	"context"
	"reflect"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/repository/postgresql-repository"
)

type AuthService interface {
	GetAppId(ctx context.Context, token string) (int64, error)
}

type authService struct {
	logFactory      log_factory.Factory
	tokenRepository postgresql_repository.TokenRepository
}

func NewAuthService(tokenRepository postgresql_repository.TokenRepository) *authService {
	return &authService{
		logFactory:      log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(authService{}), nil)),
		tokenRepository: tokenRepository,
	}
}

func (r *authService) GetAppId(ctx context.Context, token string) (int64, error) {
	appId := int64(12)

	return appId, nil
}
