package service

import "simple-basic/model"

var tasks []model.Task

type taskServiceImpl struct {
}

func NewTaskService() TaskService {
	return &taskServiceImpl{}
}

func (service *taskServiceImpl) New(task model.Task) {
	tasks = append(tasks, task)
}

func (service *taskServiceImpl) View() []model.Task {
	return tasks
}
