package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type TaskStatus string

const (
	StatusToDo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	ID        int
	Name      string
	Status    TaskStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	db, err := sql.Open("sqlite3", "./tasks.db?_busy_timeout=500&_journal_mode=WAL&_sync=NORMAL")
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Microsecond * 500)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	defer db.Close()

	if err = setupDatabase(db); err != nil {
		log.Fatal(err)
	}

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Введите хотя бы один аргумент")
	}

	switch args[1] {
	case "add":
		if len(args) < 3 {
			log.Fatal("Введите название таски")
		}
		name := args[2]
		insertTask(db, name, StatusInProgress)
	case "update":
		if len(args) < 4 {
			log.Fatal("Введите id и название таски")
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}
		newName := args[3]
		updateTask(db, id, newName, StatusInProgress)
	case "delete":
		if len(args) < 3 {
			log.Fatal("Введите id таски")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}
		deleteTask(db, id)
	case "mark-in-progress":
		if len(args) < 3 {
			log.Fatal("Введите id такси")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}
		name, err := getTaskNameById(db, id)
		if err != nil {
			log.Fatal(err)
		}
		updateTask(db, id, name, StatusInProgress)
	case "mark-done":
		if len(args) < 3 {
			log.Fatal("Введите id такси")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}
		name, err := getTaskNameById(db, id)
		if err != nil {
			log.Fatal(err)
		}
		updateTask(db, id, name, StatusDone)
	case "mark-todo":
		if len(args) < 3 {
			log.Fatal("Введите id такси")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}
		name, err := getTaskNameById(db, id)
		if err != nil {
			log.Fatal(err)
		}
		updateTask(db, id, name, StatusToDo)

	case "list":
		tasks, err := getAllTasks(db)
		if err != nil {
			log.Fatal(err)
		}

		for _, task := range tasks {
			fmt.Println("-------------")
			fmt.Println("ID:", task.ID)
			fmt.Println("Name:", task.Name)
			fmt.Println("Status:", task.Status)
			fmt.Println("Created At:", normalizeDate(task.CreatedAt))
			fmt.Println("Updated At:", normalizeDate(task.UpdatedAt))
			fmt.Println("-------------")
		}
	}
}

func normalizeDate(date time.Time) string {
	return date.Format("02.01.2006 15:02")
}

func getTaskNameById(db *sql.DB, id int) (string, error) {
	row, err := db.Query("SELECT id, name, status, created_at, updated_at FROM tasks WHERE id = ?", id)
	if err != nil {
		return "", err
	}

	defer row.Close()
	var task Task
	for row.Next() {
		err = row.Scan(&task.ID, &task.Name, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	}
	if err != nil {
		return "", err
	}

	return task.Name, nil
}

func getAllTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, name, status, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

func insertTask(db *sql.DB, name string, status TaskStatus) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (name, status) VALUES (?, ?)", name, status)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateTask(db *sql.DB, id int, name string, status TaskStatus) error {
	_, err := db.Exec("UPDATE tasks SET name = ?, status = ? WHERE id = ?", name, status, id)
	return err
}

func deleteTask(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)

	return err
}

func setupDatabase(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	status TEXT NOT NULL CHECK(status IN ('todo', 'in-progress', 'done')),
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
			CREATE TRIGGER IF NOT EXISTS update_tasks_timestamp
		AFTER UPDATE ON tasks
		FOR EACH ROW
		BEGIN 
			UPDATE tasks SET updated_at = CURRENT_TIMESTAMP
		WHERE ID = OLD.id;
		END
		`)

	return err
}
