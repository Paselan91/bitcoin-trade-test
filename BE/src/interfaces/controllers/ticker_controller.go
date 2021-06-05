package controllers

import (
	"app/src/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type TickerController struct{}

func NewTickerController() *TickerController {
	return &TickerController{}
}

type OhlcParams struct {
	Periods     string `json:"Periods"`
	BeforeAfter string `json:"BeforeAfter"`
	Time        string `json:"Time"`
}

func (t *TickerController) Past(c echo.Context) error {
	tickerUsecase := usecase.NewTickerUsecase()
	var ohlcParams OhlcParams
	if err := c.Bind(&ohlcParams); err != nil {
		// TODO: error handling
	}

	periods := ohlcParams.Periods
	beforeAfter := ohlcParams.BeforeAfter
	unitTimeStamp := ohlcParams.Time

	tickers, _ := tickerUsecase.FetchDataFlemeCandles(periods, beforeAfter, unitTimeStamp)
	return c.JSON(http.StatusOK, tickers)
}
