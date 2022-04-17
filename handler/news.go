package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

	admin := r.Context().Value("user").(*models.User)

	var news models.News

	photo, _, err := r.FormFile("photo")
	if err != nil {
		resp := APIResponse("failed pass photo", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}

	file, err := ioutil.ReadAll(photo)
	if err != nil {
		resp := APIResponse("failed read file photo", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}
	category := r.FormValue("category")
	categoryID, err := strconv.Atoi(category)
	if err != nil {
		resp := APIResponse("category is not int", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}

	news.Name = r.FormValue("title")
	news.Categroy.ID = categoryID
	news.Content = r.FormValue("content")
	news.Location = r.FormValue("location")
	news.Author = admin.Email
	fmt.Println(admin)

	news1, err := h.service.CreateNews(&news, file, context.Background())
	if err != nil {
		resp := APIResponse("failed read file photo", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(500)
		w.Write([]byte(resbyte))
		return

	}

	resp, _ := json.Marshal(news1)
	w.Write(resp)

}

func (h *NewsHandler) GetNewsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	param := r.URL.Query()
	id := param.Get("news_id")

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

func (h *NewsHandler) DeleteNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	param := r.URL.Query()
	id := param.Get("news_id")

	err := h.service.DeleteNewsByID(&id, context.Background())
	if err != nil {
		resp := APIResponse("canot decode json", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}

	resp := APIResponse("success get news", 200, "success", nil)
	resByte, _ := json.Marshal(resp)
	w.Write(resByte)
}
