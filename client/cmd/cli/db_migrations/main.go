package main

import (
	"database/sql"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/config"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"
)

func main() {
	logger.NewLogs()
	p := "Миграции базы данных:"
	logger.Info(p, "Старт")
	config := config.NewConfig()
	logger.Info(p, "Подключение к БД")
	db, err := sql.Open("sqlite3", config.DBConnStr)
	if err != nil {
		logger.Error(p, err)
	}

	defer db.Close()
	// setup database
	logger.Info(p, "Применение миграций")
	if err := goose.Up(db, "../../../db/migrations"); err != nil {
		logger.Error(p, err)
	}
	logger.Info(p, "Завершение") // run app
}
