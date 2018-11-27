package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/create", createHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../templates/root.html.tpl"))
	title := "Top Page"
	if err := t.ExecuteTemplate(w, "root.html.tpl", title); err != nil {
		log.Fatal(err)
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../templates/edit.html.tpl"))
	title := "Edit Page"
	if err := t.ExecuteTemplate(w, "edit.html.tpl", title); err != nil {
		log.Fatal(err)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../templates/create.html.tpl"))
	title := "Create Page"
	if err := t.ExecuteTemplate(w, "create.html.tpl", title); err != nil {
		log.Fatal(err)
	}
}
