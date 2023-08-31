package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type API struct {
	App *fiber.App
	C   *Controller
}

func NewAPI(cfg *Config, app *fiber.App) *API {

	api := new(API)
	api.App = app

	api.App.Use(logger.New())
	api.App.Use(recover.New())
	api.setupRouter()

	api.C = NewController()

	return api
}

func (api *API) setupRouter() {

	api.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.App.Post("/signup", func(c *fiber.Ctx) error {
		return api.C.Singup(c)
	})

	api.App.Post("/signin", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.App.Get("/insert", func(c *fiber.Ctx) error {
		return api.C.Insert(c)
	})

}
