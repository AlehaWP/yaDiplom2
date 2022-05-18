package database

import (
	"context"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func AddAccount(ctx context.Context, a models.Account) error {
	logger.Info("Запрос добавления аккаунта пользователя в БД")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()

	q := `INSERT INTO accounts (login, password, uuid) VALUES (?,?,?)` // ON CONFLICT (user_id, order_id, sum_in, sum_out)  DO NOTHING `
	if _, err := pdb.ExecContext(ctx, q, a.Login, a.Password, a.UUID); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}

func AddFile(ctx context.Context, f models.File) error {
	logger.Info("Запрос добавления файла пользователя в БД")
	ctx, cancelfunc := context.WithTimeout(ctx, 50*time.Second)
	defer cancelfunc()
	q := `INSERT INTO files (name, data, uuid) VALUES (?,?,?)` // ON CONFLICT (user_id, order_id, sum_in, sum_out)  DO NOTHING `
	if _, err := pdb.ExecContext(ctx, q, f.Name, f.Data, f.UUID); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}

func AddCard(ctx context.Context, c models.Card) error {
	logger.Info("Запрос добавления карты пользователя в БД")
	ctx, cancelfunc := context.WithTimeout(ctx, 50*time.Second)
	defer cancelfunc()
	q := `INSERT INTO cards (number, month, year, owner, uuid) VALUES (?,?,?,?,?) ON CONFLICT IGNORE` // ON CONFLICT (user_id, order_id, sum_in, sum_out)  DO NOTHING `
	if _, err := pdb.ExecContext(ctx, q, c.Number, c.Month, c.Year, c.Owner, c.UUID); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}

// func (s serverDB) NewDBBalanceRepo() models.BalanceRepo {
// 	br := new(DBBalanceRepo)
// 	br.serverDB = s
// 	return br
// }
