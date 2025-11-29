package routing

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

const URL_PREFIX = "/api/v0"

func StartServer() {

	addURL("/", HelloWorld)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addURL(url string, callback func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(URL_PREFIX+url, func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprintf("%s Request at %s", r.Method, URL_PREFIX+url))
		callback(w, r)
	})
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
