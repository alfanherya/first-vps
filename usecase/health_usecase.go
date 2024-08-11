package usecase

import (
	"context"
	"first-app/model/response"
	"first-app/repository"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HealtUsecase struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	Validate         *validator.Validate
	HealthRepository *repository.HealthRepository
}

func NewHealthUsecase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, healtdRepository *repository.HealthRepository) *HealtUsecase {
	return &HealtUsecase{
		DB:               db,
		Log:              logger,
		Validate:         validate,
		HealthRepository: healtdRepository,
	}
}

func (c *HealtUsecase) All(ctx context.Context) (*response.HealthResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	health := new(response.HealthResponse)
	var wg sync.WaitGroup
	var dbErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := c.HealthRepository.CheckDB(tx, &health.Database); err != nil {
			c.Log.Warnf("Failed check DB : %+v", err)
			dbErr = fiber.ErrInternalServerError
		}
	}()

	wg.Wait() // Wait for the goroutine to finish

	if dbErr != nil {
		return nil, dbErr
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed Commit Transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return health, nil
}

func (c *HealtUsecase) CheckDB(ctx context.Context) (*response.CheckDBResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	health := new(response.CheckDBResponse)

	if err := c.HealthRepository.CheckDB(tx, health); err != nil {
		c.Log.Warnf("Failed check DB : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return health, nil
}
