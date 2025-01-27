package data

type Task struct {
	ID int
	Title string
	Completed bool
}

var Tasks = []Task{}
var nextID = 1

func AddTask(title string) {
	Tasks = append(Tasks, Task{ID: nextID, Title: title, Completed: false})
	nextID++
}

func ToggleTask(id int) {
	for i, task := range Tasks {
		if task.ID == id {
			Tasks[i].Completed =  !Tasks[i].Completed
			break
		}
	}
}

func DeleteTask(id int) {
	for i, taks := range Tasks {
		if taks.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
		}
	}
}