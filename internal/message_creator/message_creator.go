package message_creator

import (
	"fmt"
	"strings"
	"tg_weather_bot/pkg/client/weather_api"
	"tg_weather_bot/pkg/logging"
)

type MessageCreator struct {
	logger *logging.Logger
}

func NewMessageCreator() *MessageCreator {
	logger := logging.GetLogger()

	return &MessageCreator{
		logger: logger,
	}
}

func (MC *MessageCreator) CreateWeatherMessage(incomeMessage string) string {

	//var message string

	cityName := strings.TrimPrefix(incomeMessage, "/weather ")

	if cityName == "" {
		return "Уточните для какого города вы хотите получить погоду"
	}
	weather, err := weather_api.GetWeatherByCity(cityName)
	if err != nil {
		MC.logger.Errorf("Cant get weather for this city")
		return fmt.Sprintf("%s? Не знаю такого города. Может быть вы опечатлись?", cityName)
	}
	return fmt.Sprintf("В городе %s сейчас %d градусов °C. Ощущается как %d °C. Скорость ветра %f м/с, UV индекс: %d",
		cityName, int(weather.Temp), int(weather.TempApparent), weather.WindSpeed, weather.UVIndex)

}

func (MC *MessageCreator) CreateStatMessage(tgUserID int64) { /*todo return stat struct*/

	/*make db request to get statistic*/

}
