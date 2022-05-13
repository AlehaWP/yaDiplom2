package database

import (
	"context"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

// type DBBalanceRepo struct {
// 	serverDB
// }

// func (db *DBBalanceRepo) Get(ctx context.Context, userID int) (*models.CurrentBalance, error) {
// 	logger.Info("Проверка баланса пользователя")
// 	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
// 	defer cancelfunc()
// 	cb := new(models.CurrentBalance)
// 	cb.UserID = userID
// 	q := `SELECT current_balance, withdrawn FROM customers WHERE user_id=$1`

// 	// q := `select 729.98 as current_balance, sum(bl.sum_out) as withdrawn
// 	// 		from balance_log bl WHERE user_id=$1`
// 	row := db.QueryRowContext(ctx, q, userID)
// 	// row := db.QueryRowContext(ctx, q)

// 	if err := row.Scan(&cb.CurBalance, &cb.Withdrawn); err != nil && err != sql.ErrNoRows {
// 		logger.Info(err)
// 		return nil, err
// 	}
// 	return cb, nil
// }

// func (db *DBBalanceRepo) GetAll(ctx context.Context, userID int) ([]models.BalanceOut, error) {
// 	logger.Info("Запрос заказов пользователя")
// 	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
// 	defer cancelfunc()
// 	q := `SELECT order_id, sum_out, date_add FROM balance_log WHERE user_id=$1 and sum_out != 0`
// 	rows, err := db.QueryContext(ctx, q, userID)
// 	if err != nil {
// 		logger.Info(err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var aBalanceOut []models.BalanceOut
// 	for rows.Next() {
// 		var bo models.BalanceOut
// 		if err := rows.Scan(&bo.OrderID, &bo.Sum, &bo.Processed); err != nil {
// 			if err == sql.ErrNoRows {
// 				return aBalanceOut, nil
// 			}
// 			logger.Info(err)
// 			return nil, err
// 		}
// 		bo.Status = models.OrderStatusProcessed
// 		aBalanceOut = append(aBalanceOut, bo)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		logger.Info(err)
// 		return nil, err
// 	}

// 	return aBalanceOut, nil
// }

func AddAccount(ctx context.Context, a models.Account) error {
	logger.Info("Запрос добавления аккаунта пользователя в БД")
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()

	q := `INSERT INTO accounts (login, password) VALUES (?,?)` // ON CONFLICT (user_id, order_id, sum_in, sum_out)  DO NOTHING `
	if _, err := pdb.ExecContext(ctx, q, a.Login, a.Password); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}

func AddFile(ctx context.Context, f models.File) error {
	logger.Info("Запрос добавления файла пользователя в БД")
	ctx, cancelfunc := context.WithTimeout(ctx, 50*time.Second)
	defer cancelfunc()
	q := `INSERT INTO files (name, data) VALUES (?,?)` // ON CONFLICT (user_id, order_id, sum_in, sum_out)  DO NOTHING `
	if _, err := pdb.ExecContext(ctx, q, f.Name, f.Data); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}

func AddCard(ctx context.Context, c models.Card) error {
	logger.Info("Запрос добавления карты пользователя в БД")
	ctx, cancelfunc := context.WithTimeout(ctx, 50*time.Second)
	defer cancelfunc()
	q := `INSERT INTO cards (number, month, year, owner) VALUES (?,?,?,?)` // ON CONFLICT (user_id, order_id, sum_in, sum_out)  DO NOTHING `
	if _, err := pdb.ExecContext(ctx, q, c.Number, c.Month, c.Year, c.Owner); err != nil {
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
