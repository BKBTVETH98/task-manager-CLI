package main

import (
	"fmt"
	"strconv"
	"strings"
	actionfile "task-manager/actionFile"
	task "task-manager/task"

	color "github.com/fatih/color"
)

func main() {

	m := map[int]func(){
		1: func() {
			if err := actionfile.Newtask(); err != nil {
				color.Red(err.Error())
			}
		},
		2: func() {
			if err := actionfile.GetTask(); err != nil {
				color.Red(err.Error())
			}
		},
		3: func() {
			if err := actionfile.FoundTaskId(); err != nil {
				color.Red(err.Error())
			}
		},
		4: func() {
			if err := actionfile.ViewTaskId(); err != nil {
				color.Red(err.Error())
			}
		},
	}

	for {
		fmt.Println()
		color.Yellow("выберите действие, введя цифру от 1 до %d", len(m))
		color.Green("1, Создать таску")
		color.Green("2, вывести все таски")
		color.Green("3, изменить статус таски по Id")
		color.Green("4, вывести таску по Id \n")
		fmt.Print("Ваш выбор: ")

		choice, err := task.GetReader() //чтение sdtin

		if err != nil {
			color.Red(err.Error())
			continue
		}

		choiceInt, err := strconv.Atoi(strings.TrimSpace(choice))

		if err != nil {
			color.Red("ошибка преобразования в INT: ", err)
			continue
		}
		if j, ok := m[choiceInt]; ok {
			j()
		}
	}

}
