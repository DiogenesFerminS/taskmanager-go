package main

import (
	"fmt"

	"github.com/DiogenesFerminS/taskmanager-go/internal/tasks"
)

func main() {
	store := tasks.NewStore()

	store.AddItem("La luna en la gran ciudad")
	store.AddItem("Y la adrenalina corre")
	store.AddItem("Sera que estoy alegre y no tengo pretenciones")
	store.AddItem("Y vos me das lo que hace faltaaa")

	taskList := store.List()

	fmt.Println("\n === MIS TAREAS === \n")

	tasks.PrintListOfTask(taskList)

	fmt.Println("\n === Before complete === \n")

	store.Complete(1)
	store.Complete(3)

	tasks.PrintListOfTask(taskList)
}
