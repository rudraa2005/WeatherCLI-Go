# Weather CLI App

A simple command-line weather application written in Go that fetches current weather data for any city using the OpenWeatherMap API.

## Features

- Get current weather information for any city
- Displays temperature in Celsius
- Shows weather description (e.g., clear sky, light rain)
- Clean command-line interface
- Environment variable configuration for API key security

## Prerequisites

- Go 1.16 or higher
- OpenWeatherMap API key (free at [openweathermap.org](https://openweathermap.org/api))

## Installation

1. Clone or download this repository
2. Navigate to the project directory
3. Install dependencies:
   ```bash
   go mod init weather-app
   go get github.com/joho/godotenv
   ```

## Setup

1. Sign up for a free API key at [OpenWeatherMap](https://openweathermap.org/api)
2. Create a `.env` file in the project root directory
3. Add your API key to the `.env` file:
   ```
   API_KEY=your_openweathermap_api_key_here
   ```

## Usage

1. Run the application:
   ```bash
   go run main.go
   ```

2. When prompted, enter the name of the city you want to check:
   ```
   Enter city name: London
   ```

3. The app will display the current weather:
   ```
   Weather in London: 15.32°C, few clouds
   Weather data fetched successfully.
   ```

## Building

To create an executable binary:

```bash
go build -o weather-app main.go
```

Then run the executable:
```bash
./weather-app
```

## Code Structure

- `WeatherResponse`: Struct that maps the JSON response from OpenWeatherMap API
- `getCity()`: Main function that handles user input and API communication
- `main()`: Entry point that calls getCity() and displays completion message

## API Response

The application uses OpenWeatherMap's Current Weather Data API endpoint:
```
https://api.openweathermap.org/data/2.5/weather?q={city}&appid={API_KEY}&units=metric
```

## Error Handling

The application handles several error conditions:
- Empty city name input
- Missing or invalid `.env` file
- Missing API key in environment variables
- Network errors when fetching data
- Invalid JSON response parsing

## Dependencies

- `github.com/joho/godotenv` - For loading environment variables from `.env` file

## Example Output

```
Enter city name: Tokyo
Weather in Tokyo: 22.15°C, clear sky
Weather data fetched successfully.
```

## Security Note

Never commit your `.env` file or expose your API key in your code. The `.env` file should be added to your `.gitignore` file:

```gitignore
.env
```

## License

This project is open source and available under the MIT License.