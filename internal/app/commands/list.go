package commands

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(message *tgbotapi.Message) {
	outputMsg := "Here all the products: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outputMsg += p.Title
		outputMsg += "\n"
	}

	serializedData, _ := json.Marshal(CommandData{Offset: 10})
	msg := tgbotapi.NewMessage(message.Chat.ID, outputMsg)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)

	c.bot.Send(msg)
}
