package manager

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/EmeraldLS/task-manager/model"
)

type UserActorManager struct {
	mu    sync.Mutex
	Users map[string]*model.User
	wg    sync.WaitGroup
}

func NewUserActorManager() *UserActorManager {
	return &UserActorManager{
		Users: make(map[string]*model.User),
	}
}

func (uam *UserActorManager) RegisterUser(user *model.User) {
	if uam.Users[user.Username] == nil {
		uam.mu.Lock()
		defer uam.mu.Unlock()

		uam.Users[user.Username] = user
	} else {
		slog.Error("user with username already exist", "username", user.Username)
		return
	}
}

func (uam *UserActorManager) AssignTask(taskManager *TaskActorManager, from, to *model.User, pid string) {
	task, err := taskManager.GetTaskByPID(&pid)
	if err != nil {
		slog.Error("unable to find task", "err", err)
		return
	}

	uam.wg.Add(2)
	go func() {
		defer uam.wg.Done()
		to.Tasks <- task
	}()

	go func() {
		defer uam.wg.Done()
		to.Notification <- fmt.Sprintf("Hi %v a new task has been assigned to you by: %v", to.Username, from.Username)
	}()

	uam.wg.Wait()
}

/*
	The problem is that I'm waiting for msg to be recieved before task would work
	Solution:
		I need to find a way to recieve msg first before i send to the task to.Tasks channel

*/

func (uam *UserActorManager) ReceieveTask(to *model.User) {
	for {
		select {
		case task, ok := <-to.Tasks:
			if !ok {
				slog.Error("Task channel closing...")
				os.Exit(1)
			} else {
				slog.Info("New task received", "message", task.Message.Body, "pid", task.Pid, "time stamp", task.CreatedAt)

			}

		case notif, ok := <-to.Notification:
			if !ok {
				slog.Error("Notification channel closing")
			} else {
				slog.Info("New Notification", "notification", notif)
			}

			// default:
			// 	slog.Info("No task received")
		}
	}
}
