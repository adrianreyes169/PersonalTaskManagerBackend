package tasks

import (
	"database/sql"
	"errors"
	"time"
)

func createTaskService(db *sql.DB, task Task) (int64, error) {
	if task.Title == "" {
		return 0, errors.New("Title has no content")
	}

	if len(task.Title) > 100 {
		return 0, errors.New("Title is too long")
	}

	if task.Content == "" {
		return 0, errors.New("Content has no text")
	}

	if len(task.Content) > 1000 {
		return 0, errors.New("Too much text on task content")
	}

	now := time.Now().UTC()
	task.createdAt = now

	ID, err := createTask(db, task)

	if err != nil {
		return 0, err
	}

	return ID, nil
}

func showTaskService(db *sql.DB) ([]Task, error) {
	tasks, err := showTask(db)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func showTaskByIDService(db *sql.DB, ID int) (Task, error) {
	if ID <= 0 {
		return Task{}, errors.New("Invalid ID number")
	}

	task, err := showTaskByID(db, ID)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func deleteTaskByIDService(db *sql.DB, ID int) (string, error) {
	if ID <= 0 {
		return "", errors.New("Invalid ID number")
	}

	msg, err := deleteTaskByID(db, ID)

	if err != nil {
		return "", err
	}

	return msg, nil

}
