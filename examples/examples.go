/*
Package main is a basic example to use the go-preev package
*/
package main

import (
	"fmt"
	"log"

	"github.com/mrz1836/go-preev"
)

func main() {
	// Create a new client
	client := preev.NewClient(nil)

	//
	// Get active pairs
	//
	pairs, err := client.GetPairs()
	if err != nil {
		log.Fatal("error: ", err.Error())
	} else if pairs == nil {
		log.Fatal("pairs was nil?")
	}

	fmt.Println("Found Active Pair(s):", pairs.BsvUsd.Name)

	//
	// Get the details of a specific pair
	//
	var pair *preev.Pair
	pair, err = client.GetPair(pairs.BsvUsd.ID)
	if err != nil {
		log.Fatal("error: ", err.Error())
	} else if pair == nil {
		log.Fatal("pair was nil?")
	}

	fmt.Println("Found BSV Pair:", pair.Name, "Quote:", pair.Base, "/", pair.Quote)

	//
	// Get active tickers
	//
	var tickers *preev.TickerList
	tickers, err = client.GetTickers()
	if err != nil {
		log.Fatal("error: ", err.Error())
	} else if tickers == nil {
		log.Fatal("tickers was nil?")
	}

	fmt.Println("Tickers Found! Current Price:", tickers.BsvUsd.Prices.Ppi.LastPrice, "Volume:", tickers.BsvUsd.Prices.Ppi.Volume)

	//
	// Get ticker by pair id
	//
	var ticker *preev.Ticker
	ticker, err = client.GetTicker(pair.ID)
	if err != nil {
		log.Fatal("error: ", err.Error())
	} else if ticker == nil {
		log.Fatal("ticker was nil?")
	}

	fmt.Println("Ticker Found! Current Price:", tickers.BsvUsd.Prices.Ppi.LastPrice, "Volume:", tickers.BsvUsd.Prices.Ppi.Volume)

	//
	// Get ticker history
	//
	var tickerHistory []*preev.Ticker
	tickerHistory, err = client.GetTickerHistory(pair.ID, 1592950680, 1592951700, 10)
	if err != nil {
		log.Fatal("error: ", err.Error())
	} else if len(tickerHistory) == 0 {
		log.Fatal("no results for ticker history?")
	}

	fmt.Println("Ticker History Found! Timestamp:", tickerHistory[0].Timestamp, "Price:", tickerHistory[0].Prices.Ppi.LastPrice)
}
