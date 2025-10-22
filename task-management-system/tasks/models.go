package taskManager

import "time"

type TaskStatus string

const (
	Pending    TaskStatus = "Pending"
	InProgress TaskStatus = "InProgress"
	Completed  TaskStatus = "Completed"
)

type User struct {
	Id    string
	Name  string
	Email string
}

type Task struct {
	Id           string
	Title        string
	Description  string
	DueDate      time.Time
	Priority     int
	Status       TaskStatus
	IsAssigned   bool
	AssignedUser *User
}
