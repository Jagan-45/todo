package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:Google/45@tcp(172.17.0.2:3306)/")
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS todoapp")
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}

	_, err = db.Exec("USE todoapp")
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
        id INT AUTO_INCREMENT PRIMARY KEY,
        description TEXT,
        completed BOOLEAN
    )`)
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, description, completed FROM tasks")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tasks []Task
		for rows.Next() {
			var task Task
			err := rows.Scan(&task.ID, &task.Description, &task.Completed)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, task)
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, tasks)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		description := r.FormValue("description")

		_, err := db.Exec("INSERT INTO tasks (description, completed) VALUES (?, false)", description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// ...
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		id := r.PostForm.Get("id")
		completed := r.PostForm.Get("complete")

		if completed == "on" {
			_, err := db.Exec("UPDATE tasks SET completed = true WHERE id = ?", id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		id := r.PostForm.Get("delete") // Get the value of the delete button

		_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	// ...

	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)
}

type Task struct {
	ID          int
	Description string
	Completed   bool
}

func parseID(idStr string) int {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1
	}
	return id
}
