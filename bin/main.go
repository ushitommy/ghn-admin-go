package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

	if r.Method == "GET" {

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

		i, _ := strconv.Atoi(r.FormValue("id"))
		newid := i - 1
		newmin := r.FormValue("min")
		newhour := r.FormValue("hour")
		newdate := r.FormValue("date")
		newmonth := r.FormValue("month")
		newday := strings.Join(r.Form["days[]"], ",")
		newtext := r.FormValue("text")
		job[newid].Min = newmin
		job[newid].Hour = newhour
		job[newid].Date = newdate
		job[newid].Month = newmonth
		job[newid].Days = newday
		job[newid].Text = newtext

		fmt.Println(job[newid])

		newJSON, err := json.MarshalIndent(job, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile("data/joblist.json", newJSON, 0666); err != nil {
			log.Fatal(err)
		}

		if err := t.ExecuteTemplate(w, "edit.html.tpl", job); err != nil {
			log.Fatal(err)
		}
	}
}

//新規にジョブを追加する
func createHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/create.html.tpl"))
	//GETのとき
	if r.Method == "GET" {
		//JSONの読み込み
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
		//POSTのとき
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

		i := job[len(job)-1].ID
		newid := i + 1
		newmin := r.FormValue("min")
		newhour := r.FormValue("hour")
		newdate := r.FormValue("date")
		newmonth := r.FormValue("month")
		newday := strings.Join(r.Form["days[]"], ",")
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

		if err := ioutil.WriteFile("data/joblist.json", newJSON, 0666); err != nil {
			log.Fatal(err)
		}

		if err := t.ExecuteTemplate(w, "create.html.tpl", job); err != nil {
			log.Fatal(err)
		}
	}
}
