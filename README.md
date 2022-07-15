# go-market-stock
List of SDK Market Stock Libraries.
| Source | Method | Data | Website |
| ------ | ------ | ---- | ------ |
| IDX | Scraping Web | Historical Data | https://www.idxchannel.com/market-stock |
| IQPlus | Dial TCP | Realtime Data | http://www.iqplus.info/produk/data_feed/


## IDX
Example Data
![alt](/assets/banner.png)
### Usage
```go
package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/kecci/go-market-stock/idx"
)

func main() {
    stock, _ := idx.GetMarketStockByCode("BBCA")
	stockByte, _ := json.Marshal(stock)
	log.Println(string(stockByte))
}
```

### Run Example
```sh
go run ./idx/example/main.go
```

## IQPlus
Docs: https://github.com/kecci/go-market-stock/tree/master/iqplus