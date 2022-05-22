package syncing

import (
	"context"
	"fmt"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/grcp_client"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

var cl grcp_client.Client

func Syncronize(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Служба синхронизации завершена")
			return
		default:
			var ok bool
			cl, ok = grcp_client.NewClient()

			if !ok {
				logger.Info("Нет соединения с сервером, синхронизация не возможна")
				return
			}

			sendFiles(ctx)
			sendCards(ctx)
			sendAccs(ctx)
			time.Sleep(30 * time.Second)
		}
	}
}
