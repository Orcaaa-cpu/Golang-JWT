package tess

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/orcaaa/echo-rest/helper"
)

func CreatCon() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/echo_rest")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func Uniq(usernames string) bool {
	conn := CreatCon()

	script := "SELECT username FROM users"

	rows, err := conn.Query(script)
	helper.PanicErr(err)
	defer rows.Close()

	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		helper.PanicErr(err)
		if username == usernames {
			return false
		}
	}

	return true
}

func TestOk(t *testing.T) {

	fmt.Println(Uniq("arga"))

}
