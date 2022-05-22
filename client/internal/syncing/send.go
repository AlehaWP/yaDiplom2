package syncing

import (
	"context"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/database"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func sendFiles(ctx context.Context) {
	lf, err := database.GetListFilesToServ(ctx)
	if err != nil {
		logger.Info("Ошибка получения файлов для отправки на сервер", err)
		return
	}

	for _, f := range lf {
		cl.SendFile(ctx, f)
		database.SetSyncing(ctx, "files", f.UUID)
	}
}

func sendCards(ctx context.Context) {
	lс, err := database.GetListCardsToServ(ctx)
	if err != nil {
		logger.Info("Ошибка получения карт для отправки на сервер", err)
		return
	}

	for _, ca := range lс {
		cl.SendCard(ctx, ca)
		database.SetSyncing(ctx, "cards", ca.UUID)
	}
}

func sendAccs(ctx context.Context) {
	la, err := database.GetListAccountsToServ(ctx)
	if err != nil {
		logger.Info("Ошибка получения аккаунтов для отправки на сервер", err)
		return
	}

	for _, a := range la {
		cl.SendAcc(ctx, a)
		database.SetSyncing(ctx, "accounts", a.UUID)
	}
}
