package models

import (
	// "github.com/markcheno/go-talib"
	// "time"
)

type DataFrameCandle struct {
	Candles CwCandles `json:"candles"`
	Smas    []Sma     `json:"smas,omitempty"`
}

type Sma struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}


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

// func (df *DataFrameCandle) Times() []time.Time {
// 	s := make([]time.Time, len(df.Candles))
// 	for i, candle := range df.Candles {
// 		s[i] = candle.Time
// 	}
// 	return s
// }

// func (df *DataFrameCandle) Opens() []float64 {
// 	s := make([]float64, len(df.Candles))
// 	for i, candle := range df.Candles {
// 		s[i] = candle.Open
// 	}
// 	return s
// }

// func (df *DataFrameCandle) Closes() []float64 {
// 	s := make([]float64, len(df.Candles))
// 	for i, candle := range df.Candles {
// 		s[i] = candle.Close
// 	}
// 	return s
// }

// func (df *DataFrameCandle) Highs() []float64 {
// 	s := make([]float64, len(df.Candles))
// 	for i, candle := range df.Candles {
// 		s[i] = candle.High
// 	}
// 	return s
// }

// func (df *DataFrameCandle) Low() []float64 {
// 	s := make([]float64, len(df.Candles))
// 	for i, candle := range df.Candles {
// 		s[i] = candle.Low
// 	}
// 	return s
// }

// func (df *DataFrameCandle) Volume() []float64 {
// 	s := make([]float64, len(df.Candles))
// 	for i, candle := range df.Candles {
// 		s[i] = candle.Volume
// 	}
// 	return s
// }

// func (df *DataFrameCandle) AddSma(period int) bool {
// 	if len(df.Candles) > period {
// 		df.Smas = append(df.Smas, Sma{
// 			Period: period,
// 			Values: talib.Sma(df.Closes(), period),
// 		})
// 		return true
// 	}
// 	return false
// }
