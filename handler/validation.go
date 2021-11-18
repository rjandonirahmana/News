package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rjandonirahmana/news/auth"
	"github.com/rjandonirahmana/news/repository"
	"github.com/rjandonirahmana/news/usecase"
)

type MiddleWare struct {
	auth      auth.Auth
	service   usecase.UsecaseUser
	adminRepo repository.AdminRepository
}

func NewMiddleWare(auth auth.Auth, service usecase.UsecaseUser, repoAdmin repository.AdminRepository) *MiddleWare {
	return &MiddleWare{auth: auth, service: service, adminRepo: repoAdmin}
}

func (m *MiddleWare) AuthenticationUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "aplication/json")

		accessToken := r.Header["X-Access-Token"]

		id, err := m.auth.ValidateToken(&accessToken[0])
		if err != nil {
			resp := APIResponse("failed to get validate id", 422, "error", err.Error())
			resbyte, _ := json.Marshal(resp)
			w.WriteHeader(422)
			w.Write(resbyte)
			return

		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		user, err := m.service.GetUserByID(id, ctx)
		if err != nil {
			resp := APIResponse("failed to get user", 422, "error", err.Error())
			resbyte, _ := json.Marshal(resp)
			w.WriteHeader(422)
			w.Write(resbyte)
			return

		}

		ctxValue := context.WithValue(ctx, "user", user)

		r = r.WithContext(ctxValue)
		next.ServeHTTP(w, r)

	}
}

func (m *MiddleWare) AuthenticationAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "aplication/json")

		accessToken := r.Header["X-Access-Token"]

		id, err := m.auth.ValidateAdmin(&accessToken[0])
		if err != nil {
			resp := APIResponse("failed to get validate id admin, or you are not admin", 422, "error", err.Error())
			resbyte, _ := json.Marshal(resp)
			w.WriteHeader(422)
			w.Write(resbyte)
			return

		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		admin, err := m.adminRepo.GetAdminByID(id, ctx)
		if err != nil {
			resp := APIResponse("failed to get user", 422, "error", err.Error())
			resbyte, _ := json.Marshal(resp)
			w.WriteHeader(422)
			w.Write(resbyte)
			return

		}

		ctxValue := context.WithValue(ctx, "admin", admin)
		fmt.Println(admin)

		r = r.WithContext(ctxValue)
		next.ServeHTTP(w, r)

	}
}
