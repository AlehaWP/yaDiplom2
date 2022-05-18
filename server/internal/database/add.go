package database

import (
	"context"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/logger"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/models"
)

func AddAccount(ctx context.Context, a models.Account) error {
	logger.Info("Запрос добавления аккаунта пользователя в БД")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()

	q := `INSERT INTO accounts (user_uuid, login, password, uuid) VALUES ($1,$2,$3,$4) ON CONFLICT (uuid) DO NOTHING ` // ON CONFLICT (user_id, order_id, sum_in, sum_out)  DO NOTHING `
	if _, err := pdb.ExecContext(ctx, q, a.User.UUID, a.Login, a.Password, a.UUID); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}

func AddFile(ctx context.Context, f models.File) error {
	logger.Info("Запрос добавления файла пользователя в БД")
	ctx, cancelfunc := context.WithTimeout(ctx, 50*time.Second)
	defer cancelfunc()
	q := `INSERT INTO files (user_uuid, name, data, uuid) VALUES ($1,$2,$3,$4)  ON CONFLICT (uuid) DO NOTHING ` // ON CONFLICT (user_id, order_id, sum_in, sum_out)  DO NOTHING `
	if _, err := pdb.ExecContext(ctx, q, f.User.UUID, f.Name, f.Data, f.UUID); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}

func AddCard(ctx context.Context, c models.Card) error {
	logger.Info("Запрос добавления карты пользователя в БД")
	ctx, cancelfunc := context.WithTimeout(ctx, 50*time.Second)
	defer cancelfunc()
	q := `INSERT INTO cards (user_uuid, number, month, year, owner, uuid) VALUES ($1,$2,$3,$4,$5)  ON CONFLICT (uuid) DO NOTHING ` // ON CONFLICT (user_id, order_id, sum_in, sum_out)  DO NOTHING `
	if _, err := pdb.ExecContext(ctx, q, c.User.UUID, c.Number, c.Month, c.Year, c.Owner, c.UUID); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}
