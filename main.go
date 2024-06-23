package main

import (
	"fmt"
	"github.com/joho/godotenv"
	//"github.com/tidwall/pretty"
	"os"
	"weather_api_test/weather_api"
)

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Country   string `json:"country"`
	}
	Current struct {
		TempC     float32    `json:"temp_c"`
		TempF     float32    `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		}
	}
}

func main() {

	err :=loadEnv()
	if err != nil {
		fmt.Println("Env Error: ", err)
	}

	var city string

	fmt.Print("Enter city: ")
	fmt.Scan(&city)

	weather, err := weather_api.GetWeather(city, os.Getenv("WEATHER_API_KEY"))

	printData(weather)


	//fmt.Println(string(pretty.Pretty(body)))

}

func loadEnv() (err error) {
	err = godotenv.Load()
	return
}

func printData(weather weather_api.Weather) {
	fmt.Println("Location: ", weather.Location.Name + ",", weather.Location.Country)
	fmt.Println("Celsius: ", weather.Current.TempC,"c")
	fmt.Println("Fahrenheit: ", weather.Current.TempF,"f")
	fmt.Println("Condition: ", weather.Current.Condition.Text)
}
