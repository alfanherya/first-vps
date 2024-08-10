package usecase

import (
	"context"
	"first-app/model/response"
	"first-app/repository"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HealtUsecase struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	Validate         *validator.Validate
	HealthRepository *repository.HealtRepository
}

func NewHealthUsecase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, healtdRepository *repository.HealthRepository) *HealtUsecase {
	return &HealtUsecase{
		DB:               db,
		Log:              logger,
		Validate:         validate,
		HealthRepository: healthRepository,
	}
}

func (c *HealtUsecase) All(ctx context.Context) (*response.HealthResponse, error){
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	health := new(response.HealthResponse)
	var wg sync.WaitGroup
	var dbErr,

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		if err := c.HealthRepository.checkDB(tx, &health.Database); err !=nil {
			c.Log.Warnf("Failed check DB : %+v", err)
			dbErr = fiber
		}
	}()
}
