package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const darkSkyAPI = "https://api.darksky.net/forecast/"

type WeatherResponse struct {
	Currently struct {
		Summary     string  `json:"summary"`
		Temperature float64 `json:"temperature"`
	}
}

func darkSkyURL(x string, y string, z string) string {
	var requestURL bytes.Buffer
	requestURL.WriteString(darkSkyAPI)
	requestURL.WriteString(x)
	requestURL.WriteString("/")
	requestURL.WriteString(y)
	requestURL.WriteString("/")
	requestURL.WriteString("?units=")
	requestURL.WriteString(z)
	requestURL.WriteString("&exclude=minutely,hourly,daily,alerts,tags")
	return requestURL.String()
}

func setVariable(key, envVariable string) (string, error) {
	envKey := os.Getenv(envVariable)
	if len(key) == 0 && len(envKey) == 0 {
		return "", errors.New("not set. Use --help for help.")
	}
	if len(key) == 0 {
		return envKey, nil
	}
	return key, nil
}

func main() {

	var apiKeyArg, weatherUnitsArg, latLongArg string
	flag.StringVar(&apiKeyArg, "k", "", "Dark Sky API Key")
	flag.StringVar(&weatherUnitsArg, "u", "si", "Weather units")
	flag.StringVar(&latLongArg, "l", "", "Latitude and Longitude co-ordinates")
	flag.Parse()

	key, err := setVariable(apiKeyArg, "FREYR_KEY")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Dark Sky API Key", err)
		os.Exit(1)
	}

	latLong, err := setVariable(latLongArg, "FREYR_LATLON")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Latitude and longitude coordinates", err)
		os.Exit(1)
	}

	units, err := setVariable(weatherUnitsArg, "FREYR_UNITS")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Weather units", err)
		os.Exit(1)
	}

	var darkSkyRequest = darkSkyURL(key, latLong, units)

	var weatherResponse WeatherResponse

	weatherRes, err := http.Get(darkSkyRequest)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if weatherRes.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, "Unexpected status code", weatherRes.StatusCode)
		os.Exit(1)
	}
	defer weatherRes.Body.Close()

	decoder := json.NewDecoder(weatherRes.Body)

	err = decoder.Decode(&weatherResponse)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Printf("It is currently %v°C and %v\n", weatherResponse.Currently.Temperature, weatherResponse.Currently.Summary)
	os.Exit(0)
}
