package manager

import (
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/EmeraldLS/task-manager/model"
)

type TaskActorManager struct {
	mu       sync.RWMutex
	AllTasks map[string]*model.Task
}

func NewTaskActorManager() *TaskActorManager {
	return &TaskActorManager{
		AllTasks: make(map[string]*model.Task),
	}
}

// Process ID is the id related to a particular task
func (tam *TaskActorManager) AddTask(task *model.Task) error {

	if tam.AllTasks[task.Pid] == nil {
		slog.Info("Adding task...")
		tam.mu.Lock()
		defer tam.mu.Unlock()

		tam.AllTasks[task.Pid] = task

		slog.Info("Task added", "process id", task.Pid)
		return nil
	} else {
		return errors.New("task with pid already exist")
	}
}

func (tam *TaskActorManager) GetTaskByPID(pid *string) (*model.Task, error) {
	// tam.mu.RLock()
	// defer tam.mu.Unlock()

	task := tam.AllTasks[*pid]
	if task == nil {
		return &model.Task{}, fmt.Errorf(`task with pid="%v" not found`, *pid)
	}

	return task, nil
}

func (tam *TaskActorManager) OutputAllTasks() {
	for pid, task := range tam.AllTasks {
		slog.Info("\tPid = \"%v\" Task = \"%v\"", pid, task)

	}
}
