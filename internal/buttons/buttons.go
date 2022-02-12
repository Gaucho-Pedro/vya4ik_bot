package buttons

import tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"

func MainMenu() *tgBotApi.ReplyKeyboardMarkup {
	keyboard := tgBotApi.NewReplyKeyboard(tgBotApi.NewKeyboardButtonRow(tgBotApi.NewKeyboardButton("Главное меню")))
	return &keyboard
}

func InlineKeyboard() *tgBotApi.InlineKeyboardMarkup {
	keyboard := tgBotApi.NewInlineKeyboardMarkup(
		tgBotApi.NewInlineKeyboardRow(
			tgBotApi.NewInlineKeyboardButtonData("Статистика за весь период", "CovidData"),
			tgBotApi.NewInlineKeyboardButtonData("Статистика за сутки", "СovidChange")))
	return &keyboard
}
