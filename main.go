package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type name struct {
	FieldName string `json:"fieldNameInJson"`
}

func getCity() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter city name: ")
	scanner.Scan()
	city := scanner.Text()
	if city == "" {
		fmt.Println("City name cannot be empty.")
		return
	}
	apiKey := "2754d1e3393485b5d5e8ec6e866cb816"
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey + "&units=metric"
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}
	var weatherResponse WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)

	fmt.Printf("Weather in %s: %.2f°C, %s\n", weatherResponse.CityName, weatherResponse.Main.Temp, weatherResponse.Weather[0].Description)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}
}

type WeatherResponse struct {
	CityName    string  `json:"name"`
	Temperature float64 `json:"main.temp"`
	Main        struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	getCity()
	fmt.Println("Weather data fetched successfully.")
}
