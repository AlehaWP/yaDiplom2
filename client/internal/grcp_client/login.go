package grcp_client

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/AlehaWP/yaDiplom2.git/client/internal/grcp_client/proto"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/models"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/stdin"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func (c Client) login(ctx context.Context) (models.User, bool, error) {
	user := models.User{}
	fmt.Println("Для синхронизации с сервером, введите рег.данные пользователя")
	// r := &pb.LoginRequest{
	// 	User: &pb.User{},
	// }
	r := new(pb.LoginRequest)
	r.User = new(pb.User)
	r.User.Login = stdin.Read("Логин")
	r.User.Password = stdin.Read("Введите пароль")

	a, err := c.pbcl.Login(ctx, r)
	if err != nil {
		logger.Info("Ошибка регистрации", err)
		fmt.Println("Ошибка регистрации", err)
		return user, false, err
	}

	if a.Finded == false {
		fmt.Println("Пользователь не найден")
		t := stdin.Read("Зарегистрировать введенные данные пользователя на сервере? Д/Н")
		if strings.ToUpper(t) != "Д" {
			return user, false, nil
		}

		user, err = c.regUser(ctx, a.User)
		if err != nil {
			return user, false, err
		}

		return user, true, nil
	}

	if a.Auth == false {
		fmt.Println("Не верная связка пользователь/пароль")
		return user, false, err
	}
	fmt.Println("Авторизация прошла успешно")
	user.UUID = a.User.Uuid
	return user, true, nil
}

func (c Client) regUser(ctx context.Context, u *pb.User) (models.User, error) {
	user := models.User{}
	r := &pb.LoginRequest{
		User: u,
	}
	a, err := c.pbcl.RegUser(ctx, r)
	if err != nil {
		fmt.Println("Ошибка регистрации данных на сервере")
		return user, err
	}
	fmt.Println("Регистрация прошла успешно")
	user.UUID = a.User.Uuid
	return user, nil
}
