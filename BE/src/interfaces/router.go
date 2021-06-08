package interfaces

import (
	"app/src/config"
	"app/src/interfaces/controllers"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// Run start server
func Run(e *echo.Echo, port string) {
	log.Printf("Server running at http://localhost:%s/", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

// Routes returns the initialized router
func Routes(e *echo.Echo) {

	// Health Check
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Good morning, Golang + Nuxt.js !")
	})

	// api v1
	apiV1 := e.Group("api/v1")
	apiV1.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "api v1 status 200")
	})

	// candle
	t := apiV1.Group("/candles")
	t.GET("", func(c echo.Context) error {
		tickerController := controllers.NewTickerController()
		return tickerController.Past(c)
	})

	// Migration Route
	e.GET("/api/v1/migrate", migrate)
	e.GET("/api/v1/seed", Seeds)
}

// =============================
//    MIGRATE
// =============================
func migrate(c echo.Context) error {
	_, err := config.DBMigrate()
	if err != nil {
		return c.String(http.StatusNotFound, "failed migrate")
	} else {
		return c.String(http.StatusOK, "success migrate")
	}
}

// =============================
//    SEED
// =============================
func Seeds(c echo.Context) error {
	_, err := config.Seeds()
	if err != nil {
		return c.String(http.StatusNotFound, "failed seed")
	} else {
		return c.String(http.StatusOK, "success seed")
	}
}
