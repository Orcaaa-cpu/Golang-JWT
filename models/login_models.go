package models

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/orcaaa/echo-rest/db"
	"github.com/orcaaa/echo-rest/helper"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlSTMT := "SELECT * from users where username = ?"

	err := con.QueryRow(sqlSTMT, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("querry error")
		return false, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("hash and password doesn't match.")
		return false, err
	}

	return true, nil
}

func SingUp(username, password string) (Response, error) {
	var res Response

	v := validator.New()

	user := User{
		Username: username,
		Password: password,
	}

	err := v.Struct(user)
	if err != nil {
		return res, err
	}

	db := db.CreateCon()

	script := "insert into users(username, password) values(?, ?)"

	stmt, err := db.Prepare(script)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, password)
	if err != nil {
		return res, err
	}

	rows, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUCCESS"
	res.Data = map[string]int64{
		"SUCCESS SING UP ACCOUNT": rows,
	}

	return res, nil
}

func UniqUsername(usernames string) (bool, error) {
	var obj User

	con := db.CreateCon()

	script := "SELECT * from users where username = ? "

	err := con.QueryRow(script, usernames).Scan(&obj.Id, &obj.Username, &obj.Password)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if len(obj.Username) != 0 {
		return true, err
	}

	return false, nil
}
