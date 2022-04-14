package buttons

import tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"

func MainMenuButton() *tgBotApi.ReplyKeyboardMarkup {
	keyboard := tgBotApi.NewReplyKeyboard(tgBotApi.NewKeyboardButtonRow(tgBotApi.NewKeyboardButton("Главное меню")))
	return &keyboard
}
func FeaturesKeyboard() *tgBotApi.InlineKeyboardMarkup {
	keyboard := tgBotApi.NewInlineKeyboardMarkup(
		tgBotApi.NewInlineKeyboardRow(tgBotApi.NewInlineKeyboardButtonData("Статистика по Covid-19 в России", "Covid")))
	return &keyboard
}
func CovidKeyboard() *tgBotApi.InlineKeyboardMarkup {
	keyboard := tgBotApi.NewInlineKeyboardMarkup(
		tgBotApi.NewInlineKeyboardRow(
			tgBotApi.NewInlineKeyboardButtonData("За весь период", "CovidData"),
			tgBotApi.NewInlineKeyboardButtonData("За сутки", "СovidChange")),
		tgBotApi.NewInlineKeyboardRow(tgBotApi.NewInlineKeyboardButtonData("Назад", "Back")))
	return &keyboard
}
