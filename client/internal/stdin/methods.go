package stdin

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func Read(invite string) string {

	// fmt.Print("Введите команду: ")
	// var input string
	// fmt.Scanln(&input)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(invite)
	b, _, _ := reader.ReadLine()
	input := bytes.NewBuffer(b).String()

	return strings.TrimSpace(input)
}
