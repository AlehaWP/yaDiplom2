package grcp_server

import (
	"context"
	"errors"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/database"
	pb "github.com/AlehaWP/yaDiplom2.git/server/internal/grcp_server/proto"
)

func (s *GophePassServer) CheckConnection(ctx context.Context, in *pb.CheckRequest) (*pb.CheckResponse, error) {
	resp := &pb.CheckResponse{}
	if ok := database.CheckDBConnection(ctx); !ok {
		resp.Message = "Ошибка соединения с базой данных"
		resp.Ok = false
		return nil, errors.New("ошибка соединения с базой данных на сервере")
	}

	resp.Message = "Ок"
	resp.Ok = true
	return resp, nil
}
