package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type task struct {
	id          int
	name        string
	description string
}

func main() {
	tasks := []task{}
	for {
		ptr := &tasks
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Привет! Это треккер твоих задач!\n1. Посмотреть задачи\n2. Редактировать задачу\n3. Добавить задачу")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода")
			break
		}

		text := scanner.Text()
		fields := strings.Fields(text)
		if len(fields) > 1 {
			fmt.Println("Необходимо выбрать одно действие")
		} else if len(fields) < 1 {
			fmt.Println("Введите хотя бы одно действие")
			continue
		}

		if fields[0] == "1" {
			fmt.Println(tasks)
		} else if fields[0] == "3" {
			addTask(ptr)
		} else if fields[0] == "2" {
			fmt.Println("Ввдетие айди таска")
			if ok := scanner.Scan(); !ok {
				fmt.Println("Error")
				break
			}
			id := scanner.Text()
			idToInt, _ := strconv.Atoi(id)

			for _, t := range tasks {
				if t.id == idToInt {
					fmt.Println("Found ID", t.id, t)
					ptr := &t
					updateTask(ptr)
				}
			}
		}
	}
}

func addTask(t *[]task) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Write name: ")
	if ok := scanner.Scan(); !ok {
		fmt.Println("Error")
		return
	}
	name := scanner.Text()
	fmt.Print("Write description: ")
	if ok := scanner.Scan(); !ok {
		fmt.Println("erorr")
		return
	}
	desc := scanner.Text()

	maxID := 0
	for _, t := range *t {
		if t.id > maxID {
			maxID = t.id
		}
		maxID++
	}

	newTask := task{
		id:          maxID,
		name:        name,
		description: desc,
	}

	*t = append(*t, newTask)
}

func updateTask(t *task) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Любое поле которое осталось пустым - оставить без изменений")
	fmt.Print("Введите название: ")
	if ok := scanner.Scan(); !ok {
		fmt.Println("error")
		return
	}
	name := scanner.Text()
	//if name == "" {
	//name = t.name
	//}
	fmt.Println(name)

	fmt.Print("Введите описание: ")
	if ok := scanner.Scan(); !ok {
		fmt.Println("error")
		return
	}
	desc := scanner.Text()
	fmt.Println(desc)
	//if desc == "" {
	//	desc = t.description
	//}

	*t = task{
		id:          t.id,
		name:        name,
		description: desc,
	}
	fmt.Println(*t)
}
