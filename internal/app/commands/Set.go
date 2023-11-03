package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Set(inputMessage *tgbotapi.Message) {
	if len(inputMessage.CommandArguments()) == 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "wrong argument")
		c.bot.Send(msg)
		return
	}
	c.productService.Set(inputMessage.CommandArguments())
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("value: %s added", inputMessage.CommandArguments()))
	c.bot.Send(msg)
}
