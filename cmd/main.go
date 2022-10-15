package main

import (
	nestedFormatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/go-co-op/gocron"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
	"vya4ikBot/internal/config"
	"vya4ikBot/internal/handlers"
	"vya4ikBot/internal/model"
)

func main() {
	log.SetFormatter(&nestedFormatter.Formatter{
		NoColors:        true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	f, err := os.OpenFile("bot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	config := config.GetConfig()

	level, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)

	data := model.NewCovidData()
	data.GetData()
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}
	scheduler := gocron.NewScheduler(location)
	if _, err := scheduler.Cron(config.Cron).Do(data.GetData); err != nil {
		log.Fatal(err)
	}
	scheduler.StartAsync()

	bot, err := tgBotApi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = config.BotDebug

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
