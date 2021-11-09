package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Model struct {
	date string

	sick       string
	sickChange string

	healed       string
	healedChange string

	died       string
	diedChange string
}

func parsing() Model {
	var model Model
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://стопкоронавирус.рф/information/", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			model.date = strings.ToLower(r.HTMLDoc.Find("div.cv-section__title-wrapper").Find("small").Last().Text())

			r.HTMLDoc.Find("div.cv-stats-virus__item").Each(func(i int, s *goquery.Selection) {
				switch i {
				case 0:
					model.sick = strings.Trim(s.Find("H3").Text(), " \n")
				case 1:
					model.sickChange = strings.Trim(s.Find("H3").Text(), " \n")
				case 2:
					model.healed = strings.Trim(s.Find("H3").Text(), " \n")
				case 3:
					model.healedChange = strings.Trim(s.Find("H3").Text(), " \n")
				case 4:
					model.died = strings.Trim(s.Find("H3").Text(), " \n")
				case 5:
					model.diedChange = strings.Trim(s.Find("H3").Text(), " \n")
				}
			})
		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()
	return model
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
			log.Printf("[%s] %s", update.CallbackQuery.From.UserName, update.CallbackQuery.Data)
			bot.Send(tgBotApi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Делаю запрос..."))
			model := parsing()
			text := fmt.Sprintf("*В России %s:*\n\n"+
				"*Выявлено случаев: *%s\n"+
				"*Человек выздоровело: *%s\n"+
				"*Человек умерло: *%s\n\n"+
				"*Выявлено случаев за сутки: *%s\n"+
				"*Человек выздоровело за сутки: *%s\n"+
				"*Человек умерло за сутки: *%s\n\n"+
				"[Источник](https://стопкоронавирус.рф)",
				model.date, model.sick, model.healed, model.died, model.sickChange, model.healedChange, model.diedChange)
			msg := tgBotApi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
			msg.ParseMode = "markdown"
			bot.Send(msg)
		}

		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			command := update.Message.Command()
			message := update.Message.Text

			if command == "" {
				switch message {
				case "Главное меню":
					message2 := tgBotApi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					message2.ReplyMarkup = tgBotApi.NewInlineKeyboardMarkup(tgBotApi.NewInlineKeyboardRow(tgBotApi.NewInlineKeyboardButtonData("Корона", "/covid")))
					bot.Send(message2)
				}
			} else {
				switch command {
				case "start":
					/*TODO: Вынести клаву в отдельный класс*/
					message1 := tgBotApi.NewMessage(update.Message.Chat.ID, "Привет, я Vya4ikBot!")
					message1.ReplyMarkup = tgBotApi.NewReplyKeyboard(tgBotApi.NewKeyboardButtonRow(tgBotApi.NewKeyboardButton("Главное меню")))
					bot.Send(message1)
				case "covid":
					bot.Send(tgBotApi.NewMessage(update.Message.Chat.ID, "Делаю запрос..."))
					model := parsing()
					text := fmt.Sprintf("*%s:*\n\n"+
						"*Выявлено случаев: *%s\n"+
						"*Человек выздоровело: *%s\n"+
						"*Человек умерло: *%s\n\n"+
						"*Выявлено случаев за сутки: *%s\n"+
						"*Человек выздоровело за сутки: *%s\n"+
						"*Человек умерло за сутки: *%s\n\n"+
						"[Источник](https://стопкоронавирус.рф)",
						model.date, model.sick, model.healed, model.died, model.sickChange, model.healedChange, model.diedChange)
					msg := tgBotApi.NewMessage(update.Message.Chat.ID, text)
					msg.ParseMode = "markdown"
					bot.Send(msg)
				default:
					bot.Send(tgBotApi.NewMessage(update.Message.Chat.ID, "К сожалению, я не знаю такую команду"))
				}
			}
		}
	}
}

func main() {
	bot()
}
