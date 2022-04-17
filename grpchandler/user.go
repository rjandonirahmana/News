package grpchandler

import (
	"context"
	"errors"

	"github.com/rjandonirahmana/news/auth"
	"github.com/rjandonirahmana/news/grpc/user"
	"github.com/rjandonirahmana/news/models"
	"github.com/rjandonirahmana/news/usecase"
)

type GrpcUser struct {
	service usecase.UsecaseUser
	auth    auth.Auth
	user.UnimplementedUserServer
}

func NewGrpcUser(usecase usecase.UsecaseUser, auth auth.Auth) *GrpcUser {
	return &GrpcUser{service: usecase, auth: auth}
}

func (g *GrpcUser) RegisterUser(ctx context.Context, req *user.Register) (*user.RegisterResponse, error) {

	if req.Password != req.Confirmpassword {
		return nil, errors.New("password and confirm password doenst match")
	}
	user2, err := g.service.CreateUser(&models.User{
		UserName: req.Name,
		Email:    req.Email,
		Password: req.Password,
	}, ctx)

	if err != nil {
		return nil, err
	}

	token, err := g.auth.CreateToken(&user2.Email)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{
		Userid: user2.ID,
		Name:   user2.Email,
		Email:  user2.Email,
		Token:  *token,
	}, nil
}
