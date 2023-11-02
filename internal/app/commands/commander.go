package commands

import (
	"github.com/CRPNTRPINK/telegram_bot_go_practice/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{bot: bot, productService: productService}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil { // If we got a message
		return
	}
	command, ok := registeredCommands[update.Message.Command()]
	if ok {
		command(c, update.Message)
	} else {
		c.Default(update.Message)
	}
}