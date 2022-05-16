package grcp_client

import (
	"context"
	"fmt"

	pb "github.com/AlehaWP/yaDiplom2.git/client/internal/grcp_client/proto"
	"google.golang.org/grpc"
)

// AddUser реализует интерфейс добавления пользователя.
func AddFile(c pb.GophePassClient) {
	file := &pb.File{
		Name: "111.txt",
		Data: []byte("Тест соединение"),
		Uuid: "123",
	}

	user := &pb.User{
		Login: "Тест юзер",
	}

	resp, _ := c.AddFile(context.Background(), &pb.AddFileRequest{
		File: file,
		User: user,
	})
	fmt.Println(resp.Message)

}

// AddUser реализует интерфейс добавления пользователя.
func AddAcc(c pb.GophePassClient) {

	acc := &pb.Account{
		Login:    "asdqwewqe",
		Password: "asdqwe1237283120938",
		Uuid:     "123",
	}

	user := &pb.User{
		Uuid: "123123213sedasdasd",
	}
	resp, _ := c.AddAcc(context.Background(), &pb.AddAccRequest{
		Account: acc,
		User:    user,
	})
	fmt.Println(resp.Message)

}

func Start(ctx context.Context) {
	// cfg := config.NewConfig()

	// определяем порт для сервера
	conn, err := grpc.Dial(":3200", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	defer conn.Close()

	c := pb.NewGophePassClient(conn)

	AddAcc(c)
	// регистрируем сервис
	// pb.RegisterGophePassServer(s, &GophePassServer{})

	// fmt.Println("сервер gRPC начал работу")
	// // получаем запрос gRpc
	// if err := s.Serve(listen); err != nil {
	// 	fmt.Println(err)
	// }
}
