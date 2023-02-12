package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey`
}

type WeatherForecast struct {
	Name string `json:"Name"`
	Main struct {
		Temp float64 `json:"Temp"`
	} `json:"Main"`
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// loads API key from .apiConfig file
func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the Weather Console\n</h1>")
}

func weatherFilter(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	data, err := query(city)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// function to retrieve weather data based on the Country
func query(city string) (WeatherForecast, error) {
	apiConfig, err := loadApiConfig(".apiConfig")

	if err != nil {
		return WeatherForecast{}, err
	}

	//Weather API (parameters: API key, Unit of temperature (Celsius), Country)
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&units=metric" + "&q=" + city)

	if err != nil {
		return WeatherForecast{}, err
	}
	defer resp.Body.Close()

	var d WeatherForecast
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return WeatherForecast{}, err
	}
	return d, nil
}
