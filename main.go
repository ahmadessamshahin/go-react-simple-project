// main.go
package main

import (
	"os"
	p "project/src"
	db "project/src/db"
)

func main() {
	db.GetDBInstance()
	app := p.App{}
	app.InitializeRoutes()
	app.Run(
		os.Getenv("APP_Port"),
	)
}
