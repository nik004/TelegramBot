package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	TOKEN := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		switch update.Message.Command() {

		case "Help":
			{
				CoomandHelp(&update, bot)
			}
		default:
			{
				DefaultMessage(&update, bot)
			}
		}

	}
}

func CoomandHelp(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You request HELP")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func DefaultMessage(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote "+update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
