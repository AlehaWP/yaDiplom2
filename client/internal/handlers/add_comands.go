package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/database"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/stdin"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
	"github.com/google/uuid"
)

var Handlers map[string]func(context.Context, string) (string, error)

func help(ctx context.Context, str string) (string, error) {
	fmt.Println("список доступных команд")
	fmt.Println("более подробное описание команды: <команда> ?")
	for key := range Handlers {
		fmt.Println(key)
	}
	return "", nil
}

func addAccount(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда добавления связки логин-пароль")
	}
	account := models.Account{
		UUID:     uuid.New().String(),
		Login:    stdin.Read("Укажите логин:"),
		Password: stdin.Read("Укажите пароль:"),
	}

	err := database.AddAccount(ctx, account)
	if err != nil {
		return "", err
	}

	return "пароль успешно добавлен", nil
}

func addFile(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда добавления файла")
	}
	f := stdin.Read("Укажите файл для загрузки:")
	d, err := ioutil.ReadFile(f)
	if err != nil {
		logger.Info("ошибка чтения файла", err)
	}
	fn := filepath.Base(f)

	// fmt.Println(dat)
	// err = ioutil.WriteFile(`/home/kseykseich/Go/github.com/AlehaWP/yaDiplom2.git/client/gophe`, dat, 0777)

	mf := models.File{
		UUID: uuid.New().String(),
		Name: fn,
		Data: d,
	}

	err = database.AddFile(ctx, mf)
	if err != nil {
		return "", err
	}
	// login := stdin.Read("Укажите логин:")
	// password := stdin.Read("Укажите пароль:")

	return "файл успешно добавлен", nil
}

func addCard(ctx context.Context, suffix string) (string, error) {
	if suffix == "?" {
		fmt.Println("команда добавления банковской карты")
	}
	n := stdin.Read("Укажите номер карты")
	m, err := strconv.Atoi(stdin.Read("Укажите месяц:"))
	if err != nil {
		logger.Error("не верный формат ввода месяца")
	}
	y, err := strconv.Atoi(stdin.Read("Укажите год:"))
	if err != nil {
		logger.Error("не верный формат ввода месяца")
	}
	o := stdin.Read("Укажите владельца")
	card := models.Card{
		UUID:   uuid.New().String(),
		Number: n,
		Month:  m,
		Year:   y,
		Owner:  o,
	}

	err = database.AddCard(ctx, card)
	if err != nil {
		return "", err
	}

	return "файл успешно добавлен", nil
}

func init() {
	Handlers = map[string]func(context.Context, string) (string, error){
		"":             help,
		"login":        help,
		"add_account":  addAccount,
		"list_account": listAccount,
		"save_account": saveAccount,
		"del_account":  delAccount,
		"add_card":     addCard,
		"list_card":    listCard,
		"save_card":    saveCard,
		"del_card":     delCard,
		"add_file":     addFile,
		"list_file":    listFile,
		"save_file":    saveFile,
		"del_file":     delFile,
	}
}
