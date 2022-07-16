package iqplus

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type (
	IqPlusConn interface {
		Close() error
		SetHanlder(fn func())
		Start()
		CheckCon(line string) error
		ReadLine() (string, error)
		ReadRecord() (RecordType, error)
		GetTradingStatus(string) (string, error)
	}

	iqPlusConnImpl struct {
		conn   net.Conn
		reader *bufio.Reader
		fn     func()
	}
)

// NewConnection creates a new connection
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

	return newIqPlusConn(conn, reader), nil
}

// newIqPlusConn creates a new iqPlusConn
func newIqPlusConn(conn net.Conn, reader *bufio.Reader) IqPlusConn {
	return &iqPlusConnImpl{conn: conn, reader: reader}
}

// Hanlder handles the connection
func (c *iqPlusConnImpl) SetHanlder(fn func()) {
	c.fn = fn
}

// Start starts the connection
func (c *iqPlusConnImpl) Start() {
	for {
		c.fn()
	}
}

// Close closes the connection
func (c *iqPlusConnImpl) Close() error {
	return c.conn.Close()
}

// CheckCon indicate whether the communication with the IQPlus central server is disconnected or not.
// IQP|20211222|072222|1|13|0[CR/LF]
// Data 0 = 'UP' atau 1='DOWN'
func (c iqPlusConnImpl) CheckCon(line string) error {
	lineArray := strings.Split(line, "|")
	if ReadRecordType(line) == ControlMessages && lineArray[5] != "0" {
		return errors.New("central server is disconnected")
	}
	return nil
}

// ReadLine returns a line
func (c iqPlusConnImpl) ReadLine() (string, error) {
	line, err := c.reader.ReadString('\n')
	if err != nil {
		return line, err
	}
	return strings.TrimSpace(line), nil
}

// ReadRecord returns a record
func (c iqPlusConnImpl) ReadRecord() (RecordType, error) {
	line, err := c.ReadLine()
	if err != nil {
		return "", err
	}
	return ReadRecordType(line), nil
}

// GetTradingStatus returns the trading status
// Type 57 : Trading Status
// '1' for Begin sending records '3' for Begin first session
// '4' for End first session
// '5' for Begin second session '6' for End second session
// '7' for End sending records '8' for Begin Pre-opening
// '9' for End Pre-opening
// 'a' for Begin Pre-closing
// 'b' for End Pre-closing
// 'c' for Begin Post-trading
// 'd' for End Post-trading
// `e' for trading suspension
// `f' for trading activation
// `g' for board suspension
// `h' for board activation
// `i' for Instrument suspension `j' for Instrument activation `k' for Market suspension
// `l' for Market activation
// Example: IQP|20211223|040351|3536|57|1|Begin sending records[CR/LF]
func (c iqPlusConnImpl) GetTradingStatus(string) (string, error) {
	return "", nil
}
