# WeatherCLI-Go

A simple command-line weather application written in Go that fetches current weather data for any city using the OpenWeatherMap API.

## Features

- Interactive city input
- Current temperature display in Celsius
- Weather description (e.g., clear sky, light rain)
- Error handling for invalid cities and API issues
- Environment variable configuration for API key

## Prerequisites

- Go 1.16 or higher
- OpenWeatherMap API key (free registration required)

## Installation

1. Clone or download the project files
2. Navigate to the project directory
3. Install dependencies:
   ```bash
   go mod init WeatherCLI-Go
   go get github.com/joho/godotenv
   ```

## Setup

1. Get your free API key from [OpenWeatherMap](https://openweathermap.org/api):
   - Sign up for a free account
   - Navigate to the API keys section
   - Copy your API key

2. Create a `.env` file in the project root directory:
   ```bash
   cp .env.sample .env
   ```

3. Edit the `.env` file and replace the placeholder with your actual API key:
   ```
   API_KEY=your_actual_api_key_here
   ```

## Usage

1. Run the application:
   ```bash
   go run main.go
   ```

2. Enter a city name when prompted:
   ```
   Enter city name: London
   ```

3. View the weather information:
   ```
   Weather in London: 15.32°C, broken clouds
   Weather data fetched successfully.
   ```

## Example Output

```
Enter city name: New York
Weather in New York: 22.15°C, clear sky
Weather data fetched successfully.
```

## Error Handling

The application handles various error scenarios:

- **Empty city name**: Prompts for valid input
- **Invalid city name**: Returns "no weather data found" message
- **Network issues**: Reports connection errors
- **Invalid API key**: Returns API authentication error
- **Missing .env file**: Application exits with error message

## Project Structure

```
WeatherCLI-Go/
├── main.go           # Main application code
├── .env              # Environment variables (create from .env.sample)
├── .env.sample       # Environment variables template
├── go.mod            # Go module file
├── go.sum            # Go dependencies checksum
└── README.md         # This file
```

## Dependencies

- `github.com/joho/godotenv` - For loading environment variables from .env file

## API Information

This application uses the [OpenWeatherMap Current Weather Data API](https://openweathermap.org/current):
- **Endpoint**: `https://api.openweathermap.org/data/2.5/weather`
- **Units**: Metric (Celsius)
- **Rate Limit**: 1,000 calls per day (free tier)

## Building for Production

To build a standalone executable:

```bash
go build -o WeatherCLI-Go main.go
```

Then run:
```bash
./WeatherCLI-Go
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Support

If you encounter any issues:
1. Verify your API key is correct and active
2. Check your internet connection
3. Ensure the city name is spelled correctly
4. Check the OpenWeatherMap API status

For additional help, please open an issue in the repository.