package handlers

import "context"

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
