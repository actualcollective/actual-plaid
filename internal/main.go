package main

import (
	"actual-plaid/internal/config"
	"actual-plaid/internal/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		panic(fmt.Errorf("Loading app configuration failed: %d \n", err))
	}

	app := fiber.New()

	app.Static("/", "public")
	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./public/index.html")
	})

	api := app.Group("/api")
	apiRoutes(api)

	log.Fatal(app.Listen(":" + config.Config.AppPort))
}

func apiRoutes(app fiber.Router) {
	routes.InstallRouter(app)
	routes.PlaidRouter(app)
}
