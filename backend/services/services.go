package services

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Aaryan@07@tcp(127.0.0.1:3306)/todoapp")
	if err != nil {
		log.Fatal(err)
	}
}

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func GetTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func CreateTodoWithTitle(title string) error {
	_, err := db.Exec("INSERT INTO todos (title) VALUES (?)", title)
	return err
}

func UpdateTodoTitleByID(id int, title string) error {
	_, err := db.Exec("UPDATE todos SET title = ? WHERE id = ?", title, id)
	return err
}

func DeleteTodoByID(id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
