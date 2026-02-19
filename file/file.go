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

func WriteJson(t *task.Task) {
	file, err := os.Create("task.json")
	if err != nil {
		color.Red("нет файла")
	}
	text, err := json.Marshal(t)
	if err != nil {
		color.Red("ошибка преобразования")
	}
	writeByte, err := file.Write(text)
	if err != nil {
		color.Red("ошибка записи", writeByte, err)
	}
}

func EditJson() {

}
