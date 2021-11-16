package route

import (
	"github.com/gorilla/mux"
	"github.com/rjandonirahmana/news/handler"
)

func NewsRoute(r *mux.Router, handler *handler.NewsHandler) {

	r.HandleFunc("/news", handler.GetNewsByID).Methods("GET")
	r.HandleFunc("/create", handler.CreateNews).Methods("POST")
}
