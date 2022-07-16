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

	println("connected, waiting for data...")

	// Read Line
	for {
		line, err := conn.ReadLine()
		if err != nil {
			println(err.Error())
			break
		}
		err = conn.CheckCon(line)
		if err != nil {
			println(err.Error())
			break
		}
	}

	println("iqplus terminated")
}
