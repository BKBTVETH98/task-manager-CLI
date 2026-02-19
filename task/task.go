package task

import (
	"fmt"
	"time"
)

type StatusCode string

var id int = 0 // еще не реализовано пока заглушка

const (
	runningStatus StatusCode = "running"
	pauseStatus   StatusCode = "pause"
	doneStatus    StatusCode = "done"
)

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      StatusCode `json:"status"`
	CreatedAt   string     `json:"createdAt"`
	UpdatedAt   string     `json:"updatedAt"`
}

func NewTask(des string, stat StatusCode) (*Task, error) {
	if des == "" {
		return nil, fmt.Errorf("описание задачи пустое")
	}
	if err := stat.validate(); err != nil {
		return nil, err
	}

	id++
	now := time.Now().Format(time.DateTime)

	return &Task{
		Id:          id,
		Description: des,
		Status:      stat,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (s StatusCode) validate() error {
	switch s {
	case runningStatus, pauseStatus, doneStatus:
		return nil
	default:
		return fmt.Errorf("ошибка валидации статуса: %q", s)
	}
}
