package todo

import (
	"fmt"
)

type Tags string //Перечисление для тегов
const (
	None      Tags = ""
	Study     Tags = "study"
	Housework Tags = "housework"
	Finance   Tags = "finance"
	Health    Tags = "health"
	Work      Tags = "workx	"
)

type Task struct {
	Todo_data string //Самма информацию про задачу - что надо сдедать
	Tag_data  Tags   ////Тег задачи - по умлочанию none - не забудь пропистаь в конструкторе (идея напистаь один конструктор через функциональные параметры)
	Is_done   bool   //Сделана ли задача - поле статуса ??? - надо еще подумать
}

type Todo_list struct {
	Len        int
	Daily_list []Task
	Daily_map  map[int]Task //Не забыть в main явно прописать её при создании пустого списка через : чтоб не было ошибки работы с nil мап

}

func NewTodo_list() *Todo_list { //Чтобы не было ошибки записи в nil map
	return &Todo_list{
		Daily_map: make(map[int]Task),
	}
}

func (v *Todo_list) Add(par1 Task) { //Проверку на ошибку и валидацию в Task будем в маршрутизаторе
	v.Daily_list = append(v.Daily_list, par1)
	v.Daily_map[v.Len] = par1
	v.Len++

}

func (v *Todo_list) List() {
	for i := 0; i < v.Len; i++ {
		fmt.Printf("Task number %v:\nTo do -> %v\n", (i + 1), v.Daily_map[i].Todo_data)
		if v.Daily_map[i].Tag_data == None {
			fmt.Printf("Tag: None\n")
		} else {
			fmt.Printf("Tag: %v\n", v.Daily_map[i])
		}
		if v.Daily_list[i].Is_done {
			fmt.Printf("Status: Done\n")
		} else {
			fmt.Printf("Status: Not done yet\n")

		}
	}
}

func (v *Todo_list) Update(par1 Task, ind int) {
	v.Daily_list[ind] = par1
	v.Daily_map[ind] = par1

}

func (v *Todo_list) Delete(ind int) {
	delete(v.Daily_map, ind)
	v.Daily_list = append(v.Daily_list[:ind], v.Daily_list[ind+1:]...)
	v.Len -= 1
}

/*func (v *Todo_list) filtr_by_tag(par1 Tags) []Task { //Фильтр по тегу -> надо додумать логику - может чтоб на экран печатало
	var res []Task
	for _, el := range v.Daily_list {
		if el.Tag_data == par1 {
			res = append(res, el)
		}
	}
	return res
}

func (v *Todo_list) filtr_by_status(par1 bool) []Task { //Фильтр по статусу задачи -> додумать логику - может чтоб на жкран печатало
	var res []Task
	for _, el := range v.Daily_list {
		if el.Is_done == par1 {
			res = append(res, el)
		}
	}
	return res
}*/
