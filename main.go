package main

import (
	"fmt"
	"github.com/shapeshed/darksky"
	"log"
	"os"
	"strconv"
)

func main() {

	key := os.Getenv("FREYR_KEY")
	lat := os.Getenv("FREYR_LAT")
	long := os.Getenv("FREYR_LONG")
	units := os.Getenv("FREYR_UNITS")

	latInt, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		log.Fatal(err)
	}
	longInt, err := strconv.ParseFloat(long, 64)
	if err != nil {
		log.Fatal(err)
	}

	params := darksky.RequestParams{
		Key:       key,
		Latitude:  latInt,
		Longitude: longInt,
	}

	if units != "" {
		params.Units = units
	}

	forecast, err := darksky.Get(&params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("It is currently %v and %v\n",
		forecast.Currently.Temperature,
		forecast.Currently.Summary)
	os.Exit(0)

}
