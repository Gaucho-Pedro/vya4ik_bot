package main

import (
	"fmt"
	"log"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func parsing() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://стопкоронавирус.рф/information/", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			fmt.Println(r.HTMLDoc.Find("small").Text())
			// r.HTMLDoc.Find("div.cv-stats-virus__item" /*"h3.cv-stats-virus__item-value"*/).Each(func(i int, s *goquery.Selection) {
			// 	text := strings.Trim(s.Find("H3").Text(), " \n")
			// 	fmt.Println(i, text)
			// })
		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()
}

func bot() {
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

func main() {
	//bot()
	parsing()
}
