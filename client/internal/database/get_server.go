package database

import (
	"context"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func GetFilesToServ(ctx context.Context) ([]models.File, error) {
	logger.Info("{Запрос} файлы для синхронизации")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT uuin, name, data FROM files WHERE to_serv = true and to_del != true`
	rows, err := pdb.QueryContext(ctx, q)
	if err != nil {
		logger.Info(err)
		return nil, err
	}
	defer rows.Close()

	var lF []models.File
	for rows.Next() {
		var f models.File
		if err := rows.Scan(&f.UUID, &f.Name, &f.Data); err != nil {
			logger.Info(err)
			return nil, err
		}
		lF = append(lF, f)
	}
	err = rows.Err()
	if err != nil {
		logger.Info(err)
		return nil, err
	}

	return lF, nil
}

func GetAccountsToServ(ctx context.Context) ([]models.Account, error) {
	logger.Info("Запрос аккаунтов для синхронизации")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT uuid, login, password FROM accounts WHERE to_serv = true and to_del != true`
	rows, err := pdb.QueryContext(ctx, q)
	if err != nil {
		logger.Info(err)
		return nil, err
	}
	defer rows.Close()

	var lA []models.Account
	for rows.Next() {
		var a models.Account
		if err := rows.Scan(&a.ID, &a.Login, &a.Password); err != nil {
			logger.Info(err)
			return nil, err
		}
		lA = append(lA, a)
	}
	err = rows.Err()
	if err != nil {
		logger.Info(err)
		return nil, err
	}

	return lA, nil
}

func GetCardsToServ(ctx context.Context) ([]models.Card, error) {
	logger.Info("Запрос карт для синхронизации")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT uuid, number, month, year, owner FROM cards WHERE to_del != true`
	rows, err := pdb.QueryContext(ctx, q)
	if err != nil {
		logger.Info(err)
		return nil, err
	}
	defer rows.Close()

	var lC []models.Card
	for rows.Next() {
		var c models.Card
		if err := rows.Scan(&c.UUID, &c.Number, &c.Year, &c.Month, &c.Owner); err != nil {
			logger.Info(err)
			return nil, err
		}
		lC = append(lC, c)
	}
	err = rows.Err()
	if err != nil {
		logger.Info(err)
		return nil, err
	}

	return lC, nil
}
