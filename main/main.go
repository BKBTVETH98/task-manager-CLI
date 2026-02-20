package main

import (
	task "task-manager/task"

	color "github.com/fatih/color"
)

func main() {

	// var status task.StatusCode
	// var scanTextTask string

	for {
		// fmt.Print("Введите описание задачи: ")
		// fmt.Scan(&scanTextTask)

		// color.Green("Доступные варианты: running, pause, done")
		// fmt.Print("Введите статус задачи: ")
		// fmt.Scan(&status)
		//разобраться как сделать меню
		t, err := task.NewTask("scanTextTask", "done")
		if err != nil {
			color.Red(err.Error())
		}
		v := task.NewVault()
		v.AddTasks(t)
		return
	}
}
