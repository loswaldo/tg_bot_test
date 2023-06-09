package weather_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"tg_weather_bot/internal/config"
	"tg_weather_bot/pkg/logging"
	"time"
)

type Weather struct {
	Temp          float64
	TempApparent  float64
	Humidity      float64
	UVIndex       int64
	WindSpeed     float64
	CloudCover    float64
	RainIntensity int64
	SnowIntensity int64
}

type weatherBody struct {
	Data struct {
		Time   time.Time `json:"time"`
		Values struct {
			CloudBase                float64 `json:"cloudBase"`
			CloudCeiling             float64 `json:"cloudCeiling"`
			CloudCover               float64 `json:"cloudCover"`
			DewPoint                 float64 `json:"dewPoint"`
			FreezingRainIntensity    int64   `json:"freezingRainIntensity"`
			Humidity                 float64 `json:"humidity"`
			PrecipitationProbability int64   `json:"precipitationProbability"`
			PressureSurfaceLevel     float64 `json:"pressureSurfaceLevel"`
			RainIntensity            int64   `json:"rainIntensity"`
			SleetIntensity           int64   `json:"sleetIntensity"`
			SnowIntensity            int64   `json:"snowIntensity"`
			Temperature              float64 `json:"temperature"`
			TemperatureApparent      float64 `json:"temperatureApparent"`
			UvHealthConcern          int64   `json:"uvHealthConcern"`
			UvIndex                  int64   `json:"uvIndex"`
			Visibility               int64   `json:"visibility"`
			WeatherCode              int64   `json:"weatherCode"`
			WindDirection            float64 `json:"windDirection"`
			WindGust                 float64 `json:"windGust"`
			WindSpeed                float64 `json:"windSpeed"`
		} `json:"values"`
	} `json:"data"`
	Location struct {
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
		Name    string  `json:"name"`
		ObjType string  `json:"type"`
	} `json:"location"`
}

func GetWeatherByCity(cityName string) (*Weather, error) {
	urlRequest := "https://api.tomorrow.io/v4/weather/realtime?location=%s&apikey=%s"

	cfg := config.GetAPIConfig()
	logger := logging.GetLogger()

	req, _ := http.NewRequest("GET", fmt.Sprintf(urlRequest, url.QueryEscape(cityName), cfg.WeatherAPIKey), http.NoBody)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("wrong city name")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var wBody weatherBody

	err := json.Unmarshal(body, &wBody)
	if err != nil {
		logger.Errorf("can't unmarshal json err: %v", err)
	}

	return &Weather{
		Temp:          wBody.Data.Values.Temperature,
		TempApparent:  wBody.Data.Values.TemperatureApparent,
		Humidity:      wBody.Data.Values.Humidity,
		UVIndex:       wBody.Data.Values.UvIndex,
		WindSpeed:     wBody.Data.Values.WindSpeed,
		CloudCover:    wBody.Data.Values.CloudCover,
		RainIntensity: wBody.Data.Values.RainIntensity,
		SnowIntensity: wBody.Data.Values.SnowIntensity,
	}, nil

}
