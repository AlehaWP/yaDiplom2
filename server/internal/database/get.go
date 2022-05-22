package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/logger"
	"github.com/AlehaWP/yaDiplom2.git/server/internal/models"
)

func GetUser(ctx context.Context, login string) (*models.User, error) {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT uuid, login, password, email, phone FROM users WHERE login = $1 `
	u := new(models.User)
	row := pdb.QueryRowContext(ctx, q, login)

	err := row.Scan(&u.UUID, &u.Login, &u.Password, &u.Email, &u.Phone)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		logger.Info(q, err)
		return nil, err
	}

	return u, nil
}

func GetListFiles(ctx context.Context, u models.User) ([]models.File, error) {
	logger.Info("Запрос файлов пользователя")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT uuid, name FROM files WHERE user_uuid = $1  and to_del!=true`
	rows, err := pdb.QueryContext(ctx, q, u.UUID)
	if err != nil {
		logger.Info(err)
		return nil, err
	}
	defer rows.Close()

	var lF []models.File
	for rows.Next() {
		var f models.File
		if err := rows.Scan(&f.UUID, &f.Name); err != nil {
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

func GetFile(ctx context.Context, uuid string) (*models.File, error) {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT uuid, name, data FROM files WHERE uuid=$1  and to_del!=true `
	f := new(models.File)
	row := pdb.QueryRowContext(ctx, q, uuid)
	if err := row.Scan(&f.UUID, &f.Name, &f.Data); err != nil && err != sql.ErrNoRows {
		logger.Info(q, err)
		return nil, err
	}
	return f, nil
}

func GetListAccounts(ctx context.Context, u models.User) ([]models.Account, error) {
	logger.Info("Запрос спсика аккаунтов пользователя")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT uuid, login, password FROM accounts WHERE user_uuid=$1  and to_del!=true `
	rows, err := pdb.QueryContext(ctx, q, u.UUID)
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

func GetListCards(ctx context.Context, u models.User) ([]models.Card, error) {
	logger.Info("Запрос спиcка карт пользователя")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT uuid, number, month, year, owner FROM cards WHERE user_uuid=$1 and to_del!=true `
	rows, err := pdb.QueryContext(ctx, q, u.UUID)
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
