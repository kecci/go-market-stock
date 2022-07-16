package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/kecci/go-market-stock/iqplus"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	stockMap := make(map[string]iqplus.MarketStock)

	readFile, err := os.Open("iqplus/example/dat/dat")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		quoteArray := strings.Split(fileScanner.Text(), "|")

		if iqplus.ReadRecordType(line) == iqplus.Quote {

			var stockCode, companyName string
			var lastTradedPrice float64

			for i := range quoteArray {
				if strings.HasPrefix(quoteArray[i], "0;") && !strings.Contains(quoteArray[i], "-") {
					stockCode = strings.ReplaceAll(quoteArray[i], "0;", "")
				} else if strings.HasPrefix(quoteArray[i], "1;") {
					companyName = strings.ReplaceAll(quoteArray[i], "1;", "")
				} else if strings.HasPrefix(quoteArray[i], "56;") {
					lastTradedPriceStr := strings.ReplaceAll(quoteArray[i], "56;", "")
					lastTradedPrice, err = strconv.ParseFloat(lastTradedPriceStr, 64)
					if err != nil {
						break
					}
				}
			}

			if stockCode != "" && companyName != "" && lastTradedPrice != 0 {
				stockMap[stockCode] = iqplus.MarketStock{
					StockCode:       stockCode,
					CompanyName:     companyName,
					LastTradedPrice: lastTradedPrice,
				}
			}
		}
	}

	readFile.Close()

	for i := range stockMap {
		b, _ := json.Marshal(stockMap[i])
		println(string(b))
	}
}
