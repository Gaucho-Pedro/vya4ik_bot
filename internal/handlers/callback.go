package handlers

import (
	"fmt"
	"vya4ikBot/internal/buttons"
	"vya4ikBot/internal/model"

	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

var format = "%s:\n\n" +
	"Выявлено случаев: *%s*\n" +
	"Выздоровело: *%s*\n" +
	"Умерло: *%s*\n\n"

func CallbackHandler(callbackQuery *tgBotApi.CallbackQuery, bot *tgBotApi.BotAPI, data *model.CovidData) {
	log.Debugf("[%s] %s", callbackQuery.From.UserName, callbackQuery.Data)

	msg := tgBotApi.NewMessage(callbackQuery.Message.Chat.ID, "")
	msg.ReplyMarkup = buttons.CovidKeyboard()
	msg.ParseMode = "markdown"

	switch callbackQuery.Data {
	case "CovidData":
		msg.Text = fmt.Sprintf(format, data.Date, data.Sick, data.Healed, data.Died)
		bot.Send(msg)
	case "СovidChange":
		msg.Text = fmt.Sprintf(format, data.Date, data.SickChange, data.HealedChange, data.DiedChange)
		bot.Send(msg)
	case "Covid":
		msg.Text = "Выберите период:"
		bot.Send(msg)
	case "Back":
		//bot.Send(tgBotApi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))
		msg.Text = "Главное меню"
		msg.ReplyMarkup = buttons.FeaturesKeyboard()
		bot.Send(msg)
	}
}
