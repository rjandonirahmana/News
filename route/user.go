package route

import (
	"github.com/gorilla/mux"
	"github.com/rjandonirahmana/news/handler"
)

func RouteUser(r *mux.Router, handler *handler.HanlderUser) {

	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
}
