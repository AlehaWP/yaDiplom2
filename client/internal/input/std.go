package input

import (
	"context"
	"fmt"
	"strings"

	"github.com/AlehaWP/yaDiplom2.git/client/internal/handlers"
	"github.com/AlehaWP/yaDiplom2.git/client/internal/stdin"
	"github.com/AlehaWP/yaDiplom2.git/client/pkg/logger"
)

func WaitInput(ctx context.Context) {

	var suffix string
	var key string

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Программа завершена")
			return
		default:
			key = stdin.Read("Введите команду:")

			if len(key) == 0 {
				break
			}

			if strings.HasSuffix(key, "?") {
				key = key[:len(key)-1]
				suffix = "?"
			}

			if strings.HasPrefix(key, "^") {
				break
			}

			key = strings.Trim(key, " ")

			f, ok := handlers.Handlers[key]
			if !ok {
				fmt.Println("Команда не найдена. Вызвать список доступных команд - ?")
				break
			}
			r, err := f(ctx, suffix)
			if err != nil {
				logger.Error(err)
			}

			fmt.Println(r)
		}
	}

}
