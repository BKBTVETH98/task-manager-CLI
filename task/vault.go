package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Vault struct {
	Tasks    []Task `json:"tasks"`
	UpdateAt string `json:"updateAt"`
}

func NewVault() *Vault { //чтение
	file, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Errorf("ошибка чтения:  %w", err)
		return &Vault{Tasks: []Task{}}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Errorf("не удалось преобразовать в структуру %w", err)
	}
	return &vault

}

func (vault *Vault) AddTasks(t *Task) error { //добавление
	vault.Tasks = append(vault.Tasks, *t)

	data, err := vault.ToBytes()
	vault.UpdateAt = time.Now().Format(time.DateTime)
	if err != nil {
		return fmt.Errorf("ошибка - %w ", err)
	}
	WriteJson(data)
	return nil
}

func (v *Vault) ToBytes() ([]byte, error) {
	r, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Errorf("Ошибка преобразования - %w", err)
	}
	return r, nil
}
