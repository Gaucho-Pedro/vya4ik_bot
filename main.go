package main

import (
	"log"

	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgBotApi.NewBotAPI("2096644322:AAH12TCiE78BXysiCpwvJHJ6MeBfyvHwxeo")
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgBotApi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {

		if update.CallbackQuery != nil {
			msg := tgBotApi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			bot.Send(msg)
		}

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		command := update.Message.Command()
		message := update.Message.Text

		if command == "" {
			switch message {
			case "Главное меню":
				message2 := tgBotApi.NewMessage(update.Message.Chat.ID, "Вот что я умею:")
				message2.ReplyMarkup = tgBotApi.NewInlineKeyboardMarkup(tgBotApi.NewInlineKeyboardRow(tgBotApi.NewInlineKeyboardButtonData("Корона", "Корона"), tgBotApi.NewInlineKeyboardButtonData("Старт", "/start")))
				bot.Send(message2)
			}
		} else {
			switch command {
			case "start":
				//TODO: Вынести клаву в отдельный класс
				message1 := tgBotApi.NewMessage(update.Message.Chat.ID, "Привет, я Vya4ikBot!")
				message1.ReplyMarkup = tgBotApi.NewReplyKeyboard(tgBotApi.NewKeyboardButtonRow(tgBotApi.NewKeyboardButton("Главное меню")))
				bot.Send(message1)

				message2 := tgBotApi.NewMessage(update.Message.Chat.ID, "Вот что я умею:")
				message2.ReplyMarkup = tgBotApi.NewInlineKeyboardMarkup(tgBotApi.NewInlineKeyboardRow(tgBotApi.NewInlineKeyboardButtonData("Корона", "Корона"), tgBotApi.NewInlineKeyboardButtonData("Старт", "/start")))
				bot.Send(message2)
			default:
				bot.Send(tgBotApi.NewMessage(update.Message.Chat.ID, "К сожалению я не знаю такую команду"))
			}
		}
	}
}
