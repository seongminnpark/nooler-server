package main

import (
	"fmt"
	// "html"
	"log"
	"net/http"
	"path"
	"strings"
)

func ShiftPath(p string) (head, tail string) {
	fmt.Printf(p);
	p = path.Clean("/" + p)
	fmt.Printf(p);
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

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
	mux := http.NewServeMux()
	mux.HandleFunc("/user", userHandler)

	log.Fatal(http.ListenAndServe(":2441", mux))
}