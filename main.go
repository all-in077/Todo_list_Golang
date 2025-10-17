package main

import (
	"Project_one/yadro"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var res = yadro.NewCLIYadro()

	scanner := bufio.NewScanner(os.Stdin) //Запустим сканнер
	for scanner.Scan() {
		text := scanner.Text() //text - это простой string

		if text == "" || text == "exit" { //Если введем пустую строку то ввод так же завершится
			break
		}

		if text == "help" || text == "Help" {
			yadro.ShowInformation()
			continue
		}

		//<Command>
		wordsBuff := strings.Fields(text) //массив строковых аргументов
		err := res.ParsAndRunCommand(wordsBuff)
		if err != nil {
			fmt.Println(err)
		}

	}
}
