package main

import (
	"fmt"
	"github.com/CRPNTRPINK/telegram_bot_go_practice/internal/service/product"
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

	productService := product.NewService()
	fmt.Println(productService)

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
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehaviour(bot, update.Message)
			}

		}
	}
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMsg := "Here all the products: \n\n"
	products := productService.List()
	for _, p := range products {
		outputMsg += p.Title
		outputMsg += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help"+
			"\n/list - list")
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
