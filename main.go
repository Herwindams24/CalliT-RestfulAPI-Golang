package main

import (
	"project-2-Herwindams24/db"

	"project-2-Herwindams24/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}

