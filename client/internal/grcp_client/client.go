package grcp_client

import (
	"context"
	"fmt"

	pb "github.com/AlehaWP/yaDiplom2.git/client/internal/grcp_client/proto"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
	"google.golang.org/grpc"
)

var user models.User

var conn *grpc.ClientConn

type Client struct {
	pbcl     pb.GophePassClient
	userUUID string
}

func (c Client) CheckConnection(ctx context.Context) bool {
	r := new(pb.CheckRequest)
	a, err := c.pbcl.CheckConnection(ctx, r)
	if err != nil {
		logger.Info("Нет соединения с сервером", err)
		return false
	}

	if a.Ok != true {
		return false
	}
	return true
}

func NewClient() (Client, bool) {
	if conn == nil {
		return Client{}, false
	}
	if user.UUID == "" {
		return Client{}, false
	}
	c := Client{
		pbcl:     pb.NewGophePassClient(conn),
		userUUID: user.UUID,
	}
	return c, true
}

func Start(ctx context.Context) {
	var ok bool
	var err error
	conn, err = grpc.Dial(":3200", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	c := Client{
		pbcl: pb.NewGophePassClient(conn),
	}

	if ok := c.CheckConnection(ctx); ok != true {
		logger.Info("Нет соединения с сервером")
		fmt.Println("Нет соединения с сервером")
		return
	}

	user, ok, err = c.login(ctx)
	if err != nil {
		logger.Info("Ошибка регистрации на сервере", err)
		fmt.Println("Синхронизация на данный момент не возможна")
		return
	}

	if !ok {
		fmt.Println("Синхронизация на данный момент не возможна")
		return
	}

	// AddFile(c)
	// GetFileList(c)
}

func Close() {
	conn.Close()
}
