package usecase

import (
	"encoding/json"
)

type TickerUsecase struct{}

func NewTickerUsecase() *TickerUsecase {
	return &TickerUsecase{}
}

func (t *TickerUsecase) TestCryptowatch(periods, beforeAfter, unitTimeStamp string) (*Cryptowatch, error) {
	apiClient := New("", "")
	return apiClient.GetCryptowatchTest(periods, beforeAfter, unitTimeStamp)
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
*/
type Cryptowatch struct {
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

func (api *APIClient) GetCryptowatchTest(
	periods string,
	beforeAfter string,
	time string,
) (*Cryptowatch, error) {
	url := ""
	resp, err := api.doRequest(
		"GET",
		url,
		map[string]string{
			"periods":   periods,
			beforeAfter: time,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	var cryptowatch Cryptowatch
	err = json.Unmarshal(resp, &cryptowatch)
	if err != nil {
		return nil, err
	}
	return &cryptowatch, nil
}
