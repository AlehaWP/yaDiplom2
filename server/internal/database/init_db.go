package database

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/config"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/logger"
	_ "github.com/lib/pq"
)

type progDB struct {
	*sql.DB
}

var (
	pdb  progDB
	once sync.Once
)

//CheckDBConnection trying connect to db.
func CheckDBConnection(ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	err := pdb.PingContext(ctx)
	if err != nil {
		logger.Error("Ошибка подключения к БД", err)
		return false
	}
	return true
}

// makeMigrations start here for autotests
// func (s *DB) makeMigrations() {
// 	p := "Миграции базы данных:"
// 	logger.Info(p, "Старт")
// 	// setup database
// 	logger.Info(p, "Применение миграций")
// 	if err := goose.Up(s.DB, "../../db/migrations"); err != nil {
// 		logger.Error(p, err)
// 	}
// 	logger.Info(p, "Завершение") // run app
// }

func OpenDBConnect() bool {
	var err error
	ctx := context.Background()
	cfg := config.NewConfig()

	once.Do(func() {
		pdb.DB, err = sql.Open("postgres", cfg.DBConnStr)
		if err == nil {
			CheckDBConnection(ctx)
			createTables(ctx)
		}
	})
	if err != nil {
		logger.Error("Ошибка подключения к БД", err)
		return false
	}

	// s.makeMigrations()
	return true
}

func Close() {
	pdb.Close()
}
