package apis

import "github.com/tidwall/gjson"

func FetchFromCoingecko() (float64, string) {
	body, err := FetchBody("https://api.coingecko.com/api/v3/simple/price?ids=step-finance&vs_currencies=usd")

	if err != nil || body == "{}" {
		return 0.0, "CoinGecko"
	}

	return gjson.Get(body, "step-finance.usd").Float(), "CoinGecko"
}

func FetchFromCoinCap() (float64, string) {
	body, err := FetchBody("https://api.coincap.io/v2/assets/step-finance")

	if err != nil {
		return 0.0, "CoinCap"
	}

	return gjson.Get(body, "data.priceUsd").Float(), "CoinCap"
}

func FetchFromCoinpaprika() (float64, string) {
	body, err := FetchBody("https://api.coinpaprika.com/v1/tickers/step-step-finance")

	if err != nil {
		return 0.0, "CoinPaprika"
	}

	return gjson.Get(body, "quotes.USD.price").Float(), "CoinPaprika"
}

func FetchUSDConversion(currency string) float64 {
	body, err := FetchBody("https://open.er-api.com/v6/latest/USD")

	if err != nil {
		return 0.0
	}

	return gjson.Get(body, "rates."+currency).Float()

}
