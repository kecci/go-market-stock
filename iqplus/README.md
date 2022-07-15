# IQPlus SDK Library with Go

IQPlus Data Feed (SCF) is stock market trading data streaming information, especially on the Indonesia Stock Exchange (IDX ) which has been processed in the form of text and sent (updated) to the receiving server with a certain data protocol.

The communication protocol used is TCP/IP (client-server) with the UNIX FreeBSD operating system. The use of TCP/IP is intended to make it easier for customers to receive required data.

## RECORD TYPE

| Record Type |	Record Name |
| ----------- | ----------- |
| 11 | History  Weekly |
| 13 | Control  Messages |
| 14 | Quote  |
| 15 | Trade  |
| 16 | Order  |
| 17 | Top  20 |
| 18 | Best  Quote |
| 26 | Resend  Order |
| 27 | Resend  Trade |
| 32 | History  Daily |
| 36 | News  |
| 38 | CPCA  |
| 39 | Activity  |
| 40 | Trade  Done |
| 56 | History  Monthly |
| 57 | Trading  Status |
| 58 | NBS  Stock |
| 59 | NBS  Broker |
| 130 | Trading  Summary |
| 149 (0) |	Login |
| 149 (1) |	Change Password |

## Data
| Data |	Record Type | 	Information |
| ---- | -------------- | ------------- |
| IDX Stock | 13 Control Messages 14 Quote 15 Trade 16 Order 17 Top 20 18 Best Quote 26 Resend Order 27 Resend Trade 39 Activity 40 Trade Done 57 Trading Status 130 Trading Summary 149 Login Historical ( daily, weekly, monthly) |	Permission Required |
| IDX Broker |	15 Trade ( inc. broker codes) 17 Top 20 Broker 27 Resend Trade 58 NBS Stock  59 NBS Broker | Permission Required |
| Regional Index | 13 Control Messages 14 Quote 149 Login Historical ( daily, weekly, monthly)	| Permission Required
| Commodity |	13 Control Messages 14 Quote 149 Login Historical ( daily, weekly, monthly) |	Permission Required |
| Futures |	13 Control Messages 14 Quote 149 Login Historical ( daily, weekly, monthly) | Permission Required
| Currency | 13 Control Messages 14 Quote 149 Login Historical ( daily, weekly, monthly) | Permission Required
| News |	13 Control Messages 36 News 149 Login |	Permission Required |

## FID
Is the main database of the IQPlus datafeed which consists of several FIDs (field identification numbers). Between FID separated by "|" and between FID and its data value separated by separator unit ';' (semicolon)

| FID | Type |	Description |
| --- | ---- | ------------ |
| 0 |	String |	Code |
| 1 |	String |	Name |
| 2 |	Numeric |	XGroup |
| 3 |	Byte |	Status |
| 4 |	Numeric |	Listing Date |
| 5 |	Numeric |	Group |
| 6 |	Numeric |	ISSUERID |
| 7 |	Numeric |	Listed Shares |
| 8 |	Numeric |	Tradable Shares |
| 9 |	Numeric |	IPO |
| 10 |	Numeric |	ORDERBOOKID |
| 11 |	Numeric |	Base Price |
| 12 |	String |	CURRENCY |
| 13 |	String |	Remark |
| 14 |	Numeric |	Earning per share (EPS) |
| 15 |	String |	ISIN |
| 16 |	Numeric |	FOREIGNLIMIT |
| 17 |	Numeric |	SECTORNAME |
| 18 |	Numeric |	INDUSTRYNAME |
| 19 |	Numeric |	EXPIRYDATE |
| 20 |	String |	UNDERLYING |
| 21 |	Numeric |	NTA (Net tangible asset) |
| 22 |	Numeric |	CONTRACTSIZE |
| 23 |	String |	Under writer |
| 24 |	Numeric |	Bid price |
| 25 |	Byte |	VERB |
| 26 |	Numeric |	STRIKEPRICE |
| 27 |	Numeric |	High bid price |
| 28 |	Numeric |	WEIGHT |
| 29 |	Numeric |	Low bid price |
| 30 |	Numeric |	MATURITYDATE |
| 31 |	Numeric |	Bid Volume |
| 32 |	Numeric |	XGROUP1 |
| 33 |	String |	SECURITYTYPE |
| 34 |	Numeric |	FLAG |
| 35 |	Numeric |	INDICATOR |
| 36 |	Numeric |	RSI |
| 37 |	Numeric |	HIGH5 |
| 38 |	Numeric |	Number of offer orders |
| 39 |	Numeric |	Offer price |
| 40 |	Numeric |	LOW5 |
| 41 |	Numeric |	INDEX |
| 42 |	Numeric |	High offer price |
| 43 |	Numeric |	FRGBOUGHTFREQ |
| 44 |	Numeric |	Low offer price |
| 45 |	Numeric |	DOMBOUGHTFREQ |
| 46 |	Numeric |	Offer Volume |
| 47 |	Numeric |	FRGSOLDFREQ |
| 48 |	Numeric |	DOMSOLDFREQ |
| 49 |	Numeric |	FRGBOUGHTVOL |
| 50 |	Numeric |	DOMBOUGHTVOL |
| 51 |	Numeric |	FRGSOLDVOL |
| 52 |	Numeric |	DOMSOLDVOL |
| 53 |	Numeric |	Number of bid orders |
| 54 |	Numeric |	Open price |
| 55 |	Numeric |	THEORETICALPRC |
| 56 |	Numeric |	Last traded price |
| 57 |	Numeric |	High traded price |
| 58 |	Numeric |	THEORETICALVOL |
| 59 |	Numeric |	Low trade price |
| 60 |	Numeric |	CLOSE |
| 61 |	Numeric |	Close Date |
| 62 |	Numeric |	TRDVOL |
| 63 |	Numeric |	TRDVAL |
| 64 |	Numeric |	TOTFREQ |
| 65 |	Numeric |	XBASEVAL |
| 66 |	Numeric |	XMARKETVAL |
| 67 |	Numeric |	CHANGE |
| 68 |	Numeric |	RATIO |
| 69 |	Numeric |	RECDATE |
| 70 |	String |	BOARD |
| 71 |	Subst |	SOURCE |
| 72 |	Numeric |	VOL |
| 73 |	Numeric |	SHARELOT |
| 74 |	Numeric |	FRGBOUGHTVAL |
| 75 |	Numeric |	FRGSOLDVAL |
| 76 |	Numeric |	DOMBOUGHTVAL |
| 77 |	Numeric |	DOMSOLDVAL |
| 78 |	Numeric |	AVG |
| 79 |	Numeric |	PCTCHANGE |

## 1.3 Quote
Data type 14 = Quote

Quote Is the main database of the IQPlus datafeed which consists of several FIDs (field identification numbers). Between FID separated by "|"and between FID and its data value separated by separator unit ';' (semicolon).

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
- IQPlus Data Feed Service (SCF) - Technical Specification version 4.0.0 (beta)
- Website: http://www.iqplus.info/produk/data_feed/