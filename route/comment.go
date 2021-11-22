package route

import (
	"github.com/gorilla/mux"
	"github.com/rjandonirahmana/news/handler"
)

func CommentRoutes(r *mux.Router, middleware *handler.MiddleWare, handler *handler.CommentHanlder) {

	r.HandleFunc("/create/comment", middleware.AuthenticationUser(handler.CreateComment)).Methods("PUT")
	r.HandleFunc("/like/comment", middleware.AuthenticationUser(handler.LikeComment)).Methods("PUT")
}
