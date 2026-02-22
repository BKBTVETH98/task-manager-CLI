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
				fmt.Println(err.Error())
			}
		},
		4: func() {
			if err := viewTaskId(); err != nil {
				color.Red(err.Error())
			}
		},
	}

	for {
		color.Yellow("выберите действие, введя цифру от 1 до %d", len(m)) //пофиксить двойную печать
		color.Green("1, Создать таску")
		color.Green("2, вывести все таски")
		color.Green("3, изменить статус таски по Id")
		color.Green("4, вывести таску по Id")
		fmt.Print("Ваш выбор: ")

		choice, err := task.GetReader().ReadString('\n')
		if err != nil {
			color.Red("ошибка чтения строки: ", err)
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

	testInput, err = reader.ReadString('\n')
	if err != nil {
		color.Red("ошибка чтения строки: %v", err)
		return
	}

	status = strings.TrimSpace(testInput)

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
	var status task.StatusCode

	fmt.Print("Введите id task: ")

	choice, err := task.GetReader().ReadString('\n')

	if err != nil {
		return fmt.Errorf("Ошибка считывания - %w", err)
	}

	choiceInt, err := strconv.Atoi(strings.TrimSpace(choice))

	if err != nil {
		return fmt.Errorf("Ошибка преобразования в Int - %w", err)
	}

	if vault.Tasks != nil && len(vault.Tasks) >= choiceInt && choiceInt > 0 {
		taskId := vault.Tasks[choiceInt-1]
		fmt.Print("Введите new status task")
		fmt.Scan(&status)
		if err := status.Validate(); err != nil {
			return err
		}
		taskId.Status = status
		taskId.UpdatedAt = time.Now().Format(time.DateTime)
		vault.Tasks[choiceInt-1] = taskId
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
	testInput, err := task.GetReader().ReadString('\n') //дописать
	if err != nil {
		color.Red("ошибка чтения строки: %v", err)
		return nil
	}
	shoice, err := strconv.Atoi(testInput)
	if err != nil {
		color.Red("парсинга числа: %v", err)
		return nil
	}
	if len(vault.Tasks) >= shoice && searchId > 0 && task.ReadJson() == nil {
		taskId := vault.Tasks[searchId-1]
		fmt.Printf("Id: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", taskId.Id, taskId.Description, taskId.Status, taskId.CreatedAt, taskId.UpdatedAt)
		return nil
	}
	return errors.New("задача с таким id не найдена, всего задач: " + fmt.Sprint(vault.Tasks[len(vault.Tasks)-1].Id))
}
