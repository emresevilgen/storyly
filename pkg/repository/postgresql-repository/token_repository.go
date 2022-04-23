package postgresql_repository

import (
	"context"
	"database/sql"
	"reflect"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
)

type TokenRepository interface {
	GetAppId(ctx context.Context, token string) (int64, error)
}

type tokenRepository struct {
	logFactory log_factory.Factory
	dbCluster  *sql.DB
}

func NewTokenRepository(dbCluster *sql.DB) *tokenRepository {
	return &tokenRepository{
		logFactory: log_factory.NewFactory(log.NewLoggerByType(reflect.TypeOf(tokenRepository{}), nil)),
		dbCluster:  dbCluster,
	}
}

func (r *tokenRepository) GetAppId(ctx context.Context, token string) (int64, error) {
	return 1, nil
}
