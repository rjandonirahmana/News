package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rjandonirahmana/news/models"
	"github.com/rjandonirahmana/news/usecase"
)

type NewsHandler struct {
	service usecase.UsecaseNews
}

func NewHanlderNews(service usecase.UsecaseNews) *NewsHandler {
	return &NewsHandler{service: service}
}

func (h *NewsHandler) CreateNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	admin := r.Context().Value("admin").(*models.Admin)

	var news models.News

	err := json.NewDecoder(r.Body).Decode(&news)
	if err != nil {
		resp := APIResponse("canot decode json", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}
	news.Author = admin.Email
	fmt.Println(admin)

	news1, err := h.service.CreateNews(&news, context.Background())
	if err != nil {
		w.WriteHeader(500)
		return
	}

	resp, _ := json.Marshal(news1)
	w.Write(resp)

}

func (h *NewsHandler) GetNewsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	param := r.URL.Query()
	id := param.Get("id")

	news, err := h.service.GetNewsByID(&id, context.Background())
	if err != nil {
		resp := APIResponse("canot decode json", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}

	resp := APIResponse("success get news", 200, "success", news)
	resByte, _ := json.Marshal(resp)
	w.Write(resByte)

}
