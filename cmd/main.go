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

	fmt.Println("\n === PROCESANDO === \n")

	results := store.ProcessAll()

	for _, task := range results {
		fmt.Println(task.Message)
	}
}
