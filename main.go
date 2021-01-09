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

func main() {
	resp, _ := http.Get("https://s3-ap-southeast-1.amazonaws.com/open-ws/weektimetable")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var tb Timetable
	err := json.Unmarshal([]byte(bytes), &tb)
	if err != nil {
		panic(err)
	}

	fmt.Println(tb)
}
