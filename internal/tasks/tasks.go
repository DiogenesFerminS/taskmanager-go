package tasks

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
