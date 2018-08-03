package nooler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/seongminnpark/nooler-server/internal/app/nooler/handler"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(user, password, dbName string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbName)
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) initializeRoutes() {
	userHandler := handler.UserHandler{DB: app.DB}
	app.Router.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
	app.Router.HandleFunc("/user", userHandler.GetUser).Methods("GET")
	app.Router.HandleFunc("/user", userHandler.UpdateUser).Methods("PUT")
	app.Router.HandleFunc("/user", userHandler.DeleteUser).Methods("DELETE")
	app.Router.HandleFunc("/login", userHandler.Login).Methods("POST")

	deviceHandler := handler.DeviceHandler{DB: app.DB}
	app.Router.HandleFunc("/device", deviceHandler.CreateDevice).Methods("POST")
}
