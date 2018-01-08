package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pdepip/go-binance/binance"
)

var (
	apikey    = os.Getenv("BINANCE_API_KEY")
	apisecret = os.Getenv("BINANCE_API_SECRET")
)

func getBalance(symbol string) (float64, error) {
	client := binance.New(apikey, apisecret)

	balances, err := client.GetPositions()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	if len(balances) == 0 {
		return 0, fmt.Errorf("no balances found, likely wrong api key/secret")
	}

	for i := range balances {
		if balances[i].Asset == symbol {
			return balances[i].Free, nil
		}
	}

	return 0, nil
}

func main() {
	if apikey == "" || apisecret == "" {
		log.Fatalf("binance api key and/or secret not set")
	}

	symbol := flag.String("symbol", "BNB", "coin symbol")
	treshold := flag.Float64("treshold", 0.0, "treshold below which to complain")
	verbose := flag.Bool("verbose", false, "verbosity")
	flag.Parse()

	value, err := getBalance(*symbol)
	if err != nil {
		log.Fatal(err)
	}

	if value <= *treshold {
		fmt.Printf("%v balance %f below or equal to treshold %f\n", *symbol, value, *treshold)
	} else if *verbose {
		fmt.Printf("%v balance %f above treshold %f\n", *symbol, value, *treshold)
	}

}
