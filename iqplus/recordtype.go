package iqplus

import (
	"strings"
)

type RecordType string

const (
	HistoryWeekly   RecordType = "11"
	ControlMessages RecordType = "13"
	Quote           RecordType = "14"
	Trade           RecordType = "15"
	Order           RecordType = "16"
	Top20           RecordType = "17"
	BestQuote       RecordType = "18"
	ResendOrder     RecordType = "26"
	ResendTrade     RecordType = "27"
	HistoryDaily    RecordType = "32"
	News            RecordType = "36"
	CPCA            RecordType = "38"
	Activity        RecordType = "39"
	TradeDone       RecordType = "40"
	HistoryMonthly  RecordType = "56"
	TradingStatus   RecordType = "57"
	NBSStock        RecordType = "58"
	NBSBroker       RecordType = "59"
	TradingSummary  RecordType = "130"
)

// String returns the string representation of the record type
func (r RecordType) String() string {
	return string(r)
}

// ReadRecord checks if the record type is valid
// IQP | Date | Time | Sequence # | Record Type | Data | CR/LF
func ReadRecord(line string) RecordType {
	lineArray := strings.Split(line, "|")
	if len(lineArray) < 6 {
		return ""
	}

	switch lineArray[4] {
	case "11":
		return HistoryWeekly
	case "13":
		return ControlMessages
	case "14":
		return Quote
	case "15":
		return Trade
	case "16":
		return Order
	case "17":
		return Top20
	case "18":
		return BestQuote
	case "26":
		return ResendOrder
	case "27":
		return ResendTrade
	case "32":
		return HistoryDaily
	case "36":
		return News
	case "38":
		return CPCA
	case "39":
		return Activity
	case "40":
		return TradeDone
	case "56":
		return HistoryMonthly
	case "57":
		return TradingStatus
	case "58":
		return NBSStock
	case "59":
		return NBSBroker
	case "130":
		return TradingSummary
	default:
		return ""
	}
}
