package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/orcaaa/echo-rest/models"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	pegawai := models.Pegawai{}
	c.Bind(&pegawai)

	result, err := models.StorePegawai(pegawai.Nama, pegawai.Alamat, pegawai.Telpon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePegawai(c echo.Context) error {

	pegawai := models.Pegawai{}
	c.Bind(&pegawai)

	result, err := models.UpdatePegawai(pegawai.Id, pegawai.Nama, pegawai.Alamat, pegawai.Telpon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func DeletePegawai(c echo.Context) error {
	id := c.FormValue("id")

	fmt.Println(id)

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeletePegawai(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
