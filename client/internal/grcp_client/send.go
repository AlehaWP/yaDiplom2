package grcp_client

import (
	"context"

	pb "github.com/AlehaWP/yaDiplom2.git/client/internal/grcp_client/proto"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func (c Client) SendFile(ctx context.Context, f models.File) {
	file := &pb.File{
		Name: f.Name,
		Data: f.Data,
		Uuid: f.UUID,
	}

	user := &pb.User{
		Uuid: c.userUUID,
	}
	resp, _ := c.pbcl.AddFile(ctx, &pb.AddFileRequest{
		File: file,
		User: user,
	})
	logger.Info("Отправка файла на сервер", f.UUID, resp.Message)

}

func (c Client) SendCard(ctx context.Context, ca models.Card) {
	card := &pb.Card{
		Uuid:   ca.UUID,
		Number: ca.Number,
		Month:  int32(ca.Month),
		Year:   int32(ca.Year),
		Owner:  ca.Owner,
	}

	user := &pb.User{
		Uuid: c.userUUID,
	}
	resp, _ := c.pbcl.AddCard(ctx, &pb.AddCardRequest{
		Card: card,
		User: user,
	})
	logger.Info("Отправка карты на сервер", ca.UUID, resp.Message)

}

func (c Client) SendAcc(ctx context.Context, a models.Account) {
	acc := &pb.Account{
		Uuid:     a.UUID,
		Login:    a.Login,
		Password: a.Password,
	}

	user := &pb.User{
		Uuid: c.userUUID,
	}
	resp, _ := c.pbcl.AddAcc(ctx, &pb.AddAccRequest{
		Account: acc,
		User:    user,
	})
	logger.Info("Отправка аккаунта на сервер", a.UUID, resp.Message)

}
