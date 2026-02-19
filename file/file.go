package file

import (
	"encoding/json"
	"os"
	"task-manager/task"

	"github.com/fatih/color"
)

func ReadJson(t *task.Task) error {

	file, err := os.ReadFile("task.json")
	if err != nil {
		color.Red("ошибка чтения", err)
		return err
	}
	err = json.Unmarshal(file, &t)
	if err != nil {
		color.Red("ошибка преобразования из json", err)
		return err
	}
	return nil
}

func WriteJson(t *task.Task) error {
	file, err := os.Create("task.json")
	if err != nil {
		color.Red("нет файла")
		return err
	}
	defer file.Close()
	text, err := json.Marshal(t)
	if err != nil {
		color.Red("ошибка преобразования")
		return err
	}
	writeByte, err := file.Write(text)
	if err != nil {
		color.Red("ошибка записи", writeByte, err)
		return err
	}
	return nil
}

func EditJson() {

}
