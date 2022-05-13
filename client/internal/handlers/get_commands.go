package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/database"
)

func listFile(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда отображения доступного списка файлов")
	}
	lf, err := database.GetListFiles(ctx)
	if err != nil {
		return "", err
	}
	r := ""

	for _, v := range lf {
		r = r + strconv.Itoa(v.ID) + " " + v.Name + "\n"
	}

	return r, nil
}

func listAccount(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда отображения списка сохраненных аккаунтов")
	}
	la, err := database.GetListAccounts(ctx)
	if err != nil {
		return "", err
	}
	r := ""

	for _, v := range la {
		r = r + strconv.Itoa(v.ID) + " " + v.Login + " " + v.Password + "\n"
	}

	return r, nil
}

func listCard(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда отображения списка сохраненных карт")
	}
	lc, err := database.GetListCards(ctx)
	if err != nil {
		return "", err
	}
	r := ""

	for _, v := range lc {
		r = r + strconv.Itoa(v.ID) + " " + v.Number + "\n"
	}

	return r, nil
}
