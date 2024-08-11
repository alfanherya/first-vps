package config

import (
	"first-app/controller"
	"first-app/repository"
	"first-app/route"
	"first-app/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	healthRepository := repository.NewHealthRepository(config.Log)

	// setup use cases
	healthUseCase := usecase.NewHealthUsecase(config.DB, config.Log, config.Validate, healthRepository)

	// setup controller
	healthController := controller.NewHealthController(healthUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:              config.App,
		HealthController: healthController,
	}

	routeConfig.Setup()
}
