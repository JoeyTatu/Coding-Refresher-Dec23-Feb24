package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

// ----- Consts and vars -----
const endpoint = "https://api.open-meteo.com/v1/forecast"

var pollInterval = time.Second * 5

// ----- Structs -----

type WeatherData struct {
	Elevation            float64           `json:"elevation"`
	Hourly               map[string][]any  `json:"hourly"`
	UTCOffsetSeconds     int               `json:"utc_offset_seconds"`
	Timezone             string            `json:"timezone"`
	TimezoneAbbreviation string            `json:"timezone_abbreviation"`
	HourlyUnits          map[string]string `json:"hourly_units"`
}

type WPoller struct {
	ClosedChannel chan struct{}
}

// ----- Structs as funcs -----
func (wp *WPoller) start() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file. Error:", err)
	}

	openCageApi := os.Getenv("OPEN_CAGE_API")

	lat, long, err := getCityCoordinates(openCageApi)
	if err != nil {
		log.Fatal("Error:", err)
	}

	ticker := time.NewTicker(pollInterval)

free:
	for {
		select {
		case <-ticker.C:
			data, err := getWeatherResults(lat, long)
			if err != nil {
				log.Fatal("Error:", err)
			}
			// fmt.Println(data)
			if err := wp.handleData(data); err != nil {
				log.Fatal("Error:", err)
			}
		case <-wp.ClosedChannel:
			break free
		}
	}
}

func (wp *WPoller) handleData(data *WeatherData) error {
	fmt.Printf("Time\t\tTemperature\tRelative Humidity\tDew Point\n")

	for i, time := range data.Hourly["time"] {
		temperature := data.Hourly["temperature_2m"][i]
		relativeHumidity := data.Hourly["relative_humidity_2m"][i]
		dewPoint := data.Hourly["dew_point_2m"][i]

		fmt.Printf("%v\t%.2f°C\t%.2f%%\t\t%.2f°C\n", time, temperature, relativeHumidity, dewPoint)
	}

	fmt.Printf("\nAdditional Information:\n")
	fmt.Printf("UTC Offset: %v seconds\n", data.UTCOffsetSeconds)
	fmt.Printf("Timezone: %v\n", data.Timezone)
	fmt.Printf("Timezone Abbreviation: %v\n", data.TimezoneAbbreviation)
	fmt.Printf("Elevation: %.2f meters\n", data.Elevation)
	fmt.Printf("Hourly Units: %v\n", data.HourlyUnits)

	return nil
}

func (wp *WPoller) close() {
	close(wp.ClosedChannel)
}

// ----- Funcs -----
func NewWPoller() *WPoller {
	return &WPoller{
		ClosedChannel: make(chan struct{}),
	}
}

func getCityCoordinates(apiKey string) (float64, float64, error) {
	fmt.Print("Enter a city (e.g., Dublin, Ireland):\n> ")
	var cityName string
	fmt.Scanln(&cityName)

	client := resty.New()

	resp, err := client.R().
		SetQueryParam("q", cityName).
		SetQueryParam("key", apiKey).
		Get("https://api.opencagedata.com/geocode/v1/json")

	if err != nil {
		return 0, 0, err
	}

	if resp.IsSuccess() {
		var result map[string]interface{}
		err := json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return 0, 0, err
		}

		coordinates := result["results"].([]interface{})[0].(map[string]interface{})
		geometry := coordinates["geometry"].(map[string]interface{})
		lat := geometry["lat"].(float64)
		lon := geometry["lng"].(float64)

		return lat, lon, nil
	}

	return 0, 0, fmt.Errorf("API request failed. Status code: %d", resp.StatusCode())
}

func getWeatherResults(lat, long float64) (*WeatherData, error) {
	// latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	// longStr := strconv.FormatFloat(long, 'f', -1, 64)

	fmt.Printf("Latitude: %.2f - Longitude: %.2f\n\n", lat, long)

	uri := fmt.Sprintf("%s?latitude=%.2f&longitude=%.2f&hourly=temperature_2m,relative_humidity_2m,dew_point_2m,apparent_temperature,precipitation,rain,showers,weather_code,cloud_cover,cloud_cover_low,cloud_cover_mid,cloud_cover_high,visibility,wind_speed_10m,wind_direction_10m,wind_gusts_10m,uv_index", endpoint, lat, long)

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

// ----- Main func -----
func main() {
	fmt.Println("Starting weather poller")
	wPoller := NewWPoller()
	wPoller.start()
}
