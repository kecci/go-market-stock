package iqplus

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type MarketStock struct {
	StockCode       string  `json:"stock_code"`
	CompanyName     string  `json:"company_name"`
	LastTradedPrice float64 `json:"last_traded_price"`
}

type (
	Config struct {
		ServerHost, ServerPort, Username, Password string
	}
)

// ChangePassword changes the password
func ChangePassword(newPassword string, config Config) error {
	// Connect
	conn, err := net.Dial("tcp", net.JoinHostPort(config.ServerHost, config.ServerPort))
	if err != nil {
		return err
	}

	// Change Password Request
	// IQP|auth_record_type|sub type|encryption_method|user|new password|old password
	msg := fmt.Sprintf("IQP|149|1|1|%s|%s|%s\r\n", config.Username, newPassword, config.Password)
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return err
	}

	// Reader
	reader := bufio.NewReader(conn)
	reader.ReadLine()

	// Login Check
	err = isChangePasswordSuccess(reader)
	if err != nil {
		return err
	}

	return nil
}

// isLogin checks if the login is successful
func isLogin(reader *bufio.Reader) error {

	for {
		lineByte, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		line := string(lineByte)[1:]

		// IQP|149|0|0|OK[CR/LF]
		if strings.Contains(line, "IQP|149|0|") {
			if strings.Contains(line, "OK") {
				return nil
			}
			return errors.New(line)
		}
	}
}

// isLogin checks if the login is successful
func isChangePasswordSuccess(reader *bufio.Reader) error {
	lineByte, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	line := string(lineByte)[1:]

	// IQP|149|1|0|OK[CR/LF]
	if !strings.Contains(line, "IQP|149|1|0|") {
		return errors.New(line)
	}

	println(line)
	return nil
}

// MapToMarketStock converts a map to MarketStock
func MapToMarketStock(line string) *MarketStock {
	quoteArray := strings.Split(line, "|")

	if ReadRecordType(line) == Quote {

		var stockCode, companyName string
		var lastTradedPrice float64

		for i := range quoteArray {
			if strings.HasPrefix(quoteArray[i], "0;") && !strings.Contains(quoteArray[i], "-") {
				stockCode = strings.ReplaceAll(quoteArray[i], "0;", "")
			} else if strings.HasPrefix(quoteArray[i], "1;") {
				companyName = strings.ReplaceAll(quoteArray[i], "1;", "")
			} else if strings.HasPrefix(quoteArray[i], "56;") {
				lastTradedPriceStr := strings.ReplaceAll(quoteArray[i], "56;", "")
				lastTradedPriceFloat, err := strconv.ParseFloat(lastTradedPriceStr, 64)
				if err != nil {
					break
				}
				lastTradedPrice = lastTradedPriceFloat
			}
		}
		if stockCode != "" && companyName != "" && lastTradedPrice != 0 {
			return &MarketStock{
				StockCode:       stockCode,
				CompanyName:     companyName,
				LastTradedPrice: lastTradedPrice,
			}
		}
		return nil
	}
	return nil
}
