package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	db, err := openDB()
	if err != nil {
		fmt.Println("Error abriendo la base de datos:", err)
		return
	}
	defer db.Close()

	err = createTable(db)
	if err != nil {
		fmt.Println("Error creando la tabla:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearScreen()
		tasks, err := listTasks(db)
		if err != nil {
			fmt.Println("Error listando tareas:", err)
			return
		}

		fmt.Println("\nTareas pendientes:")
		pendientes := 0
		hechas := 0
		for _, t := range tasks {
			if !t.Done {
				fmt.Printf("%d. %s\n", t.ID, t.Description)
				pendientes++
			}
		}
		if pendientes == 0 {
			fmt.Println("(No hay tareas pendientes)")
		}
		fmt.Println("\nTareas hechas:")
		for _, t := range tasks {
			if t.Done {
				fmt.Printf("%d. %s\n", t.ID, t.Description)
				hechas++
			}
		}
		if hechas == 0 {
			fmt.Println("(No hay tareas hechas)")
		}

		fmt.Println("\n¿Qué quieres hacer?")
		fmt.Println("1. Añadir tarea")
		fmt.Println("2. Terminar tarea")
		fmt.Println("3. Editar tarea")
		fmt.Println("4. Borrar tarea")
		fmt.Println("5. Salir")
		fmt.Print("> ")
		scanner.Scan()
		opcion := strings.TrimSpace(scanner.Text())

		switch opcion {
		case "1":
			fmt.Print("Descripción de la nueva tarea: ")
			scanner.Scan()
			desc := scanner.Text()
			if err := addTask(db, desc); err != nil {
				fmt.Println("Error añadiendo tarea:", err)
			}
		case "2":
			fmt.Print("ID de la tarea a marcar como hecha: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			if err := completeTask(db, id); err != nil {
				fmt.Println("Error marcando tarea como hecha:", err)
			}
		case "3":
			fmt.Print("ID de la tarea a editar: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Nueva descripción: ")
			scanner.Scan()
			desc := scanner.Text()
			if err := editTask(db, id, desc); err != nil {
				fmt.Println("Error editando tarea:", err)
			}
		case "4":
			fmt.Print("ID de la tarea a borrar: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			if err := deleteTask(db, id); err != nil {
				fmt.Println("Error borrando tarea:", err)
			}
		case "5":
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}
