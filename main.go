package main

import (
	"github.com/orcaaa/echo-rest/db"
	"github.com/orcaaa/echo-rest/routes"
)

func main() {

	db.CreateCon()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
