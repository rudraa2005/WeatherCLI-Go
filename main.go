package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	CityName string `json:"name"`
	Main     struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
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

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY not found in environment variables.")
		return
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}
	defer resp.Body.Close()

	var weatherResponse WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	fmt.Printf("Weather in %s: %.2f°C, %s\n",
		weatherResponse.CityName,
		weatherResponse.Main.Temp,
		weatherResponse.Weather[0].Description)
}

func main() {
	getCity()
	fmt.Println("Weather data fetched successfully.")
}
