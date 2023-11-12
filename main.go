package main

import (
	"flag"
	"sp/apis"
	"strings"

	"github.com/fatih/color"
)

var USD_rate float64 = 1

func main() {

	currency_flag := flag.String("c", "USD", "3 letters currency name. Display price in that currency.")
	flag.Parse()

	currency := strings.ToUpper(*currency_flag)

	if len(currency) != 3 {
		color.Yellow("Sorry stepper, your currency must have exactly 3 letter. Example, it ain't EUROS but EUR lol :)")
		return
	}

	if currency != "USD" {
		USD_rate = apis.FetchUSDConversion(currency)
	}

	if USD_rate == 0 {
		color.Yellow("Sorry stepper, can't get current price in " + currency + " rn :(\n")
		color.Yellow("Try again later ! (Or make sure you did not misspelled your currency lol)\n")
		color.Yellow("Will try to fetch in USD ...\n")
		currency = "USD"
		USD_rate = 1
	}

	apis.FetchStepPrices(USD_rate, currency)

}
