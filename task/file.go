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

func ReadJson() error { // нужно сделать чтение массива json
	_, err := os.Stat("task.json")
	if os.IsNotExist(err) {
		fmt.Println("Файл не существует или ошибка доступа:", err)
		return err
	}
	v := NewVault()
	fmt.Println(v)
	return nil
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
func GetReader() *bufio.Reader {
	r := bufio.NewReader(os.Stdin)
	return r
}
