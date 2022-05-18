package grcp_server

import (
	"context"
	"encoding/json"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/database"
	pb "github.com/AlehaWP/yaDiplom2.git/server/internal/grcp_server/proto"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/logger"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/models"
)

func (s *GophePassServer) GetFileList(ctx context.Context, in *pb.GetDataRequest) (*pb.GetDataResponse, error) {

	user := models.User{
		UUID: in.User.Uuid,
	}

	l, err := database.GetListFiles(ctx, user)
	if err != nil {
		logger.Info("Ошибка запроса списка файлов пользователя", user.UUID)
		return nil, err
	}

	b, err := json.Marshal(l)
	if err != nil {
		logger.Info("Ошибка маршализации списка файлов пользователя", user.UUID)
		return nil, err
	}
	resp := &pb.GetDataResponse{
		Message: "Список сформирован",
		Data:    b,
	}
	return resp, nil
}

func (s *GophePassServer) GetFile(ctx context.Context, in *pb.GetDataRequest) (*pb.GetDataResponse, error) {

	return nil, nil
}

func (s *GophePassServer) GetCard(ctx context.Context, in *pb.GetDataRequest) (*pb.GetDataResponse, error) {

	return nil, nil
}

func (s *GophePassServer) GetAcc(ctx context.Context, in *pb.GetDataRequest) (*pb.GetDataResponse, error) {

	return nil, nil
}
