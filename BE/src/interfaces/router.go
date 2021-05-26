package interfaces

import (
	"app/src/config"
	"app/src/interfaces/controllers"
	// "app/src/usecase"
	"fmt"
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Good morning, Golang + Nuxt.js !")
	})

	t := e.Group("/ticker")
	t.GET("/past", func(c echo.Context) error {
		// TODO: Groupの中に入れたい　何回も呼び出したくない
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
