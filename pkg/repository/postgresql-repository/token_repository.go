package postgresql_repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"reflect"
	commonErrors "storyly/pkg/errors"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/model/documents/tokens"
)

type TokenRepository interface {
	GetAppId(ctx context.Context, token string) (int64, error)
}

type tokenRepository struct {
	logFactory log_factory.Factory
	dbCluster  *gorm.DB
}

func NewTokenRepository(dbCluster *gorm.DB) *tokenRepository {
	return &tokenRepository{
		logFactory: log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(tokenRepository{}), nil)),
		dbCluster:  dbCluster,
	}
}

func (r *tokenRepository) GetAppId(ctx context.Context, token string) (int64, error) {
	r.logFactory.For(ctx).Info("Starting to query token")

	var tokenDoc tokens.Token

	result := r.dbCluster.First(&tokenDoc, "token = ?", token)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, commonErrors.CreateAuthenticationError()
		}

		return 0, commonErrors.CreateDbError(result.Error)
	}

	r.logFactory.For(ctx).Info("Finished gathering token")
	return tokenDoc.Id, nil
}
