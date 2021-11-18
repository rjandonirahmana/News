package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rjandonirahmana/news/models"
	"github.com/rjandonirahmana/news/usecase"
)

type commentHanlder struct {
	usecase usecase.UsecaseComment
}

func NewCommentHandler(service usecase.UsecaseComment) *commentHanlder {
	return &commentHanlder{usecase: service}
}

func (h *commentHanlder) CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	user := r.Context().Value("user").(*models.User)
	var comment models.CommentNews

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		resp := APIResponse("canot decode json", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}

	comment.UserID = user.ID
	comment.UserName = user.UserName

	newsID := r.URL.Query().Get("news_id")
	err = h.usecase.CreateComment(&newsID, &comment, context.TODO())
	if err != nil {
		resp := APIResponse("canot decode json", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}

	resp := APIResponse("success create news", 200, "success", nil)
	resByte, _ := json.Marshal(resp)
	w.Write(resByte)

}
