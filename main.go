package main

import (
	"html/template"
	"log"
	"net/http"
)

type ToDo struct {
	Title string
	Done  bool
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

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	taskAppend := []ToDo{{r.FormValue("task"), false}}
	data.Todos = append(data.Todos, taskAppend...)

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
