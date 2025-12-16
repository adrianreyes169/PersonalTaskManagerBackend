package tasks

import (
	"database/sql"
	"errors"
)

func createTask(db *sql.DB, t Task) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (title, content, createdAt) VALUES (?,?,?) ",
		t.Title,
		t.Content,
		t.createdAt)

	if err != nil {
		return 0, err
	}

	id, err2 := result.LastInsertId()

	if err2 != nil {
		return 0, err2
	}

	return id, nil
}

func showTask(db *sql.DB) ([]Task, error) {
	var tasks []Task
	rows, err := db.Query("SELECT * FROM tasks")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.createdAt)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	err2 := rows.Err()

	if err2 != nil {
		return nil, err
	}

	return tasks, nil
}

func showTaskByID(db *sql.DB, ID int) (Task, error) {
	var t Task

	row := db.QueryRow("SELECT * FROM tasks WHERE id = (?)",
		ID)

	err2 := row.Scan(&t.ID, &t.Title, &t.Content, &t.createdAt)

	if err2 != nil {
		return Task{}, err2
	}

	return t, nil
}

func deleteTaskByID(db *sql.DB, ID int) (string, error) {
	result, err := db.Exec("DELETE FROM tasks WHERE id = (?)",
		ID)

	if err != nil {
		return "", err
	}

	rowsAffected, err2 := result.RowsAffected()

	if err2 != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", errors.New("No task was found with that ID")
	}

	return "Task deleted successfully", nil
}
