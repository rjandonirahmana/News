package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rjandonirahmana/news/database"
	"github.com/rjandonirahmana/news/handler"
	"github.com/rjandonirahmana/news/repository"
	"github.com/rjandonirahmana/news/route"
	"github.com/rjandonirahmana/news/usecase"
)

func main() {
	// clientoption.ApplyURI("mongodb://localhost:27017")

	dbmongo, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())

	if err != nil {
		log.Fatal(err)
	}

	repoNews := repository.NewRepoNews(dbmongo)
	servicenews := usecase.NewServiceNews(repoNews)
	handlerNews := handler.NewHanlderNews(servicenews)

	repoComment := repository.NewRepoComments(dbmongo)
	serviceComment := usecase.NewUseCaseComent(repoComment)
	handlerComment := handler.NewCommentHandler(serviceComment)

	r := mux.NewRouter()
	route.NewsRoute(r, handlerNews)

	r.HandleFunc("/create/comment", handlerComment.CreateComment).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:7000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
