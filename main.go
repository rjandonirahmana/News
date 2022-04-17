package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rjandonirahmana/news/auth"
	"github.com/rjandonirahmana/news/database"
	"github.com/rjandonirahmana/news/grpc/user"
	"github.com/rjandonirahmana/news/grpchandler"
	"github.com/rjandonirahmana/news/handler"
	"github.com/rjandonirahmana/news/repository"
	"github.com/rjandonirahmana/news/route"
	"github.com/rjandonirahmana/news/usecase"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

func main() {

	dbmongo, err := database.ConnectionMongo("mongodb://localhost:27017", "News", context.Background())
	log.Println(dbmongo)
	log.Println(err)

	if err != nil {
		log.Fatal(err)
	}

	redis, err := database.RedisConnection("", "")
	if err != nil {
		log.Fatal(err)
	}

	authentication := auth.NewAuth("coba", "test")
	authredis := auth.NewAuthRedis(redis)

	repoNews := repository.NewRepoNews(dbmongo)
	servicenews := usecase.NewServiceNews(repoNews)
	handlerNews := handler.NewHanlderNews(servicenews)
	repoAdmin := repository.NewAdminRepo(dbmongo)

	repoUser := repository.NewRepoUser(dbmongo)
	serviceUser := usecase.NewUsecCaseUser(repoUser, "secret", repoAdmin)
	handlerUser := handler.NewHandlerUser(serviceUser, authentication, authredis)

	repoComment := repository.NewRepoComments(dbmongo)
	serviceComment := usecase.NewUseCaseComent(repoComment)
	handlerComment := handler.NewCommentHandler(serviceComment)

	// serviceAdmin := usecase.NewAdminUsecase(repoAdmin, "SECRETBANGET")
	// handlerAdmin := handler.NewAdminHandler(serviceAdmin, authentication)

	middleWare := handler.NewMiddleWare(authentication, serviceUser, repoAdmin, authredis)
	handlerGrpcUser := grpchandler.NewGrpcUser(serviceUser, authentication)

	r := mux.NewRouter()
	//route News
	route.NewsRoute(r, middleWare, handlerNews)
	//route admin
	// route.RouteAdmin(r, handlerAdmin)
	//route user
	route.RouteUser(r, handlerUser)

	go func() {
		listen, err := net.Listen("tcp", ":11000")
		if err != nil {
			log.Fatalf("[ERROR] Failed to listen tcp: %v", err)
		}

		grpcServer := grpc.NewServer()
		user.RegisterUserServer(grpcServer, handlerGrpcUser)
		grpcServer.Serve(listen)
	}()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"127.0.0.1:7000"},
		AllowedMethods: []string{"POST", "PUT", "GET", "DELETE", "PATCH"},
		AllowedHeaders: []string{"X-Access-Token", "Accept-Language", "Content-Type", "Content-Language", "Origin"},
	})

	r.HandleFunc("/create/comment", middleWare.AuthenticationUser(handlerComment.CreateComment)).Methods("PUT")
	r.HandleFunc("/like/comment", middleWare.AuthenticationUser(handlerComment.LikeComment)).Methods("PUT")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:7000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	handlerCors := c.Handler(r)

	log.Fatal(srv.ListenAndServe(), handlerCors)

}
