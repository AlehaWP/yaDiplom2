package ossignal

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func HandleQuit(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-c
	logger.Info("Получен сигнал на закрытие клиента")
	cancel()
}
