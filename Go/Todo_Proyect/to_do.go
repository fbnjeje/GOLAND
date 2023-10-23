package main

import (
	"bufio"
	"fmt"
	"os"
)

type task struct {
	name        string
	description string
	completed   bool
}

func main() {

	tasks := make(map[int]task)

	lastId := 0

	for {
		fmt.Println("Selecciona  la opcion deseada:")
		fmt.Println("1. ver tareas")
		fmt.Println("2. agregar tareas")
		fmt.Println("3. Marcar tarea como completada")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Salir")

		fmt.Print("opcion: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
	}
}
