package models

import (
	// "github.com/markcheno/go-talib"
	"time"
)

/*
	OHLC Candlesticks
	Period values

	https://docs.cryptowat.ch/rest-api/markets/ohlc

	60     1m
	180    3m
	300    5m
	900    15m
	1800   30m
	3600   1h
	7200   2h
	14400  4h
	21600  6h
	43200  12h
	86400  1d
	259200 3d
	604800 1w


	[
		CloseTime   float,
		OpenPrice   float,
		HighPrice   float,
		LowPrice    float,
		ClosePrice  float,
		Volume      float,
		QuoteVolume float,
    ]
*/

type CwCandles struct {
	Result struct {
		Num60     [][]float64 `json:"60"`
		Num180    [][]float64 `json:"180"`
		Num300    [][]float64 `json:"300"`
		Num900    [][]float64 `json:"900"`
		Num1800   [][]float64 `json:"1800"`
		Num3600   [][]float64 `json:"3600"`
		Num7200   [][]float64 `json:"7200"`
		Num14400  [][]float64 `json:"14400"`
		Num21600  [][]float64 `json:"21600"`
		Num43200  [][]float64 `json:"43200"`
		Num86400  [][]float64 `json:"86400"`
		Num259200 [][]float64 `json:"259200"`
		Num604800 [][]float64 `json:"604800"`
	} `json:"result"`
}

type Candle struct {
	Duration    time.Duration `json:"duration"`
	Closetime   float64       `json:"close_time"`
	Open        float64       `json:"open"`
	High        float64       `json:"high"`
	Low         float64       `json:"low"`
	Close       float64       `json:"close"`
	Volume      float64       `json:"volume"`
	Quotevolume float64       `json:"quote_volume"`
}
