package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env doesn't found")
	}

	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			default:
				defaultBehaviour(bot, update.Message)
			}

		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help")
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func defaultBehaviour(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID

	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
