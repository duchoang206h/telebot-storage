package main

import (
	"log"

	"github.com/duchoang206h/telebot-storage/bot"
	"github.com/duchoang206h/telebot-storage/config"
	"github.com/duchoang206h/telebot-storage/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	bot.InitBot()
	router.SetupRoute(app)
	log.Fatal(app.Listen(config.Config("APP_PORT")))
}
