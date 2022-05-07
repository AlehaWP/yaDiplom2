package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/config"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/database"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/ossignal"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World</h1>"))
}

func main() {
	//makeMigrations()
	var wg sync.WaitGroup
	logger.NewLogs()
	defer logger.Close()
	logger.Info("Старт клиента")

	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.NewConfig()
	fmt.Println(cfg)
	sDB := database.OpenDBConnect()
	defer sDB.Close()

	wg.Add(1)
	go func() {
		ossignal.HandleQuit(cancel)
		wg.Done()
	}()

	// w := workers.NewWorkersPool(10)
	// defer w.Close()

	// l := accrual.NewSurveyAccrual(sDB.NewDBOrdersRepo(), sDB.NewDBBalanceRepo(), w)
	// go func() {
	// 	l.GetOrdersForSurveyFromDB(ctx)
	// 	wg.Done()
	// }()

	// s := new(server.Server)
	// s.ServerDB = sDB
	// s.Start(ctx)
	wg.Wait()
	logger.Info("Клиент остановлен")

}
