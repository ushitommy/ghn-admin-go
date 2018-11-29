package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
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

//struct for Job JSON data
type Job []struct {
	ID    string `json:"id"`
	Min   string `json:"min"`
	Hour  string `json:"hour"`
	Date  string `json:"date"`
	Month string `json:"month"`
	Day   string `json:"day"`
	Text  string `json:"text"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	bytes, err := ioutil.ReadFile("data/joblist.json")
	if err != nil {
		log.Fatal(err)
	}

	var job Job
	if err := json.Unmarshal(bytes, &job); err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.ParseFiles("templates/root.html.tpl"))
	if err := t.ExecuteTemplate(w, "root.html.tpl", job); err != nil {
		log.Fatal(err)
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/edit.html.tpl"))

	bytes, err := ioutil.ReadFile("data/joblist.json")
	if err != nil {
		log.Fatal(err)
	}

	var job Job
	if err := json.Unmarshal(bytes, &job); err != nil {
		log.Fatal(err)
	}

	if err := t.ExecuteTemplate(w, "edit.html.tpl", job); err != nil {
		log.Fatal(err)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/create.html.tpl"))

	bytes, err := ioutil.ReadFile("data/joblist.json")
	if err != nil {
		log.Fatal(err)
	}

	var job Job
	if err := json.Unmarshal(bytes, &job); err != nil {
		log.Fatal(err)
	}

	if err := t.ExecuteTemplate(w, "create.html.tpl", job); err != nil {
		log.Fatal(err)
	}
}
