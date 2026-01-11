package tasks

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func createTaskHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var t Task

		err := json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			http.Error(w, "Invaid request body", http.StatusBadRequest)
			return
		}

		id, err := createTask(db, t)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int64{"id": id})
	}
}

func showTaskHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusBadRequest)
			return
		}

		notes, err := showTask(db)

		if err != nil {
			http.Error(w, "Failed to retrieve tasks "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(&notes)
	}
}

func showTaskByIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusBadRequest)
			return
		}

		parts := strings.Split(r.URL.Path, "/")

		if len(parts) < 3 {
			http.Error(w, "ID not provided", http.StatusBadRequest)
		}

		id, err := strconv.Atoi(parts[2])

		if err != nil || id <= 0 {
			http.Error(w, "ID not valid", http.StatusBadRequest)
		}

		note, err := showTaskByID(db, id)

		if err != nil {
			http.Error(w, "Failed to retrieve task "+err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(&note)
	}
}

func deleteTaskByIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusBadRequest)
			return
		}

		parts := strings.Split(r.URL.Path, "/")

		if len(parts) < 3 {
			http.Error(w, "ID not provided", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(parts[2])

		if err != nil || id <= 0 {
			http.Error(w, "ID not valid", http.StatusBadRequest)
			return
		}

		_, err2 := deleteTaskByID(db, id)

		if err2 != nil {
			http.Error(w, "Could not delete task", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	}
}
