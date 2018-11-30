package main

import (
	"encoding/json"
	"fmt"

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

	fmt.Println(job)
	newid := ID + 1
	newjob := Job{
		ID:  newid,
		Min: "11",
	}

	job = append(job, newjob)

	fmt.Println(job)

	newJSON, err := json.Marshal(job)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(newJSON))

	ioutil.WriteFile("job.json", newJSON, 0666)

}

type Job struct {
	ID  string `json:"id"`
	Min string `json:"min"`
}
