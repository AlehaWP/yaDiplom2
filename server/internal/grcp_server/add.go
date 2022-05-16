package grcp_server

import (
	"context"
	"fmt"
	"os"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/database"
	pb "github.com/AlehaWP/yaDiplom2.git/server/internal/grcp_server/proto"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/logger"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/models"
)

// AddUser реализует интерфейс добавления пользователя.
func (s *GophePassServer) AddFile(ctx context.Context, in *pb.AddFileRequest) (*pb.AddResponse, error) {
	response := &pb.AddResponse{
		Message: "нет ошибок",
	}

	user := in.User

	file := in.File.Data

	os.WriteFile(in.File.Name, file, 0777)
	fmt.Println(user)

	// userID := in.User
	// retURL, err := repo.SaveURL(ctx, in.Url.Url, baseURL, userID)
	// if err != nil {
	// 	return nil, err
	// }
	// response.Url.Url = retURL

	return response, nil

}

func (s *GophePassServer) AddAcc(ctx context.Context, in *pb.AddAccRequest) (*pb.AddResponse, error) {

	response := &pb.AddResponse{
		Message: "нет ошибок",
	}

	acc := models.Account{
		User: models.User{
			UUID: in.User.Uuid,
		},
		Login:    in.Account.Login,
		Password: in.Account.Password,
		UUID:     in.Account.Uuid,
	}

	if err := database.AddAccount(ctx, acc); err != nil {
		logger.Info("AddAcc", "Ошибка добавления данных", acc, err)
		response.Message = "Ошибка добавления данных"
		return response, err
	}

	return response, nil

}

func (s *GophePassServer) AddCard(ctx context.Context, in *pb.AddCardRequest) (*pb.AddResponse, error) {
	response := &pb.AddResponse{
		Message: "нет ошибок",
	}

	// user := in.User

	// file := in.File.Data

	// os.WriteFile(in.File.Name, file, 0777)
	// fmt.Println(user)

	// userID := in.User
	// retURL, err := repo.SaveURL(ctx, in.Url.Url, baseURL, userID)
	// if err != nil {
	// 	return nil, err
	// }
	// response.Url.Url = retURL

	return response, nil

}
