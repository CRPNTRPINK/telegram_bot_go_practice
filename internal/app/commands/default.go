package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Default(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote: "+message.Text)
	msg.ReplyToMessageID = message.MessageID

	c.bot.Send(msg)
}
