package yadro

import (
	"Project_one/todo"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Comands string //Алиасы команд
const (
	add_task    Comands = "at"
	show_list   Comands = "lt"
	update_task Comands = "ut"
	delete_task Comands = "dt"
)

func Show_information() { //Функция на help коотрая будет выводить информацию о командах
	fmt.Println("Information about commands:")
	fmt.Println("To add task write:  at [tag] [task]")
	fmt.Println("To show the list write:  lt")
	fmt.Println("To update the task write:  ut <index> [new_status] [new_task]")
	fmt.Println("To delete the task:  dt <index>")
	fmt.Println("Available tags: study, housework, finance, health, work or nothing")
	fmt.Println("Available status: done, not done yet")
	fmt.Println("In <> you have the required arguments, and in [] optional, if you want to skip them print ---")

}

func Check_tag(arg string) string {
	switch todo.Tags(arg) { //Проверили тег на валидность
	case "---":
		return string(todo.None)
	case todo.Study:
		return string(todo.Study)
	case todo.Housework:
		return string(todo.Housework)
	case todo.Finance:
		return string(todo.Finance)
	case todo.Health:
		return string(todo.Health)
	case todo.Work:
		return string(todo.Work)
	default:
		return "<->"
	}
}

type Command interface {
	Run(args []string) error
}

type Add_commannd struct {
	common_info string
}

// Формат: at [tag] [task], если что то отстутвует - то писать надо ---
func (v *Add_commannd) Run(args []string) error { //Первый аргумент - тег (если он есть) Если его нет, передается None и мы продолжаем работу, если он не валидный выбрасываем ошибку сообщением
	info_tag := Check_tag(args[0])
	if info_tag == "<->" {
		return errors.New("uknown tag, please choose one of the available ones")
	}
	var format_task todo.Task
	format_task.Tag_data = todo.Tags(info_tag)
	info := strings.Join(args[1:], " ")
	format_task.Todo_data = info
	our_todo_list.Add(format_task) //Добавим задачу к нашему списку
	v.common_info = strings.Join(args, " ")
	return nil
}

type Show_list_command struct {
	common_info string
}

// lt - можно добавить чтоб выводил как срез 1:4 к примеру
func (v *Show_list_command) Run(args []string) error {
	our_todo_list.List()
	v.common_info = strings.Join(args, " ")
	return nil
}

type Update_task_command struct {
	common_info string
}

// ut <index> [status] [task], если что то отстутвует - то писать надо ---
func (v *Update_task_command) Run(args []string) error { //Формат ввода: <Номер задачи - Ind> [New status] [New data]  пропуск необязательного параметра симво "---"
	if len(args) < 2 {
		return errors.New("please make the update by format")
	}
	ind, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("please enter the index")
	}
	if ind > our_todo_list.Len {
		return errors.New("index of the task is out of range")
	}
	ind -= 1

	var format_task todo.Task

	switch args[1] {
	case "done":
		format_task.Is_done = true
	case "not done yet":
		format_task.Is_done = false
	case "---":
		format_task.Is_done = our_todo_list.Daily_map[ind].Is_done
	default:
		return errors.New("please chose one of the available status")
	}
	if len(args) == 2 {
		return nil
	}

	if args[2] != "---" {
		format_task.Todo_data = strings.Join(args[2:], " ")
	} else {
		format_task.Todo_data = our_todo_list.Daily_map[ind].Todo_data
	}

	our_todo_list.Update(format_task, ind)
	v.common_info = strings.Join(args, " ")
	return nil
}

type Delete_command struct {
	common_info string
}

// dt <ind> ind- индекс задачи которую надо удалить
func (v *Delete_command) Run(args []string) error {
	ind, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("please enter the index")
	}

	if ind > (our_todo_list.Len) {
		return errors.New("task index is out of range")
	}
	ind -= 1

	our_todo_list.Delete(ind)
	v.common_info = strings.Join(args, " ")
	return nil

}

type CLI_Yadro struct {
	map_of_commands map[int]Command
	len_of_the_map  int
	our_todo_list   *todo.Todo_list
}

func NewCLI_Yadro() *CLI_Yadro { //Конструктор чтоб не работать с nil указателем
	var m = todo.NewTodo_list()
	return &CLI_Yadro{
		map_of_commands: make(map[int]Command),
		our_todo_list:   m,
	}
}

func (v *CLI_Yadro) Pars_and_run_command(args []string) error {

	comm := args[0]
	switch comm {
	case "at":
		var param Add_commannd
		err := param.Run(args[1:])
		if err != nil {
			fmt.Println("Error:")
			return err
		}
		param.common_info = strings.Join(args, " ")
		v.map_of_commands[v.len_of_the_map+1] = &param //Тонкйи момент-всё из за того что метод принимает указательна структуру
		v.len_of_the_map += 1
		return nil
	case "lt":
		var param Show_list_command
		err := param.Run(args[1:])
		if err != nil {
			return err
		}
		param.common_info = strings.Join(args, " ")
		v.map_of_commands[v.len_of_the_map+1] = &param
		v.len_of_the_map += 1
		return nil
	case "ut":
		var param Update_task_command
		err := param.Run(args[1:])
		if err != nil {
			return err
		}
		param.common_info = strings.Join(args, " ")
		v.map_of_commands[v.len_of_the_map+1] = &param
		v.len_of_the_map += 1
		return nil
	case "dt":
		var param Delete_command
		err := param.Run(args[1:])
		if err != nil {
			return err
		}
		param.common_info = strings.Join(args, " ")
		v.map_of_commands[v.len_of_the_map+1] = &param
		v.len_of_the_map += 1
		return nil
	default:
		err := errors.New("error: Unacceptable command  please choose one of the available ones")
		return err
	}
}

var our_todo_list = todo.NewTodo_list() //Сам список с которым буждем работать
