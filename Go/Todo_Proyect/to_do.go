package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		fmt.Print("Opcion: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			if len(tasks) == 0 {
				fmt.Println("no hay nada xd")
			} else {
				for id, task := range tasks {
					fmt.Printf("%d. %s. %s (Completdado: %t)\n", id, task.name, task.description, task.completed)
				}
			}
		case 2:
		case 3:
		case 4:
		case 5:
			fmt.Print("Saliendo ...")
			os.Exit(0)
		}
	}
}
