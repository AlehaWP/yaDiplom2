package grcp_server

import (
	"context"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/database"
	pb "github.com/AlehaWP/yaDiplom2.git/server/internal/grcp_server/proto"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/logger"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/models"
)

func (s *GophePassServer) AddFile(ctx context.Context, in *pb.AddFileRequest) (*pb.AddResponse, error) {
	logger.Info("Получен запрос на добавление файла", in.File.Uuid)
	response := &pb.AddResponse{}
	f := models.File{
		User: models.User{
			UUID: in.User.Uuid,
		},
		Name: in.File.Name,
		Data: in.File.Data,
		UUID: in.File.Uuid,
	}
	m := "Добавления данных файла UUID= " + in.File.Uuid + " на сервер "
	if err := database.AddFile(ctx, f); err != nil {
		logger.Info("AddFile", m, f, err)
		response.Message = m + " вызвало ошибку"
		return response, err
	}
	response.Message = m + " прошло успешно"
	return response, nil
}

func (s *GophePassServer) AddAcc(ctx context.Context, in *pb.AddAccRequest) (*pb.AddResponse, error) {
	logger.Info("Получен запрос на добавление аккаунта", in.Account.Uuid)
	response := &pb.AddResponse{}

	a := models.Account{
		User: models.User{
			UUID: in.User.Uuid,
		},
		Login:    in.Account.Login,
		Password: in.Account.Password,
		UUID:     in.Account.Uuid,
	}
	m := "Добавление данных аккаунта UUID= " + in.Account.Uuid + " на сервер"
	if err := database.AddAccount(ctx, a); err != nil {
		logger.Info("AddAcc", m, a, err)
		response.Message = m + " вызвало ошибку"
		return response, err
	}
	response.Message = m + " прошло успешно"
	return response, nil

}

func (s *GophePassServer) AddCard(ctx context.Context, in *pb.AddCardRequest) (*pb.AddResponse, error) {
	logger.Info("Получен запрос на добавление карты", in.Card.Uuid)
	response := &pb.AddResponse{}

	c := models.Card{
		User: models.User{
			UUID: in.User.Uuid,
		},
		Number: in.Card.Number,
		Month:  int(in.Card.Month),
		Year:   int(in.Card.Month),
		Owner:  in.Card.Owner,
		UUID:   in.Card.Uuid,
	}
	m := "Добавления данных карты UUID= " + in.Card.Uuid + " на сервер "
	if err := database.AddCard(ctx, c); err != nil {
		logger.Info("AddCard", m, c, err)
		response.Message = m + " вызвало ошибку"
		return response, err
	}
	response.Message = m + " прошло успешно"
	return response, nil
}
