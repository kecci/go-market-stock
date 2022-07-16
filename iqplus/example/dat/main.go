package main

import (
	"bufio"
	"encoding/json"
	"os"
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

			var stockCode, companyName, lastClose string

			for i := range quoteArray {
				if strings.HasPrefix(quoteArray[i], "0;") && !strings.Contains(quoteArray[i], "-") {
					stockCode = strings.ReplaceAll(quoteArray[i], "0;", "")
				} else if strings.HasPrefix(quoteArray[i], "1;") {
					companyName = strings.ReplaceAll(quoteArray[i], "1;", "")
				} else if strings.HasPrefix(quoteArray[i], "56;") {
					lastClose = strings.ReplaceAll(quoteArray[i], "56;", "")
				}
			}

			if stockCode != "" && companyName != "" && lastClose != "" {
				stockMap[stockCode] = iqplus.MarketStock{
					StockCode:       stockCode,
					CompanyName:     companyName,
					LastTradedPrice: lastClose,
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
