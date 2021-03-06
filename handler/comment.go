package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rjandonirahmana/news/models"
	"github.com/rjandonirahmana/news/usecase"
)

type CommentHanlder struct {
	usecase usecase.UsecaseComment
}

func NewCommentHandler(service usecase.UsecaseComment) *CommentHanlder {
	return &CommentHanlder{usecase: service}
}

func (h *CommentHanlder) CreateComment(w http.ResponseWriter, r *http.Request) {
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

func (h *CommentHanlder) LikeComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	querry := r.URL.Query()
	newsID := querry.Get("news_id")
	commentID := querry.Get("comment_id")

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
	defer cancel()

	err := h.usecase.LikeComment(&newsID, &commentID, ctx)
	if err != nil {
		resp := APIResponse("failed to likes comment", 422, "error", err.Error())
		resbyte, _ := json.Marshal(resp)
		w.WriteHeader(422)
		w.Write([]byte(resbyte))
		return
	}

	resp := APIResponse("success like comment", 200, "success", nil)
	resByte, _ := json.Marshal(resp)
	w.Write(resByte)

}
