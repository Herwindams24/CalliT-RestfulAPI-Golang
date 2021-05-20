package controllers

import (
	"net/http"
	"project-2-Herwindams24/models"
	"strconv"

	"github.com/labstack/echo"
)

func FetchAllDaftarInstansi(c echo.Context) error {
	result, err := models.FetchAllDaftarInstansi()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreDaftarInstansi(c echo.Context) error {
	namaInstansi := c.FormValue("namaInstansi")
	Kota := c.FormValue("Kota")
	telfon := c.FormValue("telfon")

	result, err := models.StoreDaftarInstansi(namaInstansi, Kota, telfon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, result)

}

func UpdateDaftarInstansi(c echo.Context) error{
	id := c.FormValue("id")
	namaInstansi := c.FormValue("namaInstansi")
	Kota := c.FormValue("Kota")
	telfon := c.FormValue("telfon")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateDaftarInstansi(conv_id, namaInstansi, Kota, telfon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
