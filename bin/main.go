package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
type Job struct {
	ID    int    `json:"id"`
	Min   string `json:"min"`
	Hour  string `json:"hour"`
	Date  string `json:"date"`
	Month string `json:"month"`
	Days  string `json:"days"`
	Text  string `json:"text"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	bytes, err := ioutil.ReadFile("data/joblist.json")
	if err != nil {
		log.Fatal(err)
	}

	var job []Job
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

	var job []Job
	if err := json.Unmarshal(bytes, &job); err != nil {
		log.Fatal(err)
	}

	if err := t.ExecuteTemplate(w, "edit.html.tpl", job); err != nil {
		log.Fatal(err)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		t := template.Must(template.ParseFiles("templates/create.html.tpl"))

		bytes, err := ioutil.ReadFile("data/joblist.json")
		if err != nil {
			log.Fatal(err)
		}

		var job []Job
		if err := json.Unmarshal(bytes, &job); err != nil {
			log.Fatal(err)
		}

		if err := t.ExecuteTemplate(w, "create.html.tpl", job); err != nil {
			log.Fatal(err)
		}
	} else {
		bytes, err := ioutil.ReadFile("data/joblist.json")
		if err != nil {
			log.Fatal(err)
		}

		var job []Job
		if err := json.Unmarshal(bytes, &job); err != nil {
			log.Fatal(err)
		}

		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		days := strings.Join(r.Form["days[]"], ",")

		i := job[len(job)-1].ID
		newid := i + 1
		newmin := r.FormValue("min")
		newhour := r.FormValue("hour")
		newdate := r.FormValue("date")
		newmonth := r.FormValue("month")
		newday := days
		newtext := r.FormValue("text")

		newjob := Job{
			ID:    newid,
			Min:   newmin,
			Hour:  newhour,
			Date:  newdate,
			Month: newmonth,
			Days:  newday,
			Text:  newtext,
		}

		job = append(job, newjob)

		newJSON, err := json.MarshalIndent(job, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(newJSON))

		if err := ioutil.WriteFile("data/joblist.json", newJSON, 0666); err != nil {
			log.Fatal(err)
		}

		t := template.Must(template.ParseFiles("templates/create.html.tpl"))
		if err := t.ExecuteTemplate(w, "create.html.tpl", job); err != nil {
			log.Fatal(err)
		}

	}
}
