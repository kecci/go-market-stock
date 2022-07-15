package iqplus

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type MarketStock struct {
	StockCode       string `json:"stock_code"`
	CompanyName     string `json:"company_name"`
	LastTradedPrice string `json:"last_traded_price"`
}

type (
	Config struct {
		ServerHost, ServerPort, Username, Password string
	}

	IqPlusConn interface {
		Close() error
		ReadLine() (string, error)
	}

	iqPlusConnImpl struct {
		conn   net.Conn
		reader *bufio.Reader
	}
)

// Close closes the connection
func (c iqPlusConnImpl) Close() error {
	return c.conn.Close()
}

// ReadLine returns a line
func (c iqPlusConnImpl) ReadLine() (string, error) {
	return c.reader.ReadString('\n')
}

// Connect returns a connection and a reader
func NewConnection(config Config) (IqPlusConn, error) {
	// Connect
	conn, err := net.Dial("tcp", net.JoinHostPort(config.ServerHost, config.ServerPort))
	if err != nil {
		return nil, err
	}

	// Login
	msg := fmt.Sprintf("IQP|149|0|1|%s|%s\r\n", config.Username, config.Password)
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return nil, err
	}

	// Reader
	reader := bufio.NewReader(conn)
	reader.ReadLine()

	// Login Check
	err = isLogin(reader)
	if err != nil {
		return nil, err
	}

	return iqPlusConnImpl{conn: conn, reader: reader}, nil
}

// isLogin checks if the login is successful
func isLogin(reader *bufio.Reader) error {
	lineByte, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	line := string(lineByte)[1:]

	if !strings.Contains(line, "IQP|149|0|0|") {
		return errors.New(line)
	}
	return nil
}

// MapToMarketStock converts a map to MarketStock
func MapToMarketStock(line string) *MarketStock {
	quoteArray := strings.Split(line, "|")

	if len(quoteArray) > 5 && quoteArray[4] == "14" {

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

		if stockCode == "" {
			return nil
		}

		return &MarketStock{
			StockCode:       stockCode,
			CompanyName:     companyName,
			LastTradedPrice: lastClose,
		}
	}
	return nil
}
