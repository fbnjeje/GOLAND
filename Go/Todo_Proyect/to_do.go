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
			fmt.Print("Nombre de la tarea: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Println("ingrese descripcion: ")
			scanner.Scan()
			description := scanner.Text()

			lastId++

			tasks[lastId] = task{name: name, description: description, completed: false}
			fmt.Printf("su tarea %d ha sido agregada", lastId)
		case 3:
			fmt.Println("ingrese el id de la tarea completada")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			if task, ok := tasks[id]; ok {
				task.completed = true
				tasks[id] = task

				fmt.Print("tarea  completada")
			} else {
				fmt.Println("no existe")
			}

		case 4:

			fmt.Println("ingrese el id de la tarea completada")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			if _, ok := tasks[id]; ok {
				delete(tasks, id)
				fmt.Printf("tarea con id %d ha sido eliminada", id)
			} else {
				fmt.Println("no existe")
			}

		case 5:
			fmt.Print("Saliendo ...")
			os.Exit(0)
		}
	}
}
