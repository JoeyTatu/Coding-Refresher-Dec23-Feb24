package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	APIKey string `json:"OpenWeatherAppApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		TemperatureKelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(fileName string) (apiConfigData, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return apiConfigData{}, err
	}

	var config apiConfigData
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return apiConfigData{}, err
	}

	return config, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go!\n"))
}

func query(city string) (weatherData, error) {
	apiConfigData, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric&appid=" + apiConfigData.APIKey)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var data weatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return weatherData{}, err
	}

	return data, nil

}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		responseData := struct {
			Name string `json:"name"`
			Main struct {
				TemperatureKelvin float64 `json:"temp"`
			} `json:"main"`
		}{
			Name: data.Name,
			Main: struct {
				TemperatureKelvin float64 `json:"temp"`
			}{TemperatureKelvin: data.Main.TemperatureKelvin},
		}

		json.NewEncoder(w).Encode(responseData)
	})

	http.ListenAndServe(":9000", nil)
}
