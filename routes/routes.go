package routes

import (
	"net/http"

	"project-2-Herwindams24/controllers"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/daftarInstansi", controllers.FetchAllDaftarInstansi)
	e.POST("/daftarInstansi", controllers.StoreDaftarInstansi)
	e.PUT("/daftarInstansi", controllers.UpdateDaftarInstansi)

	return e
}
