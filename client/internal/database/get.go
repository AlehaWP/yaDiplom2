package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func GetListFiles(ctx context.Context) ([]models.File, error) {
	logger.Info("Запрос файлов пользователя")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id, name FROM files WHERE to_del = false`
	rows, err := pdb.QueryContext(ctx, q)
	if err != nil {
		logger.Info(err)
		return nil, err
	}
	defer rows.Close()

	var lF []models.File
	for rows.Next() {
		var f models.File
		if err := rows.Scan(&f.ID, &f.Name); err != nil {
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

func GetListFilesToServ(ctx context.Context) ([]models.File, error) {
	logger.Info("Запрос файлов пользователя для синхронизации")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT name, data, uuid FROM files WHERE to_serv = true and to_del = false`
	rows, err := pdb.QueryContext(ctx, q)
	if err != nil {
		logger.Info(err)
		return nil, err
	}
	defer rows.Close()

	var lF []models.File
	for rows.Next() {
		var f models.File
		if err := rows.Scan(&f.Name, &f.Data, &f.UUID); err != nil {
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

func GetFile(ctx context.Context, id int) (*models.File, error) {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT name, data FROM files WHERE id=?`
	f := new(models.File)
	f.ID = id
	row := pdb.QueryRowContext(ctx, q, id)
	if err := row.Scan(&f.Name, &f.Data); err != nil && err != sql.ErrNoRows {
		logger.Info(q, err)
		return nil, err
	}
	return f, nil
}

func GetListAccounts(ctx context.Context) ([]models.Account, error) {
	logger.Info("Запрос списка аккаунтов пользователя")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id, login, password FROM accounts WHERE to_del = false`
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

func GetListAccountsToServ(ctx context.Context) ([]models.Account, error) {
	logger.Info("Запрос списка аккаунтов пользователя для синхронизации")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT uuid, login, password FROM accounts WHERE to_del = false and to_serv=true`
	rows, err := pdb.QueryContext(ctx, q)
	if err != nil {
		logger.Info(err)
		return nil, err
	}
	defer rows.Close()

	var lA []models.Account
	for rows.Next() {
		var a models.Account
		if err := rows.Scan(&a.UUID, &a.Login, &a.Password); err != nil {
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

func GetListCards(ctx context.Context) ([]models.Card, error) {
	logger.Info("Запрос спиcка карт пользователя")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id, number, month, year, owner FROM cards WHERE to_del = false`
	rows, err := pdb.QueryContext(ctx, q)
	if err != nil {
		logger.Info(err)
		return nil, err
	}
	defer rows.Close()

	var lC []models.Card
	for rows.Next() {
		var c models.Card
		if err := rows.Scan(&c.ID, &c.Number, &c.Year, &c.Month, &c.Owner); err != nil {
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

func GetListCardsToServ(ctx context.Context) ([]models.Card, error) {
	logger.Info("Запрос спиcка карт пользователя для синхронизации")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT uuid, number, month, year, owner FROM cards WHERE to_del = false and to_serv = true`
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
