package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type Timetable []struct {
	Intake string `json:"INTAKE"`
	Module string `json:"MODID"`
	Day string `json:"DAY"`
	Location string `json:"LOCATION"`
	Room string `json:"ROOM"`
	LectID string `json:"LECTID"`
	LectName string `json:"NAME"`
	Date string `json:"DATESTAMP"`
	StartTime string `json:"TIME_FROM"`
	EndTime string `json:"TIME_TO"`
}

func removeDup(intakes []string) []string {
	intakeMap := make(map[string]bool)
	intakeList := []string{}

	for _, intake := range intakes {
		if _, value := intakeMap[intake]; !value {
			intakeMap[intake] = true
			intakeList = append(intakeList, intake)
		}
	}

	return intakeList
}

func main() {
	// Create HTTPS Get request from open web service API
	resp, _ := http.Get("https://s3-ap-southeast-1.amazonaws.com/open-ws/weektimetable")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Parse JSON into Timetable struct
	var tb Timetable
	err := json.Unmarshal([]byte(bytes), &tb)
	if err != nil {
		panic(err)
	}

	// Add all intake codes into a slice
	intakeListDup := []string{}

	for i := range tb {
		intakeListDup = append(intakeListDup, tb[i].Intake)
	}

	// Remove redundant intake codes
	intakes := removeDup(intakeListDup)

	fmt.Println(intakes)
}
