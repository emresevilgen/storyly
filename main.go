package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	inmemcache "github.com/patrickmn/go-cache"
	"github.com/swaggo/echo-swagger"
	"github.com/tylerb/graceful"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"storyly/configs"
	_ "storyly/docs"
	"storyly/pkg/controller/event-controller"
	"storyly/pkg/controller/story-controller"
	"storyly/pkg/middleware/auth-middleware"
	"storyly/pkg/middleware/tracing-middleware"
	"storyly/pkg/repository/postgresql-repository"
	"storyly/pkg/service/auth-service"
	"storyly/pkg/service/event-service"
	"storyly/pkg/service/inmemory-cache-service"
	"storyly/pkg/service/story-service"
	customValidator "storyly/pkg/validator"
	"time"
)

const (
	appPort = "8086"
)

func init() {
	configs.InitConfigs()
}

// @title Storyly
// @description This is for storyly assignment.
func main() {
	// Init db cluster
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=storyly sslmode=disable",
		configs.AppConfig.PostgreSql.Host, configs.AppConfig.PostgreSql.Port, configs.Secrets.PostgreSqlUser, configs.Secrets.PostgreSqlPassword)

	dbCluster, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Db connection error,err:", err.Error())
		return
	}

	db, err := dbCluster.DB()
	if err != nil {
		fmt.Println("Db error,err:", err.Error())
		return
	}

	defer func() {
		_ = db.Close()
	}()

	// Init in memory caches
	inMemoryCacheForAuth := inmemcache.New(9999*time.Hour, 9999*time.Hour)
	inMemoryCacheForStories := inmemcache.New(9999*time.Hour, 9999*time.Hour)

	// Init repositories
	tokenRepository := postgresql_repository.NewTokenRepository(dbCluster)
	storyRepository := postgresql_repository.NewStoryRepository(dbCluster)
	eventRepository := postgresql_repository.NewEventRepository(dbCluster)

	// Init services
	cacheServiceForAuth := inmemory_cache_service.New(inMemoryCacheForAuth)
	cacheServiceForStories := inmemory_cache_service.New(inMemoryCacheForStories)
	authService := auth_service.NewAuthService(tokenRepository, cacheServiceForAuth)
	storyService := story_service.NewStoryService(storyRepository, cacheServiceForStories)
	eventService := event_service.NewEventService(eventRepository, storyService)

	// Init controllers
	storyController := story_controller.NewStoryController(storyService)
	eventController := event_controller.NewEventController(eventService)

	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Validator = customValidator.CustomValidator{Validator: validator.New()}

	srv := &http.Server{Addr: ":" + appPort, Handler: e}

	authMiddleware := auth_middleware.NewAuthMiddleware(authService)
	tracingMiddleware := tracing_middleware.NewTracingMiddleware(true, true, true)

	storiesGroup := e.Group("/stories")
	storiesGroup.Use(tracingMiddleware.GetTracingInformationFromEcho)
	{
		storiesGroup.GET("/:token", storyController.GetStories, authMiddleware.Authenticate)
	}

	eventGroup := e.Group("/event")
	eventGroup.Use(tracingMiddleware.GetTracingInformationFromEcho)
	{
		eventGroup.POST("/:token", eventController.PostEvent, authMiddleware.Authenticate)
		eventGroup.GET("/metrics/:token", eventController.GetEventMetrics, authMiddleware.Authenticate)
	}

	e.GET("/swagger*", echoSwagger.WrapHandler)
	e.GET("/_monitoring/ready", apiHealthCheck)
	e.GET("/_monitoring/live", apiHealthCheck)

	log.Println("Server running at", appPort)

	err = graceful.ListenAndServe(srv, time.Duration(configs.AppConfig.GracefulShutdownTimeout)*time.Second)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
}

func apiHealthCheck(c echo.Context) error {
	return c.JSON(200, "Healthy")
}
