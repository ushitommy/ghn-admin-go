package main

import (
	"encoding/json"

	//	"html/template"
	"io/ioutil"
	"log"
	//	"net/http"
)

func main() {
	bytes, err := ioutil.ReadFile("job.json")
	if err != nil {
		log.Fatal(err)
	}

	var job []Job
	if err := json.Unmarshal(bytes, &job); err != nil {
		log.Fatal(err)
	}

	newjob := Job{
		ID:  "3",
		Min: "11",
	}

	job = append(job, newjob)

	newJSON, err := json.Marshal(job)
	if err != nil {
		log.Fatal(err)
	}

	newFile, err := ioutil.WriteFile("job.json", newJSON, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

type Job struct {
	ID  string `json:"id"`
	Min string `json:"min"`
}
