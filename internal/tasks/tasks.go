package tasks

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

func NewTask(id int, title string) Task {
	return Task{
		ID:    id,
		Title: title,
		Done:  false,
	}
}

func formatTitle(title string) string {
	return "[" + title + "]"
}

func PrintListOfTask(tasks []Task) {
	for _, value := range tasks {
		status := "[ ]"

		if value.Done {
			status = "[X]"
		}

		fmt.Printf("%s %d. %s \n", status, value.ID, value.Title)
	}
}
