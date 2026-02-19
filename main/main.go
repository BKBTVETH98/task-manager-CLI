package main

import (
	"fmt"
	"task-manager/file"
	"task-manager/task"

	color "github.com/fatih/color"
)

func main() {

	var status task.StatusCode
	var scanTextTask string

	for {
		fmt.Print("Введите описание задачи: ")
		fmt.Scan(&scanTextTask)

		color.Green("Доступные варианты: running, pause, done")
		fmt.Print("Введите статус задачи: ")
		fmt.Scan(&status)

		t, err := task.NewTask(scanTextTask, status)
		if err != nil {
			color.Red(err.Error())
			continue
		}

		file.WriteJson(t)
		newtask := &task.Task{}
		err = file.ReadJson(newtask)
		fmt.Println(*newtask)

	}
}
