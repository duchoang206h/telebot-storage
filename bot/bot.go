package bot

import (
	"log"

	"github.com/duchoang206h/telebot-storage/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TeleBot struct {
	bot *tgbotapi.BotAPI
}

var teleBot *TeleBot

func InitBot() error {
	_bot, err := tgbotapi.NewBotAPI(config.Config("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	_bot.Debug = true
	if err != nil {
		return err
	}
	teleBot = &TeleBot{bot: _bot}
	return nil
}

func GetBot() *TeleBot {
	return teleBot
}

func (teleBot *TeleBot) GetFileUrl(FileID string) (string, error) {
	file, err := teleBot.bot.GetFile(tgbotapi.FileConfig{FileID: FileID})
	if err != nil {
		return "", err
	}
	fileUrl := file.Link(teleBot.bot.Token)
	return fileUrl, nil
}

func (teleBot *TeleBot) UploadFile(fileBytes []byte, fileName string, chatID int64) (string, error) {
	fileConfig := tgbotapi.FileBytes{
		Name:  fileName,
		Bytes: fileBytes,
	}
	document := tgbotapi.NewDocumentUpload(chatID, fileConfig)
	message, err := teleBot.bot.Send(document)
	if err != nil {
		return "", err
	}
	return message.Document.FileID, nil
}
