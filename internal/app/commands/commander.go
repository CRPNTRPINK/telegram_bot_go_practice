package commands

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/CRPNTRPINK/telegram_bot_go_practice/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{bot: bot, productService: productService}
}

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		c.bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprint(parsedData)))
		return
	}

	if update.Message != nil { // If we got a message
		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		case "set":
			c.Set(update.Message)
		default:
			c.Default(update.Message)
		}

		return
	}
}
