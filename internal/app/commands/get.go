package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	arg, err := strconv.Atoi(args)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "wrong arguments. The right example: /get 1")
		c.bot.Send(msg)
		return
	}

	product, err := c.productService.Get(arg)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)

}
