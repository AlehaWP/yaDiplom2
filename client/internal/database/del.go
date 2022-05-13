package database

import (
	"context"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func DelFile(ctx context.Context, id int) error {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `UPDATE files SET to_del=1 WHERE id=?`
	if _, err := pdb.ExecContext(ctx, q, id); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}

func DelAccount(ctx context.Context, id int) error {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `UPDATE accounts SET to_del=1 WHERE id=?`
	if _, err := pdb.ExecContext(ctx, q, id); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}

func DelCard(ctx context.Context, id int) error {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	q := `UPDATE cards SET to_del=1 WHERE id=?`
	if _, err := pdb.ExecContext(ctx, q, id); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}
