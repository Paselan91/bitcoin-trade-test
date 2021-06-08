package usecase

import (
	"app/src/domain/models"
	"encoding/json"
	"log"
)

type TickerUsecase struct{}

func NewTickerUsecase() *TickerUsecase {
	return &TickerUsecase{}
}

func (t *TickerUsecase) FetchDataFlameCandles(periods string, beforeAfter string, unitTimeStamp string, isSma bool, smas []int) (*models.DataFrameCandle, error) {
	apiClient := New("", "")
	cwCandles, err := apiClient.FetchCwCandles(periods, beforeAfter, unitTimeStamp)
	if err != nil {
		log.Println(err)
	}
	candles := cwCandles.ConvertCwCandlesToCandles(periods)

	var df models.DataFrameCandle
	df.Candles = candles
	if isSma {
		df.AddSma(smas[0])
		df.AddSma(smas[1])
		df.AddSma(smas[2])
	}
	return &df, err
}

func (api *APIClient) FetchCwCandles(
	periods string,
	beforeAfter string,
	time string,
) (*models.CwCandles, error) {
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
	var cwCandles models.CwCandles
	err = json.Unmarshal(resp, &cwCandles)
	if err != nil {
		return nil, err
	}
	return &cwCandles, nil
}
