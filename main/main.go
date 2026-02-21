package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	task "task-manager/task"
	"time"

	color "github.com/fatih/color"
)

func main() {

	m := map[int]func(){
		1: newtask,
		2: getTask,
		3: func() {
			if err := foundTask(); err != nil {
				color.Red(err.Error())
			}
		},
		4: func() {
			if err := viewTaskId(); err != nil {
				color.Red(err.Error())
			}
		},
	}

	for {
		color.Green("выберите действие, введя цифру от 1 до %d", len(m)) //пофиксить двойную печать
		color.Green("1, Создать таску")
		color.Green("2, вывести все таски")
		color.Green("3, изменить статус таски по Id")
		color.Green("4, вывести таску по Id")
		fmt.Print("Ваш выбор: ")
		choice, err := task.GetReader().ReadString('\n')
		if err != nil {
			color.Red("ошибка чтения строки: %v", err)
			continue
		}

		choiceInt, _ := strconv.Atoi(strings.TrimSpace(choice))

		if j, ok := m[choiceInt]; ok {
			j()
		}
	}

}

func newtask() {
	var status string
	var scanTextTask string

	reader := task.GetReader()

	fmt.Print("Введите описание задачи: ")
	testInput, err := reader.ReadString('\n')
	if err != nil {
		color.Red("ошибка чтения строки: %v", err)
		return
	}
	scanTextTask = strings.TrimSpace(testInput)

	color.Green("Доступные варианты: running, pause, done")
	fmt.Print("Введите статус задачи: ")

	statusInput, err := reader.ReadString('\n')
	if err != nil {
		color.Red("ошибка чтения строки: %v", err)
		return
	}
	status = strings.TrimSpace(statusInput)

	t, err := task.NewTask(scanTextTask, task.StatusCode(status))
	if err != nil {
		color.Red(err.Error())
		return
	}

	v := task.NewVault()
	v.AddTasks(*t)

}

func getTask() {
	v := task.NewVault()
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		color.Red("ошибка вывода", err)
	}
	color.Yellow(string(data))

}

func foundTask() error {
	vault := task.NewVault()
	var searchId int
	var status task.StatusCode

	fmt.Print("Введите id task ")
	fmt.Scan(&searchId)

	if len(vault.Tasks) >= searchId {
		taskId := vault.Tasks[searchId-1]
		fmt.Print("Введите new status task")
		fmt.Scan(&status)
		if err := status.Validate(); err != nil {
			return err
		}
		taskId.Status = status
		taskId.UpdatedAt = time.Now().Format(time.DateTime)
		vault.Tasks[searchId-1] = taskId
		data, err := vault.ToBytes()
		vault.UpdateAt = time.Now().Format(time.DateTime)
		if err != nil {
			color.Red("не удалось перезаписать")
		}
		task.WriteJson(data)
		color.Green("статус успешно изменен")
		return nil
	}
	return errors.New("задача с таким id не найдена")

}

func viewTaskId() error {
	vault := task.NewVault()
	var searchId int

	fmt.Print("Введите id task ")
	fmt.Scan(&searchId)

	if len(vault.Tasks) >= searchId && searchId > 0 && task.ReadJson() == nil {
		taskId := vault.Tasks[searchId-1]
		fmt.Printf("Id: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", taskId.Id, taskId.Description, taskId.Status, taskId.CreatedAt, taskId.UpdatedAt)
		return nil
	}
	return errors.New("задача с таким id не найдена, всего задач: " + fmt.Sprint(vault.Tasks[len(vault.Tasks)-1].Id))
}
