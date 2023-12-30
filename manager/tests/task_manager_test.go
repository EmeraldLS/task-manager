package tests

import (
	"testing"

	"github.com/EmeraldLS/task-manager/manager"
	"github.com/EmeraldLS/task-manager/model"
)

var taskManager = manager.NewTaskActorManager()

func TestAddtask(t *testing.T) {
	msg := model.NewMessage("Wash Plates")
	task := model.NewTask("submitted", "1234765", msg)

	if err := taskManager.AddTask(task); err != nil {
		t.Errorf("FAILED: %v", err)
	}
}

func TestGetTaskByID(t *testing.T) {
	msg := model.NewMessage("Wash Plates")
	task := model.NewTask("submitted", "1234577", msg)

	if err := taskManager.AddTask(task); err != nil {
		t.Errorf("FAILED: %v", err)
	}
	pid := "1234577"
	task, err := taskManager.GetTaskByPID(&pid)
	if err != nil {
		t.Errorf("FAILED: %v", err)
	} else {
		t.Log("PASSED: Task gotten")
	}
}
