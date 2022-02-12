package handlers

import (
	"fmt"
	"vya4ikBot/internal/data"

	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

func CallbackHandler(callbackQuery *tgBotApi.CallbackQuery, bot *tgBotApi.BotAPI, data *data.CovidData) {
	log.Debugf("[%s] %s", callbackQuery.From.UserName, callbackQuery.Data)

	msg := tgBotApi.NewMessage(callbackQuery.Message.Chat.ID, "")
	msg.ParseMode = "markdown"

	switch callbackQuery.Data {
	case "CovidData":
		msg.Text = fmt.Sprintf("%s:\n\n"+
			"Выявлено случаев: *%s*\n"+
			"Человек выздоровело: *%s*\n"+
			"Человек умерло: *%s*\n\n",
			data.Date, data.Sick, data.Healed, data.Died)
		bot.Send(msg)
	case "СovidChange":
		msg.Text = fmt.Sprintf("%s:\n\n"+
			"Выявлено случаев: *%s*\n"+
			"Человек выздоровело: *%s*\n"+
			"Человек умерло: *%s*\n\n",
			data.Date, data.SickChange, data.HealedChange, data.DiedChange)
		bot.Send(msg)
	}

}
