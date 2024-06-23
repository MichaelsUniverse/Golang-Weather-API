package weather_api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	}
	Current struct {
		TempC     float32 `json:"temp_c"`
		TempF     float32 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		}
	}
}

func GetWeather(city string, api_key string) (weather Weather, err error) {

	url := "http://api.weatherapi.com/v1/current.json?key=" + os.Getenv("WEATHER_API_KEY") + "&q=" + strings.ToLower(city) + "&aqi=no"

	req, errNewReq := http.NewRequest("GET", url, nil)
	res, errDoReq := http.DefaultClient.Do(req)

	if err := errors.Join(errNewReq, errDoReq); err != nil {
		return weather, err
	}

	defer res.Body.Close()
	body, errReadAll := io.ReadAll(res.Body)

	errUnmarshal := json.Unmarshal(body, &weather)

	if err := errors.Join(errReadAll, errUnmarshal); err != nil {
		return weather, err
	}
	return weather, err
}
