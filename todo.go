package main

import (
	"database/sql"
)

type Task struct {
	ID          int
	Description string
	Done        bool
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		done BOOLEAN NOT NULL CHECK (done IN (0, 1))
	);`
	_, err := db.Exec(query)
	return err
}

// Añadir tarea
func addTask(db *sql.DB, description string) error {
	_, err := db.Exec("INSERT INTO tasks (description, done) VALUES (?, 0)", description)
	return err
}

// Listar tareas (pendientes primero, luego hechas)
func listTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, description, done FROM tasks ORDER BY done, id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		t := Task{}
		var doneInt int
		if err := rows.Scan(&t.ID, &t.Description, &doneInt); err != nil {
			return nil, err
		}
		t.Done = doneInt == 1
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// Marcar tarea como hecha
func completeTask(db *sql.DB, id int) error {
	_, err := db.Exec("UPDATE tasks SET done = 1 WHERE id = ?", id)
	return err
}

// Editar tarea
func editTask(db *sql.DB, id int, newDescription string) error {
	_, err := db.Exec("UPDATE tasks SET description = ? WHERE id = ?", newDescription, id)
	return err
}

// Borrar tarea
func deleteTask(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
