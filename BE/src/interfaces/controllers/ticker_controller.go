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

// FIXME: sma,emea,BBandsはオブジェクトにしたい
type CandleParams struct {
	Periods     string `json:"Periods"`     // FIXME: json to snake case
	BeforeAfter string `json:"BeforeAfter"` // FIXME: json to snake case
	Time        string `json:"Time"`        // FIXME: json to lowercase
	Smas        string `json:"smas"`
	Sma1        string `json:"sma1"`
	Sma2        string `json:"sma2"`
	Sma3        string `json:"sma3"`
	Emas        string `json:"emas"`
	Ema1        string `json:"ema1"`
	Ema2        string `json:"ema2"`
	Ema3        string `json:"ema3"`
	BBands      string `json:"bbands"`
	BBandsN     string `json:"bbandsN"`
	BBandsK     string `json:"bbandsK"`
	Ichimoku    string `json:"ichimoku"`
	Rsi         string `json:"rsi"`
	RsiPeriod   string `json:"rsiPeriod"`
	Macd        string `json:"macd"`
	Macd1       string `json:"macd1"`
	Macd2       string `json:"macd2"`
	Macd3       string `json:"macd3"`
}

func (t *TickerController) Past(c echo.Context) error {
	tickerUsecase := usecase.NewTickerUsecase()
	var candleParams CandleParams
	if err := c.Bind(&candleParams); err != nil {
		// TODO: error handling
	}

	periods := candleParams.Periods
	beforeAfter := candleParams.BeforeAfter
	unitTimeStamp := candleParams.Time
	smasStr := candleParams.Smas
	emasStr := candleParams.Emas
	bbandsStr := candleParams.BBands
	ichimokuStr := candleParams.Ichimoku
	rsiStr := candleParams.Rsi
	macdStr := candleParams.Macd

	isSmas := false
	smas := make([]int, 3)
	if smasStr != "" {
		isSmas = true
		sma1Str := candleParams.Sma1
		sma2Str := candleParams.Sma2
		sma3Str := candleParams.Sma3
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

	isEmas := false
	emas := make([]int, 3)
	if emasStr != "" {
		isEmas = true
		ema1Str := candleParams.Ema1
		ema2Str := candleParams.Ema2
		ema3Str := candleParams.Ema3
		ema1, err := strconv.Atoi(ema1Str)
		if ema1Str == "" || err != nil || ema1 < 0 {
			ema1 = 7
		}
		ema2, err := strconv.Atoi(ema2Str)
		if ema2Str == "" || err != nil || ema2 < 0 {
			ema2 = 14
		}
		ema3, err := strconv.Atoi(ema3Str)
		if ema3Str == "" || err != nil || ema3 < 0 {
			ema3 = 50
		}
		emas[0] = ema1
		emas[1] = ema2
		emas[2] = ema3
	}

	isBBands := false
	var bbandsN, bbandsK int
	if bbandsStr != "" {
		isBBands = true
		bbandsNStr := candleParams.BBandsN
		bbandsKStr := candleParams.BBandsK
		n, err := strconv.Atoi(bbandsNStr)
		bbandsN = n
		if bbandsNStr == "" || err != nil || bbandsN < 0 {
			bbandsN = 20
		}
		k, err := strconv.Atoi(bbandsNStr)
		bbandsK = k
		if bbandsKStr == "" || err != nil || bbandsK < 0 {
			bbandsK = 2
		}
	}

	isIchimoku := false
	if ichimokuStr != "" {
		isIchimoku = true
	}

	isRsi := false
	var rsiPeriod int
	if rsiStr != "" {
		isRsi = true
		rsiPeriodStr := candleParams.RsiPeriod
		period, err := strconv.Atoi(rsiPeriodStr)
		rsiPeriod = period
		if rsiPeriodStr == "" || err != nil || period < 0 {
			rsiPeriod = 14
		}
	}

	isMacd := false
	macds := make([]int, 3)
	if macdStr != "" {
		isMacd = true
		macd1Str := candleParams.Macd1
		macd2Str := candleParams.Macd2
		macd3Str := candleParams.Macd3
		macd1, err := strconv.Atoi(macd1Str)
		if macd1Str == "" || err != nil || macd1 < 0 {
			macd1 = 12
		}
		macd2, err := strconv.Atoi(macd2Str)
		if macd2Str == "" || err != nil || macd2 < 0 {
			macd2 = 26
		}
		macd3, err := strconv.Atoi(macd3Str)
		if macd3Str == "" || err != nil || macd3 < 0 {
			macd3 = 9
		}
		macds[0] = macd1
		macds[1] = macd2
		macds[2] = macd3
	}
	log.Println("isMacd")
	log.Println(isMacd)
	log.Println("macds")
	log.Println(macds)

	dataFlameCandles, err := tickerUsecase.FetchDataFlameCandles(
		periods,
		beforeAfter,
		unitTimeStamp,
		isSmas,
		smas,
		isEmas,
		emas,
		isBBands,
		bbandsN,
		bbandsK,
		isIchimoku,
		isRsi,
		rsiPeriod,
		isMacd,
		macds,
	)
	log.Println("dataFlameCandles.Macd")
	log.Println(dataFlameCandles.Macd)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, dataFlameCandles)
}
