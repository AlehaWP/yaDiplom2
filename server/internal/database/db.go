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
func (s progDB) CheckDBConnection(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	err := s.PingContext(ctx)
	if err != nil {
		logger.Error("Ошибка подключения к БД", err)
	}
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

	ctx := context.Background()
	cfg := config.NewConfig()
	db, err := sql.Open("postgres", cfg.DBConnStr)
	if err != nil {
		logger.Error("Ошибка подключения к БД", err)
		return false
	}

	pdb.DB = db
	pdb.CheckDBConnection(ctx)
	pdb.createTables(ctx)

	// s.makeMigrations()
	return true
}

func Close() {
	pdb.Close()
}
