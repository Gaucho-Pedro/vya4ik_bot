package main

import (
	"time"
	"vya4ikBot/internal/config"
	"vya4ikBot/internal/data"
	"vya4ikBot/internal/handlers"

	"github.com/go-co-op/gocron"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	config := config.GetConfig()

	level, _ := log.ParseLevel(config.LogLevel)
	log.SetLevel(level)

	data := data.NewCovidData()
	data.GetData()

	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Cron(config.Cron).Do(data.GetData)
	scheduler.StartAsync()

	bot, err := tgBotApi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = log.GetLevel().String() == "debug"

	log.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgBotApi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			go handlers.CallbackHandler(update.CallbackQuery, bot, data)
		} else if update.Message != nil {
			go handlers.MessageHandler(update.Message, bot)
		}
	}
}
