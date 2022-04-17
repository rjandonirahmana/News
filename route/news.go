package route

import (
	"github.com/gorilla/mux"
	"github.com/rjandonirahmana/news/handler"
)

func NewsRoute(r *mux.Router, middleware *handler.MiddleWare, handler *handler.NewsHandler) {

	//querry news_id
	r.HandleFunc("/news", handler.GetNewsByID).Methods("GET")

	r.HandleFunc("/news", middleware.AuthenticationNews(handler.CreateNews)).Methods("POST")
	//querry news_id
	r.HandleFunc("/news", middleware.AuthenticationNews(handler.DeleteNews)).Methods("DELETE")

}
