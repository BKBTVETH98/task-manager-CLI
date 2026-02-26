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
		5: func() {
			if err := actionfile.DelTaskId(); err != nil {
				color.Red(err.Error())
			}
		},
		6: func() {
			actionfile.GetRunningTask()
		},
		7: func() {
			actionfile.GetPauseTask()
		},
		8: func() {
			actionfile.GetDoneTask()
		},
	}

	for {
		fmt.Println()
		color.Yellow("выберите действие, введя цифру от 1 до %d", len(m))
		color.Green("1, Создать таску")
		color.Green("2, вывести все таски")
		color.Green("3, изменить статус таски по Id")
		color.Green("4, вывести таску по Id")
		color.Green("5, удалить таску по Id ")
		color.Green("6, вывести все таски в работе")
		color.Green("7, вывести все таски на паузе")
		color.Green("8, вывести все выполненные таски\n")
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
