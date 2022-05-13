package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/config"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/database"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/stdin"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func saveFile(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда сохранения нужного файла")
	}

	lf, err := database.GetListFiles(ctx)
	if err != nil {
		return "", err
	}
	fmt.Println(lf)

	n, err := strconv.Atoi(stdin.Read("Укажите порядковый номер файла для загрузки"))
	if err != nil {
		logger.Info("порядковый номер файла не число")
		return "ошибка сохранения", err
	}

	f, err := database.GetFile(ctx, n)
	if (err != nil) || (f.Data == nil) {
		return "ошибка загрузки файла из базы данных", err
	}

	d := stdin.Read("укажите каталог для сохранения файла")
	err = ioutil.WriteFile(d+config.PathSeparator+f.Name, f.Data, 0777)
	if err != nil {
		return "ошибка записи файла на диск", err
	}

	return "файл успешно сохранен", nil
}

func saveAccount(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда сохранения нужного аккаунта в файл")
	}

	la, err := database.GetListAccounts(ctx)
	if err != nil {
		return "", err
	}
	fmt.Println(la)

	n, err := strconv.Atoi(stdin.Read("Укажите порядковый номер аккаунта для загрузки"))
	if err != nil {
		logger.Info("порядковый номер аккаунта не число")
		return "ошибка сохранения", err
	}

	for _, v := range la {
		if v.ID != n {
			continue
		}

		d := stdin.Read("укажите каталог для сохранения файла")
		f, err := os.OpenFile(d+config.PathSeparator+v.Login+".json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return "ошибка сохранения файла", err
		}
		defer f.Close()
		jsonParser := json.NewEncoder(f)
		jsonParser.Encode(v)

	}

	return "аккаунт успешно сохранен", nil
}

func saveCard(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда сохранения карты в файл")
	}

	lс, err := database.GetListCards(ctx)
	if err != nil {
		return "", err
	}
	fmt.Println(lс)

	n, err := strconv.Atoi(stdin.Read("Укажите порядковый номер карты для загрузки"))
	if err != nil {
		logger.Info("порядковый номер не число")
		return "ошибка сохранения", err
	}

	for _, v := range lс {
		if v.ID != n {
			continue
		}

		d := stdin.Read("укажите каталог для сохранения файла")
		f, err := os.OpenFile(d+config.PathSeparator+"card"+strconv.Itoa(v.ID)+".json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return "ошибка сохранения файла", err
		}
		defer f.Close()
		jsonParser := json.NewEncoder(f)
		jsonParser.Encode(v)

	}

	return "файл успешно сохранен", nil
}
