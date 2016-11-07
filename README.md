# Freyr

A command line tool to show weather from the [Dark Sky][1] API.

## Dark Sky API Key

To use Freyr a [Dark Sky][1] API key is required. To get one [sign up][2] and
retrieve the secret key from the [account page][3].

## Installation

    go get -v github.com/shapeshed/freyr

## Options

    --help    # show help
    -u        # set forecast.io units - us, uk, ca, si or auto. This may also be set by the FREYR_UNITS environment variable
    -k apikey # the forecast.io api key. This may also be set by the FREYR_KEY environment variable
    -l latlog # the long lat co-ordinates for the location. This may also be set by the FREYR_LATLON environment variable

## Usage

Get a [Dark Sky API key][1] and the [lat-long co-ordinates][4] for the
location you want.

    export FREYR_KEY=5a078b5225925e716eef0b83499358bc 
    export FREYR_LATLON=52.847875,-0.664397 
    freyr
    It is currently 11°C and Partly Cloudy

    freyr -k 5a078b5225925e716eef0b83499358bc -l 52.847875,-0.664397 freyr
    It is currently 11°C and Partly Cloudy


[1]: https://darksky.net/
[2]: https://darksky.net/dev/register
[3]: https://darksky.net/dev/account
[4]: http://dbsgeo.com/latlon/
