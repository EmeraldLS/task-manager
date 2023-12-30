package model

import (
	"sync"
)

type User struct {
	Username     string
	Role         string
	Tasks        chan *Task
	Notification chan string
	Wg           sync.WaitGroup
}

func NewUser(username, role string) *User {

	return &User{
		Username:     username,
		Role:         role,
		Tasks:        make(chan *Task),
		Notification: make(chan string),
	}
}
