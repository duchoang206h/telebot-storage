package handler

import (
	"io"
	"strconv"

	"github.com/duchoang206h/telebot-storage/bot"
	"github.com/duchoang206h/telebot-storage/config"
	"github.com/gofiber/fiber/v2"
)

func GetFileHandler(c *fiber.Ctx) error {
	fileID := c.Params("FileID")
	telebot := bot.GetBot()
	fileUrl, err := telebot.GetFileUrl(fileID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result": nil,
			"error":  err,
		})
	}
	return c.JSON(fiber.Map{
		"result": fileUrl,
	})
}

func UploadFileHandler(c *fiber.Ctx) error {
	telebot := bot.GetBot()
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid file",
		})
	}
	f, _ := file.Open()
	// get fileBytes from f or file
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		if err != io.EOF {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error reading file: " + err.Error(),
			})
		}
	}
	chat_id, _ := strconv.Atoi(config.Config("CHAT_ID"))
	fileId, err := telebot.UploadFile(fileBytes, file.Filename, int64(chat_id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(fiber.Map{
		"result": fileId,
	})
}
