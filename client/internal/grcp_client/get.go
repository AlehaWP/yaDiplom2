package grcp_client

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/AlehaWP/yaDiplom2.git/client/internal/grcp_client/proto"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func (c Client) GetFileList(ctx context.Context) {

	user := &pb.User{
		Uuid: user.UUID,
	}

	resp, _ := c.pbcl.GetFileList(ctx, &pb.GetDataRequest{
		User: user,
	})
	lf := []models.File{}

	err := json.Unmarshal(resp.Data, &lf)
	if err != nil {
		logger.Info("GetFileList", "Ошибка маршализации", err)
	}
	fmt.Println(resp.Message)
	fmt.Println(lf)

}
