package route

import (
	"first-app/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type RouteConfig struct {
	App              *fiber.App
	HealthController *controller.HealthController
}

func (c *RouteConfig) Setup() {
	c.App.Get("/metrics", monitor.New())
	c.SetupHealthRoute()

}

func (c *RouteConfig) SetupHealthRoute() {
	group := c.App.Group("/health")
	group.Get("/db", c.HealthController.CheckDB)
}
