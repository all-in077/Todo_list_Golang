package main

import (
	"Project_one/yadro"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var res = yadro.NewCLI_Yadro()

	scanner := bufio.NewScanner(os.Stdin) //Запустим сканнер
	for scanner.Scan() {
		text := scanner.Text() //text - это простой string

		if text == "" || text == "exit" { //Если введем пустую строку то ввод так же завершится
			break
		}

		if text == "help" || text == "Help" {
			yadro.Show_information()
			continue
		}

		//<Command>
		words_buff := strings.Fields(text) //массив строковых аргументов
		err := res.Pars_and_run_command(words_buff)
		if err != nil {
			fmt.Println(err)
		}

	}
}
