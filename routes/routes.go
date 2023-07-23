package routes

import (
	"net/http"

	handlers "github.com/pashamakhilkumarreddy/golang-ci-cd/handlers"
)

func SetUpRoutes() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/ws", handlers.WsEndPoint)
}
