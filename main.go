package main

import (
	"flag"
	"fmt"
	"github.com/shapeshed/darksky"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	h := flag.Bool("h", false, "Show hourly data")
	d := flag.Bool("d", false, "Show daily data")
	a := flag.Bool("a", false, "Show alerts")
	flag.Parse()

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

	if *h {
		for i, hour := range forecast.Hourly.Data {
			if i == 0 {
				fmt.Printf("%v\t\t%v\t%v\n", "Hour", "Temp", "Summary")
			}
			t := time.Unix(hour.Time, 0)
			fmt.Printf("%v\t%.0f\t%v\n", t.Format("2 Jan 15:04"), hour.Temperature, hour.Summary)

		}
		os.Exit(0)
	}

	if *d {
		for i, day := range forecast.Daily.Data {
			if i == 0 {
				fmt.Printf("%v\t%v\t%v\t%v\n", "Day", "Min", "Max", "Summary")
			}
			t := time.Unix(day.Time, 0)
			fmt.Printf("%v\t%.0f\t%.0f\t%v\n", t.Format("2 Jan"), day.TemperatureMin, day.TemperatureMax, day.Summary)

		}
		os.Exit(0)
	}

	if *a {
		if len(forecast.Alerts) == 0 {
			fmt.Println("No alerts")
		} else {
			for _, alert := range forecast.Alerts {
				t := time.Unix(alert.Time, 0)
				fmt.Printf("%v\n%.0f\n%v\n", t.Format("2 Jan 15:04"), alert.Title, alert.Description)

			}
		}
		os.Exit(0)
	}

	fmt.Printf("%.0f %v\n",
		forecast.Currently.Temperature,
		forecast.Currently.Summary)

}
