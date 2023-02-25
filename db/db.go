package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// var db *sql.DB
// var err error

// func Init() {
// 	conf := config.GetConfig()

// 	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

// 	db, err = sql.Open("mysql", connectionString)
// 	if err != nil {
// 		panic("ConnectionString Error")
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		panic("DNS Invalid")
// 	}
// }

// func CreateCon() *sql.DB {
// 	return db
// }

func CreateCon() *sql.DB {

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
