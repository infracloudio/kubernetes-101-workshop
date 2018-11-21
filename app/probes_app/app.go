package main

import (
	"fmt"
	"log"
	"net/http"
)

var ready = 0
var health = 0

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("Path %s called ", r.RequestURI))
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("Readiness probe called  %d time", ready))
	ready++
	if ready < 5 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintf(w, "Ready")
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	/*health++
	log.Println(fmt.Sprintf("Health probe called  %d time", health))
	if health > 40 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/

	fmt.Fprintf(w, "Ready")
	return
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ready", readyHandler)
	http.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
