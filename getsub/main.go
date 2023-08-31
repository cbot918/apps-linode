package main

import (
	"log"

	i "getsub/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// cfg, err := i.NewConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	cfg := &i.Config{}
	f := fiber.New()
	api := i.NewAPI(cfg, f)

	log.Fatal(api.App.Listen(":8080"))
}
