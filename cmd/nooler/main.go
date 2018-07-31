package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
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

func userHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Path[6:]
	
	idx := strings.Index(param, "/")

	if idx == -1 {
		fmt.Fprintf(w, "/user/:id id=%s", param)
		return
	}

	static := param[idx+1:]

	if len(static) == 0 {
		http.Redirect(w, r, r.URL.Path[:len(r.URL.Path)-1], http.StatusMovedPermanently)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func main() {
	a := App()
	a.initialize("", "", "")
}