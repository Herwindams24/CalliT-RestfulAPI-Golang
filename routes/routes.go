package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"project-2-Herwindams24/controllers"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/daftarInstansi", controllers.FetchAllDaftarInstansi)

	return e
}