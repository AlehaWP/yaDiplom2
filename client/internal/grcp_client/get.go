package grcp_client

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/AlehaWP/yaDiplom2.git/client/internal/grcp_client/proto"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func GetFileList(c pb.GophePassClient) {

	user := &pb.User{
		Login: "Тест юзер",
		Uuid:  "123123213sedasdasd",
	}

	resp, _ := c.GetFileList(context.Background(), &pb.GetDataRequest{
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
