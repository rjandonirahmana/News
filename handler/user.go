package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rjandonirahmana/news/auth"
	"github.com/rjandonirahmana/news/models"
	"github.com/rjandonirahmana/news/usecase"
)

type HanlderUser struct {
	service   usecase.UsecaseUser
	auth      auth.Auth
	authredis auth.Authentication
}

func NewHandlerUser(service usecase.UsecaseUser, authentication auth.Auth, authredis auth.Authentication) *HanlderUser {
	return &HanlderUser{service: service, auth: authentication, authredis: authredis}
}

func (h *HanlderUser) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp := APIResponse("canot decode json", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write(resbyte)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	user, err = h.service.CreateUser(user, ctx)
	if err != nil {
		resp := APIResponse("failed register user", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write(resbyte)
		return
	}

	token, err := h.authredis.GenerateToken(ctx, user.Email, user.ID, "user")
	if err != nil {
		resp := APIResponse("failed to generate token", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write(resbyte)
		return
	}

	response := APIResponseToken("success create new user", 200, "success", user, token)
	respByte, _ := json.Marshal(response)
	w.WriteHeader(200)
	w.Write(respByte)

}

func (h *HanlderUser) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp := APIResponse("canot decode json", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write(resbyte)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	
	user, roles, err := h.service.LoginUser(&user.Email, &user.Password, ctx)
	if err != nil {
		resp := APIResponse("failed to login", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write(resbyte)
		return
	}

	token, err := h.authredis.GenerateToken(ctx, user.Email, user.ID, roles)
	if err != nil {
		resp := APIResponse("failed to generate token", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write(resbyte)
		return
	}

	response := APIResponseToken("success login", 200, "success", user, token)
	respByte, _ := json.Marshal(response)
	w.WriteHeader(200)
	w.Write(respByte)
}
