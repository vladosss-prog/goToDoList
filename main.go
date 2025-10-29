package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type ToDo struct {
	Title string
	Done  bool
	ID    string
}

type ToDoPageData struct {
	PageTitle string
	Todos     []ToDo
}

var data = ToDoPageData{
	PageTitle: "ToDo",
	Todos:     []ToDo{},
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	errHadnler(err)

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		errHadnler(err)

		id := uuid.New()

		for i := range data.Todos {
			if r.Form.Has(data.Todos[i].ID) {
				data.Todos[i].Done = true
			} else {
				data.Todos[i].Done = false
			}
		}

		taskAppend := []ToDo{{r.FormValue("task"), false, id.String()}}
		data.Todos = append(data.Todos, taskAppend...)

		tmpl.Execute(w, data)
		return
	}

	tmpl.Execute(w, data)

}

func errHadnler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handleMain)
	http.ListenAndServe(":8080", nil)

}
