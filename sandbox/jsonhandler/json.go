package main

import (
	"encoding/json"

	//	"html/template"
	"io/ioutil"
	"log"
	//	"net/http"
)

type Job struct {
	ID  int    `json:"id"`
	Min string `json:"min"`
}

func main() {
	//JSON読み込み
	bytes, err := ioutil.ReadFile("job.json")
	if err != nil {
		log.Fatal(err)
	}

	var job []Job
	if err := json.Unmarshal(bytes, &job); err != nil {
		log.Fatal(err)
	}

	i := job[len(job)-1].ID
	newid := i + 1
	newjob := Job{
		ID:  newid,
		Min: "11",
	}

	job = append(job, newjob)

	newJSON, err := json.Marshal(job)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("job.json", newJSON, 0666)

}
