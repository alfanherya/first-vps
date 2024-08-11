package controller

import (
	"first-app/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type HealthController struct {
	Log     *logrus.Logger
	Usecase *usecase.HealtUsecase
}

func NewHealthController(usecase *usecase.HealtUsecase, logger *logrus.Logger) *HealthController {
	return &HealthController{
		Log:     logger,
		Usecase: usecase,
	}
}

func (c *HealthController) CheckDB(ctx *fiber.Ctx) error {
	response, err := c.Usecase.CheckDB(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to check db : %+v", err)
		return err
	}
	return ctx.JSON(response)
}
