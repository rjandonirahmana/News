package main

import (
	"context"
	"fmt"
	"log"

	// "net/http"
	// "time"

	// "github.com/gorilla/mux"
	// "github.com/rjandonirahmana/news/auth"
	"github.com/rjandonirahmana/news/database"
	// "github.com/rjandonirahmana/news/handler"
	"github.com/rjandonirahmana/news/repository"
	// "github.com/rjandonirahmana/news/route"
	// "github.com/rjandonirahmana/news/usecase"
	// "github.com/rs/cors"
)

func main() {

	dbmongo, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())

	if err != nil {
		log.Fatal(err)
	}

	// authentication := auth.NewAuth("coba", "test")

	// repoNews := repository.NewRepoNews(dbmongo)
	// servicenews := usecase.NewServiceNews(repoNews)
	// handlerNews := handler.NewHanlderNews(servicenews)

	// repoUser := repository.NewRepoUser(dbmongo)
	// serviceUser := usecase.NewUsecCaseUser(repoUser, "secret")
	// handlerUser := handler.NewHandlerUser(serviceUser, authentication)

	empat := "5"
	satu := "1"
	repoComment := repository.NewRepoComments(dbmongo)
	comment, err := repoComment.GetOneComment(&empat, &satu, context.Background())

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(comment)
	// serviceComment := usecase.NewUseCaseComent(repoComment)
	// handlerComment := handler.NewCommentHandler(serviceComment)

	// repoAdmin := repository.NewAdminRepo(dbmongo)
	// serviceAdmin := usecase.NewAdminUsecase(repoAdmin, "SECRETBANGET")
	// handlerAdmin := handler.NewAdminHandler(serviceAdmin, authentication)

	// middleWare := handler.NewMiddleWare(authentication, serviceUser, repoAdmin)

	// r := mux.NewRouter()
	// route.NewsRoute(r, middleWare, handlerNews)
	// route.RouteAdmin(r, handlerAdmin)

	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"127.0.0.1:7000"},
	// 	AllowedMethods: []string{"POST", "PUT", "GET", "DELETE", "PATCH"},
	// 	AllowedHeaders: []string{"X-Access-Token", "Accept-Language", "Content-Type", "Content-Language", "Origin"},
	// })

	// r.HandleFunc("/create/comment", middleWare.AuthenticationUser(handlerComment.CreateComment)).Methods("PUT")
	// r.HandleFunc("/register", handlerUser.Register).Methods("POST")
	// r.HandleFunc("/login", handlerUser.Login).Methods("POST")

	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "127.0.0.1:7000",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 5 * time.Second,
	// 	ReadTimeout:  5 * time.Second,
	// }

	// handlerCors := c.Handler(r)

	// log.Fatal(srv.ListenAndServe(), handlerCors)

}
