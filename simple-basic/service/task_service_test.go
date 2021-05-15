package service

import (
	"fmt"
	"simple-basic/model"
	"testing"
)

func TestTask(t *testing.T)  {
	taskService := NewTaskService()
	task := model.Task{
		ToDoItem: "Belajar Golang",
	}

	taskService.New(task)

	datas := taskService.View()

	for _, v := range datas {
		fmt.Printf("Item: %s", v.ToDoItem)
	}
}
