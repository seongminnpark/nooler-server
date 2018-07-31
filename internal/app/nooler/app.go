package main

import (
	"fmt"
	"net/http"
	"strings"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) { }
func (a *App) Run(addr string) { }
