package taskManager

import (
	"strings"
	"sync"
)

type TaskManager struct {
	Tasks     map[string]*Task
	UserTasks map[string]*Task
	mu        sync.Mutex
}

var instance *TaskManager
var once sync.Once

func GetTaskManager() *TaskManager {
	once.Do(func() {
		instance = &TaskManager{
			Tasks:     make(map[string]*Task),
			UserTasks: make(map[string]*Task),
		}
	})
	return instance
}

func (tm *TaskManager) CreateTask(task *Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.Tasks[task.GetId()] = task
}

func (tm *TaskManager) UpdateTask(updatedTask *Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if existingTask, exist := tm.Tasks[updatedTask.GetId()]; exist {
		existingTask.SetTitle(updatedTask.GetTitle())
		existingTask.SetDescription(updatedTask.GetDescription())
		existingTask.SetDueDate(updatedTask.GetDueDate())
		existingTask.SetPriority(updatedTask.GetPriority())
		existingTask.SetStatus(updatedTask.GetStatus())
	}
}

func (tm *TaskManager) DeleteTask(taskId string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if _, exist := tm.Tasks[taskId]; exist {
		delete(tm.Tasks, taskId)
	}
}

func (tm *TaskManager) SearchTask(keyword string) []*Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	var matchingTasks []*Task
	for _, task := range tm.Tasks {
		if strings.Contains(task.GetTitle(), keyword) || strings.Contains(task.GetDescription(), keyword) {
			matchingTasks = append(matchingTasks, task)
		}
	}
	return matchingTasks
}



// tasks => assing a particular task to someone...
// task can change the status of the task
//
