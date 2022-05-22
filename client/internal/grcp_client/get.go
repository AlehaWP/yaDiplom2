package grcp_client

import (
	"context"
	"encoding/json"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/database"
	pb "github.com/AlehaWP/yaDiplom2.git/client/internal/grcp_client/proto"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func (cl Client) GetFileList(ctx context.Context) {

	user := &pb.User{
		Uuid: user.UUID,
	}

	resp, _ := cl.pbcl.GetFileList(ctx, &pb.GetDataRequest{
		User: user,
	})
	lf := []models.File{}

	err := json.Unmarshal(resp.Data, &lf)
	if err != nil {
		logger.Info("GetFileList", "Ошибка маршализации", err)
		return
	}

	for _, f := range lf {
		err := database.AddFile(ctx, f)
		if err != nil {
			logger.Info("GetFileList", "Ошибка сохранения файла", err)
		}
	}

}

func (cl Client) GetAccList(ctx context.Context) {

	user := &pb.User{
		Uuid: user.UUID,
	}

	resp, _ := cl.pbcl.GetAcc(ctx, &pb.GetDataRequest{
		User: user,
	})
	la := []models.Account{}

	err := json.Unmarshal(resp.Data, &la)
	if err != nil {
		logger.Info("GetAccList", "Ошибка маршализации", err)
		return
	}

	for _, a := range la {
		err := database.AddAccount(ctx, a)
		if err != nil {
			logger.Info("GetAccList", "Ошибка сохранения файла", err)
		}
	}

}

func (cl Client) GetCardList(ctx context.Context) {

	user := &pb.User{
		Uuid: user.UUID,
	}

	resp, _ := cl.pbcl.GetCard(ctx, &pb.GetDataRequest{
		User: user,
	})
	lc := []models.Card{}

	err := json.Unmarshal(resp.Data, &lc)
	if err != nil {
		logger.Info("GetCardList", "Ошибка маршализации", err)
		return
	}

	for _, c := range lc {
		err := database.AddCard(ctx, c)
		if err != nil {
			logger.Info("GetAccList", "Ошибка сохранения файла", err)
		}
	}

}
