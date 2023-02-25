package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/orcaaa/echo-rest/controllers"
	"github.com/orcaaa/echo-rest/middlewares"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai, middlewares.IsAuthenticated)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai", controllers.UpdatePegawai)
	e.DELETE("/pegawai", controllers.DeletePegawai)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	e.POST("/singUp", controllers.SingUp)

	return e

}
