package grcp_server

import (
	"context"
	"fmt"
	"net"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/config"
	pb "github.com/AlehaWP/yaDiplom2.git/server/internal/grcp_server/proto"

	"google.golang.org/grpc"
)

// UsersServer поддерживает все необходимые методы.
type GophePassServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedGophePassServer
}

func Start(ctx context.Context) {
	cfg := config.NewConfig()

	// определяем порт для сервера
	listen, err := net.Listen("tcp", cfg.ServAddr)
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
