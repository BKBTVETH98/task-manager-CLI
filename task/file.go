package task

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

const (
	fileName string = "task.json"
)

func ReadJson() ([]byte, error) { // нужно сделать чтение массива json
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Println("Файл не существует или ошибка доступа, файл будет создан:", err)
		os.Create(fileName)
		return nil, err
	}
	file, err := os.ReadFile(fileName)

	//создаем файл, если его нет
	return file, nil
}

func WriteJson(t []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer file.Close()

	_, err = file.Write(t)

	if err != nil {
		color.Red("ошибка записи - %w", err)
		return err
	}
	defer file.Close()
	return nil
}

func GetReader() (string, error) {
	r := bufio.NewReader(os.Stdin)
	text, err := r.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("ошибка считывания: %w", err)
	}
	return text, nil
}
