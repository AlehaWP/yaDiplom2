package grcp_server

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/config"
	pb "github.com/AlehaWP/yaDiplom2.git/server/internal/grcp_server/proto"
	"golang.org/x/crypto/acme/autocert"

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

	manager := &autocert.Manager{
		Cache:  autocert.DirCache("cache-dir"),
		Prompt: autocert.AcceptTOS,
		// HostPolicy: autocert.HostWhitelist(s.opt.ServAddr()),
	}

	t := manager.TLSConfig()

	// tls.Config
	// определяем порт для сервера
	listen, err := tls.Listen("tcp", cfg.ServAddr, t)
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
