package main

import (
	"github.com/Herwindams24/project-2-Herwindams24/db"

	"github.com/Herwindams24/project-2-Herwindams24/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
