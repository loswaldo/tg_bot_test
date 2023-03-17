package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg_weather_bot/internal/message_creator"

	"tg_weather_bot/internal/config"
	"tg_weather_bot/pkg/logging"
)

func main() {

	logger := logging.GetLogger()
	logger.Info("create logger")

	cfg := config.GetAPIConfig()

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPIKey)
	if err != nil {
		logger.Fatalf("can't connect to tg bot err: %v", err)
	}

	dbCfg := config.GetDBConfig()
	MC, err := message_creator.NewMessageCreator(dbCfg)
	if err != nil {
		logger.Fatalf("can't create new message creator err: %v", err)
	}

	bot.Debug = false

	logger.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "stat":
			logger.Infof("Get stat message from %d", update.Message.From.ID)
			msg.Text = MC.CreateStatMessage(update.Message.From.ID)
		case "weather":
			logger.Infof("Get weather message from %d", update.Message.From.ID)
			msg.Text = MC.CreateWeatherMessage(update.Message)
		case "start":
			logger.Infof("Get start message from %d", update.Message.From.ID)
			msg.Text = "Привет! Я бот погоды. Напиши мне /weather и  Название_города(в именительном падеже) и я покажу тебе погоду"
		default:
			msg.Text = "I don't know that command"
		}

		logger.Infof("sending message [%s]", msg.Text)

		if _, err := bot.Send(msg); err != nil {
			logger.Errorf("Can't send message err: %v", err)
		}
	}
}

//weather_api.GetWeatherByCity("Москва")
//}
