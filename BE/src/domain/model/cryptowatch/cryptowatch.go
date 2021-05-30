package cryptowatch

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Allowance struct {
	Cost      int64
	Remaining int64
}
type Market struct {
	Exchange     string `json:"exchange"`
	CurrencyPair string `json:"currencyPair"`
}

func (m *Market) toString() string {
	return m.Exchange + ":" + m.CurrencyPair
}

// type MarketsPrices map[string]float64

// type MarketsSummaries map[string]Summary

// type Price struct {
// 	Last   float64 `json:"last"`
// 	High   float64 `json:"high"`
// 	Low    float64 `json:"low"`
// 	Change Change `json:"change"`
// }

// type Change struct {
// 	Percentage float64 `json:"percentage"`
// 	Absolute   float64 `json:"absolute"`
// }

// type Summary struct {
// 	Price  Price `json:"price"`
// 	Volume float64        `json:"volume"`
// }

// type Trade struct {
// 	ID        int64 `json:"id"`
// 	Timestamp int64 `json:"timestamp"`
// 	Time      time.Time
// 	Price     float64 `json:"price"`
// 	Amount    float64`json:"amout"`
// }

// type Book struct {
// 	Price  float64 `json:"price"`
// 	Amount float64 `json:"amount"`
// }

// type OrderBook struct {
// 	Asks []Book `json:"asks"`
// 	Bids []Book `json:"bids"`
// }

type OHLC map[int64][]OHLCSummary

type OHLCSummary struct {
	CloseTimestamp int64
	CloseTime      time.Time
	OpenPrice      float64
	HighPrice      float64
	LowPrice       float64
	ClosePrice     float64
	Volume         float64
}

type OHLCParams struct {
	Before  int64
	After   int64
	Periods []int64
}

func (p *OHLCParams) ToValues() url.Values {
	v := url.Values{}
	if p.Before > 0 {
		v.Add("before", strconv.Itoa(int(p.Before)))
	}
	if p.After > 0 {
		v.Add("after", strconv.Itoa(int(p.After)))
	}
	if len(p.Periods) > 0 {
		s := make([]string, len(p.Periods))
		for i, j := range p.Periods {
			s[i] = strconv.Itoa(int(j))
		}
		v.Add("periods", strings.Join(s, ","))
	}
	return v
}
