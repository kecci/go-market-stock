# go-market-stock

![alt](/assets/banner.png)

## Use Library
```go
    stock, _ := idx.GetMarketStockByCode("BBCA")
	stockByte, _ := json.Marshal(stock)
	log.Println(string(stockByte))
```

## Run Example Project
```sh
go run ./idx/example/main.go
```

