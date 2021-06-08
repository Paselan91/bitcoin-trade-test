package controllers

import (
	"app/src/usecase"
	// "fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	// "reflect"
	"strconv"
	// "encoding/json"
)

type TickerController struct{}

func NewTickerController() *TickerController {
	return &TickerController{}
}

// FIXME: smaはオブジェクトにしたい

type CandleParams struct {
	Periods     string `json:"Periods"`
	BeforeAfter string `json:"BeforeAfter"`
	Time        string `json:"Time"`
	Smas        string `json:"smas"`
	Sma1        string `json:"sma1"`
	Sma2        string `json:"sma2"`
	Sma3        string `json:"sma3"`
}

// type Smas struct {
// 	Sma1 int `json:"sma1"`
// 	Sma2 int `json:"sma2"`
// 	Sma3 int `json:"sma3"`
// }

func (t *TickerController) Past(c echo.Context) error {
	tickerUsecase := usecase.NewTickerUsecase()
	var candleParams CandleParams
	if err := c.Bind(&candleParams); err != nil {
		// TODO: error handling
	}

	log.Println("candleParams")
	log.Println(candleParams)
	log.Println("---------- 0 -----------")
	log.Println(candleParams.Smas)
	log.Println("---------- 1 -----------")

	log.Println(candleParams.Smas == "")
	// log.Println(candleParams.Smas != "")
	log.Println("---------- 2 -----------")

	log.Println(candleParams.Sma1)
	log.Println("---------- xx -----------")

	log.Println(candleParams.Sma2)
	log.Println(candleParams.Sma3)
	log.Println("---------- 3 -----------")

	// name := c.QueryParam("smas")
	// log.Println("")
	// log.Println("name type")
	// log.Println(reflect.TypeOf(name))
	// log.Println("")

	periods := candleParams.Periods
	beforeAfter := candleParams.BeforeAfter
	unitTimeStamp := candleParams.Time
	SmasStr := candleParams.Smas
	log.Println("---------- 4 -----------")

	var isSmas bool
	// var sma1 int
	// var sma2 int
	// var sma3 int
	// var smas []int
	smas := make([]int, 3)
	if SmasStr != "" {
		log.Println("---------- hi -----------")
		sma1Str := candleParams.Sma1
		sma2Str := candleParams.Sma2
		sma3Str := candleParams.Sma3
		// sma1 = candleParams.Sma1
		// sma2 = candleParams.Sma2
		// sma3 = candleParams.Sma3
		isSmas = true
		sma1, err := strconv.Atoi(sma1Str)
		if sma1Str == "" || err != nil || sma1 < 0 {
			sma1 = 7
		}
		sma2, err := strconv.Atoi(sma2Str)
		if sma2Str == "" || err != nil || sma2 < 0 {
			sma2 = 14
		}
		sma3, err := strconv.Atoi(sma3Str)
		if sma3Str == "" || err != nil || sma3 < 0 {
			sma3 = 50
		}
		smas[0] = sma1
		smas[1] = sma2
		smas[2] = sma3
	}
	log.Println("isSmas")
	log.Println(isSmas)
	log.Println("smas")
	log.Println(smas)

	dataFlameCandles, err := tickerUsecase.FetchDataFlameCandles(
		periods,
		beforeAfter,
		unitTimeStamp,
		isSmas,
		smas,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, dataFlameCandles)
}
