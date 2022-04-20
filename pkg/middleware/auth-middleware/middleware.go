package auth_middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	commonErrors "storyly/pkg/errors"
	log "storyly/pkg/log"
	"storyly/pkg/log/log_factory"
	"storyly/pkg/service/auth-service"
	"storyly/pkg/tracing"
	stringUtil "storyly/pkg/utils/string-util"
)

const (
	paramTokenKey = "token"
	appIdKey      = "AppId"

	tokenMissingMessage = "Token missing"
)

type Auth struct {
	authService auth_service.AuthService
	logger      log_factory.Factory
}

func NewAuthMiddleware(authService auth_service.AuthService) *Auth {
	return &Auth{
		authService: authService,
		logger:      log_factory.NewFactory(log.NewLoggerByName("Authorization-Middleware", nil)),
	}
}

func (auth *Auth) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Param(paramTokenKey)
		if stringUtil.IsBlank(token) {
			return ctx.JSON(http.StatusBadRequest, commonErrors.CreateError(http.StatusBadRequest, tokenMissingMessage))
		}

		stdContext := tracing.CreateContextFromEcho(ctx)
		authorizedAppId, err := auth.getAuthorizedAppId(stdContext, token)
		if err != nil {
			auth.logger.For(stdContext).Error("Authorization failed err: %+v", err)

			errResponse := err.(commonErrors.ErrorResponse)
			return ctx.JSON(errResponse.StatusCode, errResponse)
		}

		ctx.Set(appIdKey, authorizedAppId)
		return next(ctx)
	}
}

func (auth *Auth) getAuthorizedAppId(ctx context.Context, token string) (int64, error) {
	appId, err := auth.authService.GetAppId(ctx, token)
	if err != nil {
		return 0, err
	}

	return appId, nil
}
