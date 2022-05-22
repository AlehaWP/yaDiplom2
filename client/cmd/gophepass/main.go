package main

import (
	"context"
	"sync"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/database"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/grcp_client"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/input"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/syncing"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/ossignal"
)

func main() {
	//makeMigrations()
	var wg sync.WaitGroup
	logger.NewLogs()
	defer logger.Close()
	logger.Info("Старт клиента")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// cfg := config.NewConfig()
	if ok := database.OpenDBConnect(); !ok {
		return
	}
	defer database.Close()

	grcp_client.Start(ctx)
	defer grcp_client.Close()

	wg.Add(1)
	go func() {
		ossignal.HandleQuit(cancel)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		input.WaitInput(ctx)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		syncing.Syncronize(ctx)
		wg.Done()
	}()

	wg.Wait()
	logger.Info("Клиент остановлен")

}
