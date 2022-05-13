package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/database"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/stdin"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func delAccount(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда удаления аккаунта из списка")
	}

	la, err := database.GetListAccounts(ctx)
	if err != nil {
		return "", err
	}
	fmt.Println(la)

	n, err := strconv.Atoi(stdin.Read("Укажите порядковый номер аккаунта для удаления"))
	if err != nil {
		logger.Info("порядковый номер не число")
	}

	if err := database.DelAccount(ctx, n); err != nil {
		return "ошибка удаления аккаунта из базы данных", err
	}

	return "аккаунт успешно удален", nil
}

func delCard(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда удаления карты из списка")
	}

	lc, err := database.GetListCards(ctx)
	if err != nil {
		return "", err
	}
	fmt.Println(lc)

	n, err := strconv.Atoi(stdin.Read("Укажите порядковый номер карты для удаления"))
	if err != nil {
		logger.Info("порядковый номер не число")
	}

	if err := database.DelCard(ctx, n); err != nil {
		return "ошибка удаления карты из базы данных", err
	}

	return "карта успешно удалена", nil
}

func delFile(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда удаления файла")
	}

	lf, err := database.GetListFiles(ctx)
	if err != nil {
		return "", err
	}
	fmt.Println(lf)

	n, err := strconv.Atoi(stdin.Read("Укажите порядковый номер файла для удаления"))
	if err != nil {
		logger.Info("порядковый номер файла не число")
	}

	if err := database.DelFile(ctx, n); err != nil {
		return "ошибка удаления файла из базы данных", err
	}

	return "файл успешно удален", nil
}
