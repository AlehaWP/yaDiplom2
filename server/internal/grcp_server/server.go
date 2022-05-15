package grcp_server

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/models"

	pb "github.com/AlehaWP/yaDiplom2.git/server/internal/grcp_server/proto"

	"google.golang.org/grpc"
)

var repo models.Database

// UsersServer поддерживает все необходимые методы.
type GophePassServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedGophePassServer
}

// AddUser реализует интерфейс добавления пользователя.
func (s *GophePassServer) AddFile(ctx context.Context, in *pb.AddFileRequest) (*pb.AddFileResponse, error) {
	response := &pb.AddFileResponse{
		Error: "нет ошибок",
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

func Start(ctx context.Context) {
	// cfg := config.NewConfig()

	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		fmt.Println(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()
	// регистрируем сервис
	pb.RegisterGophePassServer(s, &GophePassServer{})

	fmt.Println("сервер gRPC начал работу")
	// получаем запрос gRpc
	if err := s.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
