package database

import (
	"context"

	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func (s progDB) createTables(ctx context.Context) {
	tx, err := s.DB.Begin()
	if err != nil {
		logger.Panic("Ошибка создания таблиц", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS files (
								    id INTEGER PRIMARY KEY,
									uuid VARCHAR(36) UNIQUE ON CONFLICT IGNORE,
									name VARCHAR(255),
									data BLOB,
									to_client BOOLEAN default false,
									to_serv BOOLEAN default true,
									to_del BOOLEAN default false)
	`)
	if err != nil {
		logger.Panic("Ошибка создания таблиц", err)
	}

	_, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS cards (
									id INTEGER PRIMARY KEY,
									uuid VARCHAR(36) UNIQUE ON CONFLICT IGNORE,
									number VARCHAR(20),
									month INTEGER,
									year INTEGER,
									owner VARCHAR(100),
									to_client BOOLEAN default false,
 									to_serv BOOLEAN default true,
									to_del BOOLEAN default false)
	`)
	if err != nil {
		logger.Panic("Ошибка создания таблиц", err)
	}
	_, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS accounts (
									id INTEGER PRIMARY KEY,
									uuid VARCHAR(36) UNIQUE ON CONFLICT IGNORE,
									login VARCHAR(255),
									password VARCHAR(255),
									to_client BOOLEAN default false,
									to_serv BOOLEAN default true,
									to_del BOOLEAN default false)
	`)
	if err != nil {
		logger.Panic("Ошибка создания таблиц", err)
	}
	// _, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS orders (
	// 								id SERIAL PRIMARY KEY,
	// 								user_id INT NOT NULL,
	// 								order_id VARCHAR(50) UNIQUE,
	// 								accrual NUMERIC default 0,
	// 								order_status VARCHAR(20),
	// 								date_add TIMESTAMPTZ(0) default (NOW() at time zone 'UTC+3'))
	// `)
	// if err != nil {
	// 	logger.Panic("Ошибка создания таблиц", err)
	// }
	// _, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS balance_log (
	// 								id SERIAL PRIMARY KEY,
	// 								user_id INT NOT NULL,
	// 								order_id VARCHAR(50),
	// 								sum_in NUMERIC default 0,
	// 								sum_out NUMERIC default 0,
	// 								date_add TIMESTAMPTZ(0) default (NOW() at time zone 'UTC+3'),
	// 								CONSTRAINT orders_uk UNIQUE (user_id, order_id, sum_in, sum_out))
	// `)
	// if err != nil {
	// 	logger.Panic("Ошибка создания таблиц", err)
	// }

	// _, err = tx.ExecContext(ctx, `CREATE OR REPLACE VIEW customers AS
	// 								select bl.user_id, sum(bl.sum_in) as sum_in,sum(bl.sum_out) as withdrawn,
	// 								sum(bl.sum_in) - sum(bl.sum_out) as current_balance,
	// 								MAX(bl.date_add) as last_event_date from balance_log bl
	// 								group by bl.user_id
	// `)
	// if err != nil {
	// 	logger.Panic("Ошибка создания таблиц", err)
	// }

	tx.Commit()
}
