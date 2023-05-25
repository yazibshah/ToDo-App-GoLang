package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Todo LIst",
		Todos: []Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Like This VIdeo", Done: false},
		},
	}

	tmpl.Execute(w, data)
}
func main() {
	tmpl = template.Must(template.ParseFiles("templates/index.go.html"))
	http.HandleFunc("/todo", todo)
	http.ListenAndServe(":8080", nil)
}
