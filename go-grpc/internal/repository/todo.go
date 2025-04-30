package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Todo struct {
	Id        uint64        `db:"id"`
	UserId    uint64        `db:"user_id"`
	Title     string        `db:"title"`
	Body      string        `db:"body"`
	Status    string        `db:"status"`
	CreatedAt uint64        `db:"created_at"`
	UpdatedAt uint64        `db:"updated_at"`
	DeletedAt sql.NullInt64 `db:"deleted_at"`
}

func CreateTodo(db *sql.DB, payload Todo) {
	log.Println("Creating todo in MySQL database...")

	stmt, err := db.Prepare("INSERT INTO todos (user_id, title, body, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(payload.UserId, payload.Title, payload.Body, payload.CreatedAt, payload.UpdatedAt)
	if err != nil {
		log.Fatalf("failed to execute statement: %v", err)
	}

	log.Println("result:", result)
}

func FindListTodo(db *sql.DB) ([]Todo, error) {
	log.Println("Fetching todo list from MySQL database...")

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatalf("failed to query todos: %v", err)
	}
	// finally
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.UserId, &todo.Title, &todo.Body, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt); err != nil {
			log.Fatalf("failed to scan todo: %v", err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("failed to iterate todos: %v", err)
	}

	return todos, nil
}

func FindOneTodo(db *sql.DB, id uint64) (Todo, error) {
	log.Println("Fetching todo from MySQL database...")

	row := db.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	var todo Todo
	if err := row.Scan(&todo.Id, &todo.UserId, &todo.Title, &todo.Body, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no todo found with id %d", id)
			return Todo{}, nil
		}
		log.Fatalf("failed to scan todo: %v", err)
	}

	return todo, nil
}

func UpdateTodo(db *sql.DB, id uint64, payload Todo) {
	log.Println("Updating todo in MySQL database...")

	stmt, err := db.Prepare("UPDATE todos SET title = ?, body = ?, status = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(payload.Title, payload.Body, payload.Status, payload.UpdatedAt, id)
	if err != nil {
		log.Fatalf("failed to execute statement: %v", err)
	}

	log.Println("result:", result)
}

func DeleteTodo(db *sql.DB, id uint64) {
	log.Println("Deleting todo from MySQL database...")

	stmt, err := db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Fatalf("failed to execute statement: %v", err)
	}

	log.Println("result:", result)
}
