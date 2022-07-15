# IQPlus SDK Library with Go

IQPlus Data Feed (SCF) is stock market trading data streaming information, especially on the Indonesia Stock Exchange (IDX ) which has been processed in the form of text and sent (updated) to the receiving server with a certain data protocol.

The communication protocol used is TCP/IP (client-server) with the UNIX FreeBSD operating system. The use of TCP/IP is intended to make it easier for customers to receive required data.

## Usage
```go
package main

import "github.com/kecci/go-market-stock/iqplus"

func main() {
	// Config
	cfg := iqplus.Config{
		ServerHost: "xxx.xxx.xxx.xxx",
		ServerPort: "xxxx",
		Username:   "xxxx",
		Password:   "xxxxxxxxxxxxxxxx",
	}

	// Connect
	conn, err := iqplus.NewConnection(cfg)
	if err != nil {
		println(err.Error())
		return
	}
	defer conn.Close()

	// Read Line
	for {
		line, err := conn.ReadLine()
		if err != nil {
			println(err.Error())
			break
		}
		println(line)
	}

	println("iqplus terminated")
}
```

## Source
- Website: http://www.iqplus.info/produk/data_feed/