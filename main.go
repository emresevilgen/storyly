package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"github.com/tylerb/graceful"
	"log"
	"net/http"
	"storyly/configs"
	_ "storyly/docs"
	event_controller "storyly/pkg/controller/event-controller"
	"storyly/pkg/controller/story-controller"
	auth_middleware "storyly/pkg/middleware/auth-middleware"
	"storyly/pkg/middleware/tracing-middleware"
	postgresql_repository "storyly/pkg/repository/postgresql-repository"
	auth_service "storyly/pkg/service/auth-service"
	event_service "storyly/pkg/service/event-service"
	story_service "storyly/pkg/service/story-service"
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

	//if err != nil {
	//	fmt.Println("Db connection error,err:", err.Error())
	//	return
	//}

	// Init repositories
	storyRepository := postgresql_repository.NewStoryRepository()
	tokenRepository := postgresql_repository.NewTokenRepository()

	// Init services
	authService := auth_service.NewAuthService(tokenRepository)
	storyService := story_service.NewStoryService(storyRepository)
	eventService := event_service.NewEventService()

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
	}

	e.GET("/swagger*", echoSwagger.WrapHandler)
	e.GET("/_monitoring/ready", apiHealthCheck)
	e.GET("/_monitoring/live", apiHealthCheck)

	log.Println("Server running at", appPort)

	err := graceful.ListenAndServe(srv, time.Duration(configs.AppConfig.GracefulShutdownTimeout)*time.Second)
	if err != nil {
		log.Println("Error: ", err)

		return
	}
}

func apiHealthCheck(c echo.Context) error {
	return c.JSON(200, "Healthy")
}
