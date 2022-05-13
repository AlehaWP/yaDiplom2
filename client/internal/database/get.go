package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

// import (
// 	"context"
// 	"database/sql"
// 	"time"

// 	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
// 	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
// )

// type DBOrdersRepo struct {
// 	serverDB
// }

// func (db *DBOrdersRepo) Get(ctx context.Context, o *models.Order) (bool, error) {
// 	logger.Info("Проверка наличия заказа")
// 	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
// 	defer cancelfunc()
// 	q := `SELECT id, user_id FROM orders WHERE order_id=$1`
// 	row := db.QueryRowContext(ctx, q, o.OrderID)

// 	if err := row.Scan(&o.ID, &o.UserID); err != nil && err != sql.ErrNoRows {
// 		logger.Info(err)
// 		return false, err
// 	}
// 	if o.ID == 0 {
// 		return false, nil
// 	}
// 	return true, nil
// }

func GetListFiles(ctx context.Context) ([]models.File, error) {
	logger.Info("Запрос файлов пользователя")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id, name FROM files WHERE to_del != true`
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
	logger.Info("Запрос спсика аккаунтов пользователя")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id, login, password FROM accounts WHERE to_del != true`
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

func GetListCards(ctx context.Context) ([]models.Card, error) {
	logger.Info("Запрос спиcка карт пользователя")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT id, number, month, year, owner FROM cards WHERE to_del != true`
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
