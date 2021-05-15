package service

import "simple-basic/model"

type TaskService interface {
	New(task model.Task)
	View() []model.Task
}
