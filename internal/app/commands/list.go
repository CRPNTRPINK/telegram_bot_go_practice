package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(message *tgbotapi.Message) {
	outputMsg := "Here all the products: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outputMsg += p.Title
		outputMsg += "\n"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, outputMsg)

	c.bot.Send(msg)
}
