package routes

import (
	"github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"

	"net/http"

	middleware "project-2-Herwindams24/middlewares"

	"project-2-Herwindams24/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/daftarInstansi", controllers.FetchAllDaftarInstansi, middleware.IsAuth)
	e.POST("/daftarInstansi", controllers.StoreDaftarInstansi, middleware.IsAuth)
	e.PUT("/daftarInstansi", controllers.UpdateDaftarInstansi, middleware.IsAuth)
	e.DELETE("/daftarInstansi", controllers.DeleteDaftarInstansi, middleware.IsAuth)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	e.GET("/test-struct-validation", controllers.TestStructValidation)
	e.GET("/test-variable-validation", controllers.TestVariableValidation)

	e.GET("/documentation/*", echoSwagger.WrapHandler)

	return e
}
