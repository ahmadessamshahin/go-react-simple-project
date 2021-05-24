package src

import (
	"fmt"
	"log"
	"net/http"
	c "project/src/controller"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

// App export
type App struct {
	Router *mux.Router
}

func (app *App) InitializeRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/phone-numbers", c.GetPhoneNumbers).Methods("GET")
	app.Router.HandleFunc("/phone-number", c.CreatePhoneNumber).Methods("POST")
	app.Router.HandleFunc("/phone-number/{id:[0-9]+}", c.GetPhoneNumber).Methods("GET")
	// app.Router.HandleFunc("/phone-number/{id:[0-9]+}", c.UpdatePhoneNumber).Methods("PUT")
	app.Router.HandleFunc("/phone-number/{id:[0-9]+}", c.DeletePhoneNumber).Methods("DELETE")
}

func (app *App) Run(port string) {
	fmt.Printf("listening to port %s:\n", port)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}
