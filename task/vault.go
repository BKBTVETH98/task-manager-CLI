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
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Errorf("ошибка чтения:  %w", err)
		return &Vault{Tasks: []Task{}, UpdateAt: time.Now().Format(time.DateTime)}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Errorf("не удалось преобразовать в структуру %w", err)
		return &Vault{Tasks: []Task{}, UpdateAt: time.Now().Format(time.DateTime)}
	}
	return &vault

}

func (vault *Vault) AddTasks(t Task) error { //добавление
	t.Id = len(vault.Tasks) + 1
	vault.Tasks = append(vault.Tasks, t)

	data, err := vault.ToBytes()
	vault.UpdateAt = time.Now().Format(time.DateTime)
	if err != nil {
		return fmt.Errorf("ошибка - %w ", err)
	}
	WriteJson(data)
	return nil
}

func (v *Vault) ToBytes() ([]byte, error) {
	r, err := json.Marshal(v)
	if err != nil {
		fmt.Errorf("Ошибка преобразования - %w", err)
	}
	return r, nil
}
