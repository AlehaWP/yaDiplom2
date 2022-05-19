package grcp_server

import (
	"context"
	"errors"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/database"
	pb "github.com/AlehaWP/yaDiplom2.git/server/internal/grcp_server/proto"
)

func (s *GophePassServer) ChekConnection(ctx context.Context) (*pb.GetDataResponse, error) {
	resp := &pb.GetDataResponse{}
	if ok := database.CheckDBConnection(ctx); !ok {
		resp.Message = "Ошибка соединения с базой данных"
		return nil, errors.New("ошибка соединения с базой данных на сервере")
	}

	resp.Message = "Ок"
	return resp, nil
}
