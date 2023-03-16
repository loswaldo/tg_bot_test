package message_creator

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"tg_weather_bot/internal/config"
	"tg_weather_bot/pkg/client/postreSQL"
	"tg_weather_bot/pkg/client/weather_api"
	"tg_weather_bot/pkg/logging"
	"time"
)

type MessageCreator struct {
	logger *logging.Logger
	db     *postreSQL.PostgresDB
}

func NewMessageCreator(cfg *config.DBConfig) *MessageCreator {
	logger := logging.GetLogger()
	db, err := postreSQL.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Can't connect to db err: %v", err)
	}
	return &MessageCreator{
		logger: logger,
		db:     db,
	}
}

func (MC *MessageCreator) CreateWeatherMessage(message *tgbotapi.Message) string {
	cityName := strings.TrimPrefix(message.Text, "/weather ")

	if cityName == "" {
		return "Уточните для какого города вы хотите получить погоду"
	}
	weather, err := weather_api.GetWeatherByCity(cityName)
	if err != nil {
		MC.logger.Errorf("Can't get weather for this city")
		return fmt.Sprintf("%s? Не знаю такого города. Может быть вы опечатлись?", cityName)
	}
	err = MC.db.AddStatic(message.From.ID, cityName, time.Now())
	if err != nil {
		MC.logger.Errorf("Can't add statistic user_id :[%d], err : %v", message.From.ID, err)
	}
	return fmt.Sprintf("В городе %s сейчас %d градусов °C. Ощущается как %d °C. Скорость ветра %f м/с, UV индекс: %d",
		cityName, int(weather.Temp), int(weather.TempApparent), weather.WindSpeed, weather.UVIndex)

}

func (MC *MessageCreator) CreateStatMessage(tgUserID int64) string {

	stat, err := MC.db.GetStatisticByUserID(tgUserID)
	if err != nil {
		MC.logger.Errorf("can't get statistic for user %d", tgUserID)
		return "Что-то пошло не так и я не смог найти статиску для вас. Извините :( Попробуйте, пожалуйста, позже"
	}

	count := 0

	var firstTime time.Time

	if len(stat) != 0 {
		firstTime = stat[0].TimeStamp
	}
	count = len(stat)

	return fmt.Sprintf("Ого! вы сделали целых %d запросов. А ваш первый был %d %s %d", count, firstTime.Day(), firstTime.Month(), firstTime.Year())

}
