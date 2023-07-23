package handlers

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func WsEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo!")
}
