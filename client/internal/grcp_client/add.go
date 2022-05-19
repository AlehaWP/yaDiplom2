package grcp_client

import (
	"context"
	"fmt"

	pb "github.com/AlehaWP/yaDiplom2.git/client/internal/grcp_client/proto"
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
		Uuid:  "123123213sedasdasd",
	}
	fmt.Println("Добавляем")
	resp, _ := c.AddFile(context.Background(), &pb.AddFileRequest{
		File: file,
		User: user,
	})
	fmt.Println("До")
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
