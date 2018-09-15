package main

import (
	"os"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	return ":" + port
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	name, _ := os.Hostname()
	w.Write([]byte(name))
}
