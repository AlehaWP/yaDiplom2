package grcp_server

import (
	"context"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/database"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/logger"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/models"
	"github.com/google/uuid"

	pb "github.com/AlehaWP/yaDiplom2.git/server/internal/grcp_server/proto"
)

func (s *GophePassServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp := new(pb.LoginResponse)
	resp.User = new(pb.User)

	user, err := database.GetUser(ctx, in.User.Login)
	if err != nil {
		resp.Finded = false
		logger.Info("Ошибка запроса пользователя", in.User.Login)
		return resp, err
	}

	if user == nil {
		resp.User = in.User
		resp.User.Uuid = uuid.New().String()
		resp.Finded = false
		resp.Auth = false
		return resp, nil
	}

	if user.Password != in.User.Password {
		resp.User = in.User
		resp.User.Uuid = uuid.New().String()
		resp.Finded = true
		resp.Auth = false
		return resp, nil
	}
	resp.Finded = true
	resp.Auth = true
	resp.User.Login = user.Login
	resp.User.Password = user.Password
	resp.User.Uuid = user.UUID
	return resp, nil
}

func (s *GophePassServer) RegUser(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp := new(pb.LoginResponse)
	resp.User = new(pb.User)

	user := models.User{}

	user.Email = in.User.Email
	user.Phone = in.User.Phone
	user.Login = in.User.Login
	user.Password = in.User.Password
	user.UUID = in.User.Uuid

	if user.UUID == "" {
		user.UUID = uuid.New().String()
	}

	err := database.AddUser(ctx, user)
	if err != nil {
		resp.Auth = false
		logger.Info("Ошибка запроса пользователя", in.User.Login)
		return nil, err
	}
	resp.User.Uuid = user.UUID
	return resp, nil
}
