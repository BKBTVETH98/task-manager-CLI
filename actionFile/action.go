package actionfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"task-manager/task"
	"time"

	"github.com/fatih/color"
)

func Newtask() error {
	var status string
	var scanTextTask string

	fmt.Print("Введите описание задачи: ")
	result, err := task.GetReader()
	if err != nil {
		return err
	}
	scanTextTask = strings.TrimSpace(result)

	color.Green("Доступные варианты: running, pause, done")
	fmt.Print("Введите статус задачи: ")

	result, err = task.GetReader()
	if err != nil {
		return err
	}

	status = strings.TrimSpace(result)

	t, err := task.NewTask(scanTextTask, task.StatusCode(status))
	if err != nil {
		return err
	}

	v := task.NewVault()

	v.AddTasks(*t)

	fmt.Println()
	color.Green("новая задача успешно добавлена!")

	return nil

}

func GetTask() error {
	v := task.NewVault()

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка вывода: %w", err)
	}
	color.Yellow(string(data))
	return nil
}

func FoundTaskId() error {
	v := task.NewVault()
	var status task.StatusCode

	fmt.Print("Введите id task: ")

	choice, err := task.GetReader()

	if err != nil {
		return fmt.Errorf("Ошибка считывания - %w", err)
	}

	choiceInt, err := strconv.Atoi(strings.TrimSpace(choice))

	if err != nil {
		return fmt.Errorf("Ошибка преобразования в Int - %w", err)
	}

	if v.Tasks != nil && len(v.Tasks) >= choiceInt && choiceInt > 0 {
		taskId := v.Tasks[choiceInt-1]
		fmt.Print("Введите new status task - ")
		choice, err = task.GetReader()
		if err != nil {
			return fmt.Errorf("Ошибка считывания - %w", err)
		}
		fmt.Println(choice)
		status = task.StatusCode(choice)[:len(choice)-2]
		if err := status.Validate(); err != nil {
			return err
		}
		taskId.Status = status
		taskId.UpdatedAt = time.Now().Format(time.DateTime)
		v.Tasks[choiceInt-1] = taskId
		data, err := v.ToBytes()
		v.UpdateAt = time.Now().Format(time.DateTime)
		if err != nil {
			color.Red("не удалось перезаписать")
		}
		task.WriteJson(data)
		color.Green("статус успешно изменен")
		return nil
	}
	return errors.New("задача с таким id не найдена всего задач: " + strconv.Itoa(len(v.Tasks)))

}

func ViewTaskId() error {
	v := task.NewVault()

	fmt.Print("Введите id task ")

	result, err := task.GetReader()

	if err != nil {
		color.Red("ошибка чтения строки: %v", err)
		return nil
	}

	choice, err := strconv.Atoi(result[:len(result)-2])
	if err != nil {
		color.Red("парсинга числа: %v", err)
		return nil
	}

	if v.Tasks != nil && len(v.Tasks) >= choice && choice > 0 {
		taskId := v.Tasks[choice-1]
		fmt.Printf("Id: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
			taskId.Id,
			taskId.Description,
			taskId.Status,
			taskId.CreatedAt,
			taskId.UpdatedAt)

		return nil
	}
	_, err = task.ReadJson()
	if err != nil {
		return fmt.Errorf("задача с таким id не найдена, всего задач: %w "+strconv.Itoa(len(v.Tasks)), err)
	}
	return errors.New("задача с таким id не найдена, всего задач: " + strconv.Itoa(len(v.Tasks)))
}

func DelTaskId() error {
	v := task.NewVault()

	fmt.Println()
	fmt.Print("Введите ID задачи которую нужно удалить: ")

	text, err := task.GetReader()
	if err != nil {
		return fmt.Errorf("Не удалось считать: %w", err)
	}

	id, err := strconv.Atoi(text[:len(text)-2])
	if err != nil {
		color.Red("парсинга числа: %v", err)
		return nil
	}

	err = v.DeleteTask(id)
	if err != nil {
		return fmt.Errorf("ошибка удаления %w", err)
	}
	s, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		return fmt.Errorf("ошибка удаления %w", err)
	}
	color.Green(string(s))
	return nil

}
