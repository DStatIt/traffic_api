package main

import (
	"log"
	"net/http"

	traffic_api "github.com/DStatIt/traffic_api"
	"github.com/gorilla/mux"
)

type appHandler func(w http.ResponseWriter, r *http.Request) (int, error)

func main() {
	r := mux.NewRouter()
	r.Handle("/script.js", appHandler(traffic_api.GetScript))
	// r.Handle("/data/", appHandler(traffic_api.Data))
	// r.Handle("/read/", appHandler(traffic_api.Read))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code, err := fn(w, r)
	if err != nil {
		// http.Error(w, http.StatusText(500), 500)
		log.Println(err)
		return
	}

	switch code {
	case 200:
		return
	case 500:
		http.Error(w, http.StatusText(code), code)
	}

	return
}
