package handlers

import (
	"vya4ikBot/internal/buttons"

	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

func MessageHandler(message *tgBotApi.Message, bot *tgBotApi.BotAPI) {
	log.Debugf("[%s] %s", message.From.UserName, message.Text)
	if message.Command() == "" {
		switch message.Text {
		case "Главное меню":
			msg := tgBotApi.NewMessage(message.Chat.ID, message.Text)
			msg.ReplyMarkup = buttons.FeaturesKeyboard()
			bot.Send(msg)
		default:
			bot.Send(tgBotApi.NewMessage(message.Chat.ID, "Простите, я вас не понимаю"))
		}
	} else {
		switch message.Command() {
		case "start":
			msg := tgBotApi.NewMessage(message.Chat.ID, "Привет, я "+bot.Self.FirstName+"! На данный момент я предоставляю оперативную информацию по Covid-19 в России\n"+"[Источник](https://стопкоронавирус.рф)")
			msg.ParseMode = "markdown"
			msg.ReplyMarkup = buttons.MainMenuButton()
			bot.Send(msg)
		default:
			bot.Send(tgBotApi.NewMessage(message.Chat.ID, "К сожалению, я не знаю такую команду"))
		}
	}
}
