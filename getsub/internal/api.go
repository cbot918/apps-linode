package internal

import (
	"fmt"
	"log"

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

	// api.App.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("home")
	// })

	api.App.Get("/", func(c *fiber.Ctx) error {
		err := api.C.GetSub(c)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	})

	api.App.Get("/ping", func(c *fiber.Ctx) error {
		fmt.Println("in /ping")
		return c.SendString("pong")
	})

}
