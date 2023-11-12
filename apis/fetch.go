package apis

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
)

func FetchBody(apiURL string) (string, error) {
	client := resty.New()

	response, err := client.R().
		Get(apiURL)

	if err != nil {
		return "", err
	}

	if response.StatusCode() != 200 {
		return "", errors.New("request failed with status code")
	}

	return response.String(), nil

}

func FetchStepPrices(usd_rate float64, currency string) {

	fetchingFunctions := []func() (float64, string){
		func() (float64, string) { return FetchFromCoingecko() },
		func() (float64, string) { return FetchFromCoinCap() },
		func() (float64, string) { return FetchFromCoinpaprika() },
	}

	for _, fetchFunc := range fetchingFunctions {
		result, source := fetchFunc()

		if result != 0.0 {
			fmt.Println("👉 " + fmt.Sprintf("%f", result*usd_rate) + " " + currency + " (" + source + ")")
			return
		}

	}

	color.Yellow("Sorry stepper, but fam can't fetch step price rn :(\nPlz try again later ! (Or make sure you're connected to the internet lol)")
}
