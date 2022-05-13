package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/kecci/go-market-stock/idx"
)

func main() {
	start := time.Now()
	// Get market stocks
	stocks, err := idx.GetMarketStocks()
	if err != nil {
		log.Println(err.Error())
		return
	}
	stocksByte, _ := json.MarshalIndent(stocks, "", "    ")
	log.Println(string(stocksByte))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	// // Get market index
	// indexes, _ := idx.GetMarketIndexes()
	// b, _ := json.MarshalIndent(indexes, "", "    ")
	// log.Println(string(b))

	// Get market stock by code
	// stock, _ := idx.GetMarketStockByCode("BBCA")
	// stockByte, _ := json.MarshalIndent(stock, "", "    ")
	// log.Println(string(stockByte))

	// Get market index by code
	// index, _ := idx.GetMarketIndexByCode("IHSG")
	// b, _ := json.Marshal(index)
	// log.Println(string(b))
}
