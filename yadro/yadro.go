package yadro

import (
	todo "Project_one/tasks"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Comands string // Алиасы команд
const (
	add_task    Comands = "at"
	show_list   Comands = "lt"
	update_task Comands = "ut"
	delete_task Comands = "dt"
)

type Command interface {
	Run(args []string, ourtodolist *todo.TodoList) error
}

type AddCommand struct {
	commoninfo string
}

type ShowListCommand struct {
	commoninfo string
}

type UpdateTaskCommand struct {
	commoninfo string
}

type DeleteCommand struct {
	commoninfo string
}

type CLIYadro struct {
	mapOfCommands map[int]Command
	lenOfTheMap   int
	ourtodolist   todo.TodoList
}

func ShowInformation() { //Функция на help коотрая будет выводить информацию о командах
	fmt.Println("Information about commands:")
	fmt.Println("To add task write:  at [tag] [task]")
	fmt.Println("To show the list write:  lt")
	fmt.Println("To update the task write:  ut <index> [new_status] [new_task]")
	fmt.Println("To delete the task:  dt <index>")
	fmt.Println("Available tags: study, housework, finance, health, work or nothing")
	fmt.Println("Available status: done, not done yet")
	fmt.Println("In <> you have the required arguments, and in [] optional, if you want to skip them print ---")
}

func CheckTag(arg string) string {
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

// Формат: at [tag] [task], если что то отстутвует - то писать надо ---
func (v *AddCommand) Run(args []string, ourtodolist *todo.TodoList) error { //Первый аргумент - тег (если он есть) Если его нет, передается None и мы продолжаем работу, если он не валидный выбрасываем ошибку сообщением
	info_tag := CheckTag(args[0])
	if info_tag == "<->" {
		return errors.New("uknown tag, please choose one of the available ones")
	}
	var format_task todo.Task
	format_task.TagData = todo.Tags(info_tag)
	info := strings.Join(args[1:], " ")
	format_task.TodoData = info
	ourtodolist.Add(format_task) //Добавим задачу к нашему списку
	v.commoninfo = strings.Join(args, " ")
	return nil
}

// lt
func (v *ShowListCommand) Run(args []string, ourtodolist *todo.TodoList) error {
	ourtodolist.List()
	v.commoninfo = strings.Join(args, " ")
	return nil
}

// ut <index> [status] [task], если что то отстутвует - то писать надо ---
func (v *UpdateTaskCommand) Run(args []string, ourtodolist *todo.TodoList) error { //Формат ввода: <Номер задачи - Ind> [New status] [New data]  пропуск необязательного параметра симво "---"
	if len(args) < 2 {
		return errors.New("please make the update by format")
	}
	ind, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("please enter the index")
	}
	if ind > ourtodolist.Len {
		return errors.New("index of the task is out of range")
	}
	ind -= 1

	var format_task todo.Task

	switch args[1] {
	case "done":
		format_task.IsDone = true
	case "not done yet":
		format_task.IsDone = false
	case "---":
		format_task.IsDone = ourtodolist.Daily_map[ind].IsDone
	default:
		return errors.New("please chose one of the available status")
	}
	if len(args) == 2 {
		return nil
	}

	if args[2] != "---" {
		format_task.TodoData = strings.Join(args[2:], " ")
	} else {
		format_task.TodoData = ourtodolist.Daily_map[ind].TodoData
	}

	ourtodolist.Update(format_task, ind)
	v.commoninfo = strings.Join(args, " ")
	return nil
}

// dt <ind> ind- индекс задачи которую надо удалить
func (v *DeleteCommand) Run(args []string, ourtodolist *todo.TodoList) error {
	ind, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("please enter the index")
	}

	if ind > (ourtodolist.Len) {
		return errors.New("task index is out of range")
	}
	ind -= 1

	ourtodolist.Delete(ind)
	v.commoninfo = strings.Join(args, " ")
	return nil
}

func NewCLIYadro() *CLIYadro { //Конструктор чтоб не работать с nil указателем
	var m = todo.NewTodoList()
	return &CLIYadro{
		mapOfCommands: make(map[int]Command),
		ourtodolist:   *m,
	}
}

func (v *CLIYadro) ParsAndRunCommand(args []string) error {
	comm := args[0]
	switch comm {
	case "at":
		var param AddCommand
		err := param.Run(args[1:], &v.ourtodolist)
		if err != nil {
			fmt.Println("Error:")
			return err
		}
		param.commoninfo = strings.Join(args, " ")
		v.mapOfCommands[v.lenOfTheMap+1] = &param //Тонкйи момент-всё из за того что метод принимает указательна структуру
		v.lenOfTheMap += 1
		return nil
	case "lt":
		var param ShowListCommand
		err := param.Run(args[1:], &v.ourtodolist)
		if err != nil {
			return err
		}
		param.commoninfo = strings.Join(args, " ")
		v.mapOfCommands[v.lenOfTheMap+1] = &param
		v.lenOfTheMap += 1
		return nil
	case "ut":
		var param UpdateTaskCommand
		err := param.Run(args[1:], &v.ourtodolist)
		if err != nil {
			return err
		}
		param.commoninfo = strings.Join(args, " ")
		v.mapOfCommands[v.lenOfTheMap+1] = &param
		v.lenOfTheMap += 1
		return nil
	case "dt":
		var param DeleteCommand
		err := param.Run(args[1:], &v.ourtodolist)
		if err != nil {
			return err
		}
		param.commoninfo = strings.Join(args, " ")
		v.mapOfCommands[v.lenOfTheMap+1] = &param
		v.lenOfTheMap += 1
		return nil
	default:
		err := errors.New("error: Unacceptable command  please choose one of the available ones")
		return err
	}
}
