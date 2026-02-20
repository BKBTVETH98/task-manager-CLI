package task

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadJson(t *Task) error { // нужно сделать чтение массива json
	v := NewVault()
	fmt.Println(v)
	return nil
}

func WriteJson(t []byte) error {

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
