package main

import (
	"github.com/kecci/go-market-stock/iqplus"
)

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
	}
	defer conn.Close()

	conn.SetHanlder(func() error {
		line, err := conn.ReadLine()
		if err != nil {
			println(err.Error())
			return err
		}
		err = conn.CheckCon(line)
		if err != nil {
			println(err.Error())
			return err
		}
		println(line)
		return nil
	})

	println("connected, waiting for data...")

	conn.Start()

	println("iqplus terminated")
}
