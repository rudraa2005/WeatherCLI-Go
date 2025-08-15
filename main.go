package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

func getCity() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter city name: ")
	if !scanner.Scan() {
		return "", errors.New("failed to read input")
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("input error: %w", err)
	}

	city := scanner.Text()
	if city == "" {
		return "", errors.New("city name cannot be empty")
	}

	return city, nil
}

func fetchWeather(city, apiKey string) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	if weather.CityName == "" {
		return nil, errors.New("no weather data found — check city name")
	}

	return &weather, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not found in environment variables")
	}

	city, err := getCity()
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	weather, err := fetchWeather(city, apiKey)
	if err != nil {
		log.Fatalf("Weather fetch error: %v", err)
	}

	fmt.Printf("Weather in %s: %.2f°C, %s\n",
		weather.CityName,
		weather.Main.Temp,
		weather.Weather[0].Description)

	fmt.Println("Weather data fetched successfully.")
}
