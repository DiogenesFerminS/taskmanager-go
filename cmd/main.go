package main

import (
	"fmt"

	"github.com/DiogenesFerminS/taskmanager-go/internal/tasks"
)

func main() {
	task1 := tasks.NewTask(1, "Terminar curso de go")
	task2 := tasks.NewTask(2, "Hacer una api rest")

	fmt.Printf("Tarea %d: %s (Completada: %v)\n", task1.ID, task1.Title, task1.Done)
	fmt.Printf("Tarea %d: %s (Completada: %v)\n", task2.ID, task2.Title, task2.Done)

}
