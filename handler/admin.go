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

type AdminHandler struct {
	service usecase.AdminUseCase
	auth    auth.Auth
}

func NewAdminHandler(service usecase.AdminUseCase, auth auth.Auth) *AdminHandler {
	return &AdminHandler{service: service, auth: auth}
}

func (h *AdminHandler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	var admin *models.Admin

	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		resp := APIResponse("canot decode json to struct admin", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}

	token, err := h.auth.TokenAdmin(&admin.ID)
	if err != nil {
		resp := APIResponse("failed to generate token", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write(resbyte)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	admin, err = h.service.CreateAdmin(admin, ctx)
	if err != nil {
		resp := APIResponse("failed to create admin", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}

	response := APIResponseToken("success create new admin", 200, "success", admin, *token)
	respByte, _ := json.Marshal(response)
	w.WriteHeader(200)
	w.Write(respByte)

}
