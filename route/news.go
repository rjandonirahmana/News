package route

import (
	"github.com/gorilla/mux"
	"github.com/rjandonirahmana/news/handler"
)

func NewsRoute(r *mux.Router, middleware *handler.MiddleWare, handler *handler.NewsHandler) {

	r.HandleFunc("/news", handler.GetNewsByID).Methods("GET")
	r.HandleFunc("/create/news", middleware.AuthenticationAdmin(handler.CreateNews)).Methods("POST")
}
