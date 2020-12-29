package corgi

import "github.com/wailsapp/wails"

var rt *wails.Runtime

type TaskStatus int
const (
	ReadyToStart TaskStatus = iota
	InProgress
	Paused
	Completed
	Blocked
)

func Setup(runtime *wails.Runtime) {
	rt = runtime
}

func Start() {
	go InitializeClient()
}

type Project struct {
	ID int
}

type Task struct {
	ID int
	ProjectID int
	Name string
	Description string
	Status TaskStatus
	WorkFolder string
	SubTasks []Task
	WorkLog []WorkEntry
}

type WorkEntry struct {
	Description string
	CommittedAt string
}

