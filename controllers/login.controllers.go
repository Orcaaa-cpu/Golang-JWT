package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/orcaaa/echo-rest/helper"
	"github.com/orcaaa/echo-rest/models"
)

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func GenerateHashPassword(c echo.Context) error {
	pass := c.Param("password")

	hash, _ := helper.HashPassword(pass)

	return c.JSON(http.StatusOK, hash)
}

func SingUp(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	// ok, err := models.UniqUsername(user.Username)
	// helper.PanicErr(err)
	// if ok {
	// 	return c.JSON(http.StatusInternalServerError, "Username "+user.Username+" Telah Di Gunakan !")
	// }

	hash, _ := helper.HashPassword(user.Password)

	res, err := models.SingUp(user.Username, hash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
