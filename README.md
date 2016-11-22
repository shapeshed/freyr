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

    export FREYR_KEY=5a078b5225925e716eef0b83499358bc 
    export FREYR_LAT=40.759011
    export FREYR_LONG=-73.984472
    freyr
    It is currently 37.46 and Mostly Cloudy

[1]: https://darksky.net/
[2]: https://darksky.net/dev/register
[3]: https://darksky.net/dev/account
[4]: http://dbsgeo.com/latlon/
