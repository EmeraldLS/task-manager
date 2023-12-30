package model

import (
	"time"
)

// Every Task is an actor on it own
type Task struct {
	Pid       string
	Status    string
	CreatedAt string
	Message   *Message
}

func NewTask(status, pid string, msg *Message) *Task {

	return &Task{
		Status:    status,
		Message:   msg,
		CreatedAt: time.Now().Format("2023-12-30 22:10:05"),
		Pid:       pid,
	}
}
