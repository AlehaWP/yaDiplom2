package grcp_client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func Start(ctx context.Context) {
	// cfg := config.NewConfig()

	// определяем порт для сервера
	conn, err := grpc.Dial(":3200", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	defer conn.Close()

	// c := pb.NewGophePassClient(conn)

	// AddFile(c)
	// GetFileList(c)
	// регистрируем сервис
	// pb.RegisterGophePassServer(s, &GophePassServer{})

	// fmt.Println("сервер gRPC начал работу")
	// // получаем запрос gRpc
	// if err := s.Serve(listen); err != nil {
	// 	fmt.Println(err)
	// }
}
