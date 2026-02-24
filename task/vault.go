package task

import (
	"encoding/json"
	"fmt"
	"time"
)

type Vault struct {
	Tasks    []Task `json:"tasks"`
	UpdateAt string `json:"updateAt"`
}

func NewVault() *Vault {
	file, err := ReadJson()
	if err != nil {
		return &Vault{Tasks: []Task{}, UpdateAt: time.Now().Format(time.DateTime)}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		return &Vault{Tasks: []Task{}, UpdateAt: time.Now().Format(time.DateTime)}
	}
	return &vault

}

func (vault *Vault) AddTasks(t Task) error { //добавление
	t.Id = len(vault.Tasks) + 1
	vault.Tasks = append(vault.Tasks, t)

	data, err := vault.ToBytes() //
	vault.UpdateAt = time.Now().Format(time.DateTime)
	if err != nil {
		return fmt.Errorf("ошибка - %w ", err)
	}
	err = WriteJson(data)
	if err != nil {
		return fmt.Errorf("ошибка - %w ", err)
	}
	return nil
}

func (v *Vault) ToBytes() ([]byte, error) {
	r, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("Ошибка преобразования - %w", err)
	}
	return r, nil
}
