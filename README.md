# Freyr

[![Build Status](https://travis-ci.org/shapeshed/freyr.svg?branch=master)](https://travis-ci.org/shapeshed/freyr)

A command line client to show weather from the [Dark Sky][1] API.

## Dark Sky API Key

To use Freyr a [Dark Sky][1] API key is required. To get one [sign up][2] and
retrieve the secret key from the [account page][3].

## Installation

    go get -v github.com/shapeshed/freyr

## Usage

Get a [Dark Sky API key][1] and the [lat-long co-ordinates][4] for the
location you want.

Set environment variables for your Dark Sky Key, Latitude and Longitude.

    export FREYR_KEY=5a078b5225925e716eef0b83499358bc 
    export FREYR_LAT=40.759011
    export FREYR_LONG=-73.984472

### Current temperature

    freyr
    6 Mostly Cloudy

### Hourly Forecast

    freyr -h
    Hour            Temp    Summary
    22 Nov 21:00    6       Mostly Cloudy
    22 Nov 22:00    5       Mostly Cloudy
    22 Nov 23:00    5       Mostly Cloudy
    23 Nov 00:00    5       Mostly Cloudy
    23 Nov 01:00    5       Mostly Cloudy
    23 Nov 02:00    5       Mostly Cloudy
    ...

### Daily Forecast

    freyr -d
    Day     Min     Max     Summary
    22 Nov  5       8       Light rain in the morning.
    23 Nov  4       9       Mostly cloudy throughout the day.
    24 Nov  5       9       Partly cloudy throughout the day.
    25 Nov  4       10      Partly cloudy in the morning.
    26 Nov  3       8       Mostly cloudy throughout the day.
    27 Nov  1       9       Mostly cloudy throughout the day.
    28 Nov  0       6       Mostly cloudy until afternoon.
    29 Nov  -2      4       Flurries overnight.

### Alerts

    freyr -a

[1]: https://darksky.net/
[2]: https://darksky.net/dev/register
[3]: https://darksky.net/dev/account
[4]: http://dbsgeo.com/latlon/
