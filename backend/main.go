// main.go
package main

import (
	p "backend/src"
	db "backend/src/db"
	"os"
)

func main() {
	db.GetDBInstance()
	app := p.App{}
	app.InitializeRoutes()
	app.Run(
		os.Getenv("APP_Port"),
	)
}
