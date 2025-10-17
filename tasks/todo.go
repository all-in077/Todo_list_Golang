package tasks

import (
	"fmt"
)

type Tags string // Перечисление для тегов
const (
	None      Tags = ""
	Study     Tags = "study"
	Housework Tags = "housework"
	Finance   Tags = "finance"
	Health    Tags = "health"
	Work      Tags = "work"
)

type Task struct {
	TodoData string //Самма информацию про задачу - что надо сдедать
	TagData  Tags   ////Тег задачи - по умлочанию none - не забудь пропистаь в конструкторе (идея напистаь один конструктор через функциональные параметры)
	IsDone   bool   //Сделана ли задача - поле статуса ??? - надо еще подумать

}

type TodoList struct {
	Len        int
	Daily_list []Task
	Daily_map  map[int]Task // Не забыть в main явно прописать её при создании пустого списка через : чтоб не было ошибки работы с nil мап
}

func NewTodoList() *TodoList { //Чтобы не было ошибки записи в nil map
	return &TodoList{
		Daily_map: make(map[int]Task),
	}
}

func (v *TodoList) Add(par1 Task) { //Проверку на ошибку и валидацию в Task будем в маршрутизаторе
	v.Daily_list = append(v.Daily_list, par1)
	v.Daily_map[v.Len] = par1
	v.Len++
}

func (v *TodoList) List() {
	for i := 0; i < v.Len; i++ {
		fmt.Printf("Task number %v:\nTo do -> %v\n", (i + 1), v.Daily_map[i].TodoData)
		if v.Daily_map[i].TagData == None {
			fmt.Printf("Tag: None\n")
		} else {
			fmt.Printf("Tag: %v\n", v.Daily_map[i])
		}
		if v.Daily_list[i].IsDone {
			fmt.Printf("Status: Done\n")
		} else {
			fmt.Printf("Status: Not done yet\n")
		}
	}
}

func (v *TodoList) Update(par1 Task, ind int) {
	v.Daily_list[ind] = par1
	v.Daily_map[ind] = par1
}

func (v *TodoList) Delete(ind int) {
	delete(v.Daily_map, ind)
	v.Daily_list = append(v.Daily_list[:ind], v.Daily_list[ind+1:]...)
	v.Len -= 1
}
