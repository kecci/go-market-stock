package idx

import (
	"bytes"
	"errors"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	gomarketstock "github.com/kecci/go-market-stock"
)

const (
	IDX_MARKET_STOCK = "https://www.idxchannel.com/market-stock"
)

type (
	MarketStockModel struct {
		Index             string `json:"index"`
		StockCode         string `json:"stock_code"`
		CompanyName       string `json:"company_name"`
		PreviousClose     string `json:"previous_close"`
		LastClose         string `json:"last_close"`
		ChangesPrice      string `json:"changes_price"`
		ChangesPercentage string `json:"changes_percentage"`
	}

	MarketIndexModel struct {
		IndexCode         string `json:"index_code"`
		LastClose         string `json:"last_close"`
		ChangesPrice      string `json:"changes_price"`
		ChangesPercentage string `json:"changes_percentage"`
	}
)

// GetMarketStocks returns market stocks
func GetMarketStocks() (res []MarketStockModel, err error) {
	data, err := gomarketstock.GetDataFromURL(IDX_MARKET_STOCK)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(data))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	doc.Find("tbody").Find("tr").Each(func(i int, sel *goquery.Selection) {
		marketStockModel := MarketStockModel{}
		sel.Find("td").Each(func(i int, sel *goquery.Selection) {
			switch i {
			case 0:
				marketStockModel.Index = strings.Replace(sel.Text(), ".", "", -1)
			case 1:
				marketStockModel.StockCode = sel.Text()
			case 2:
				marketStockModel.CompanyName = sel.Text()
			case 3:
				marketStockModel.PreviousClose = sel.Text()
			case 4:
				marketStockModel.LastClose = sel.Text()
			case 5:
				marketStockModel.ChangesPrice = sel.Text()
			case 6:
				marketStockModel.ChangesPercentage = sel.Text()
			}
		})

		res = append(res, marketStockModel)
	})

	return res, nil
}

// GetMarketStock returns market stock
func GetMarketStockByCode(code string) (res MarketStockModel, err error) {
	stocks, err := GetMarketStocks()
	if err != nil {
		log.Println(err.Error())
		return res, err
	}

	for i := range stocks {
		if stocks[i].StockCode == code {
			return stocks[i], nil
		}
	}

	return res, errors.New("stock not found")
}

// GetMarketIndexes returns market indexes
func GetMarketIndexes() (res []MarketIndexModel, err error) {
	data, err := gomarketstock.GetDataFromURL(IDX_MARKET_STOCK)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(data))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	doc.Find(".market-list").Find("div .market-list--row").Children().Each(func(i int, sel *goquery.Selection) {
		if sel.Find("div").First().Text() != "" {

			marketIndexModel := MarketIndexModel{}
			marketIndexModel.IndexCode = sel.Find("div").Find("a").Text()
			marketIndexModel.LastClose = sel.Find("div").Next().Find("div").First().Text()
			marketIndexModel.ChangesPercentage = sel.Find("div").Next().Find("div").Next().Find("div").First().Text()
			marketIndexModel.ChangesPrice = sel.Find("div").Next().Next().Text()

			res = append(res, marketIndexModel)
		}
	})

	return res, nil
}

// GetMarketIndexByCode returns market index
func GetMarketIndexByCode(code string) (res MarketIndexModel, err error) {
	indexes, err := GetMarketIndexes()
	if err != nil {
		return res, err
	}

	for i := range indexes {
		if indexes[i].IndexCode == code {
			return indexes[i], nil
		}
	}

	return res, errors.New("index not found")
}
