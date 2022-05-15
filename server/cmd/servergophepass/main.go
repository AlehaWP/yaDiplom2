package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/database"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/grcp_server"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/logger"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/signal"
)

var (
	BuildVersion string = "N/A"
	BuildDate    string = "N/A"
	BuildCommit  string = "N/A"
)

// Main.
func main() {
	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n", BuildVersion, BuildDate, BuildCommit)
	logger.NewLogs()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ok := database.OpenDBConnect()
	if !ok {
		logger.Error("Ошибка при подключении к БД: ")
		return
	}
	defer database.Close()
	// s := new(server.Server)
	go signal.HandleQuit(cancel)
	// go s.Start(ctx, opt)
	go grcp_server.Start(ctx)
	<-ctx.Done()
}
