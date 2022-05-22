package database

import (
	"context"
	"time"

	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func SetSyncing(ctx context.Context, table, uuid string) error {
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()

	q := `UPDATE ` + table + ` set to_serv = false WHERE uuid=$1`
	if _, err := pdb.ExecContext(ctx, q, uuid); err != nil {
		logger.Info(q, err)
		return err
	}
	return nil
}
