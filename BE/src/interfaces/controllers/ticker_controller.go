package controllers

import (
    // "app/src/usecase"
    "github.com/labstack/echo"
	"net/http"
)

type TickerController struct {}

func NewTickerController() *TickerController {
	return &TickerController{}
}

func (t *TickerController) Past(c echo.Context) error {
    // userUid  := c.Request().Header.Get("current-user-id")
    // userDisplayName := c.QueryParam("userDisplayName")
    // oauthType := c.QueryParam("oauthType")
	// df, _ := usecase.GetAllCandle(productCode, durationTime, limit)
	return c.JSON(http.StatusOK, "ok heyheyhey")
}
