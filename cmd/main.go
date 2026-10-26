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

	fmt.Println("=== MIS TAREAS ===")

	for _, value := range taskList {
		status := "[ ]"

		if value.Done {
			status = "[X]"
		}

		fmt.Printf("%s %d. %s \n", status, value.ID, value.Title)
	}

}
