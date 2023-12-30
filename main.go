package main

import (
	"log/slog"

	"github.com/EmeraldLS/task-manager/manager"
	"github.com/EmeraldLS/task-manager/model"
)

func main() {
	user1 := model.NewUser("emeraldls", "admin")
	user2 := model.NewUser("sanni", "admin")

	userManager := manager.NewUserActorManager()
	userManager.RegisterUser(user1)
	userManager.RegisterUser(user2)

	taskManager := manager.NewTaskActorManager()
	msg := model.NewMessage("Do dishes")
	task := model.NewTask("pending", "12345", msg)

	err := taskManager.AddTask(task)
	if err != nil {
		slog.Error("unable to add task", "err", err)
	}

	go userManager.ReceieveTask(user2)

	userManager.AssignTask(taskManager, user1, user2, "12345")

	slog.Info("Program completed")
}
