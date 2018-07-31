package main

import (
	"fmt"
	"net/http"
	"strings"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a := App()
	a.initialize("", "", "")
}