package router

import (
	"github.com/duchoang206h/telebot-storage/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/api")
	fileRoute := api.Group("/file")
	fileRoute.Get("/:FileID", handler.GetFileHandler)
	fileRoute.Post("/", handler.UploadFileHandler)
}
