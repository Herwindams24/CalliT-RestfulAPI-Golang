package controllers

import (
	"net/http"
	"project-2-Herwindams24/models"
	"github.com/labstack/echo"

)


func FetchAllDaftarInstansi(c echo.Context) error {
	result, err := models.FetchAllDaftarInstansi()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}