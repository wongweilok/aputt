/*
   Copyright (C) 2021 Wong Wei Lok <wongweilok@disroot.org>

   This file is part of aputt.

   aputt is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   aputt is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with aputt.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// TimetableData holds all timetable information
type TimetableData []struct {
	Intake    string `json:"INTAKE"`
	Module    string `json:"MODID"`
	Day       string `json:"DAY"`
	Room      string `json:"ROOM"`
	LectID    string `json:"LECTID"`
	Date      string `json:"DATESTAMP"`
	DateISO   string `json:"TIME_FROM_ISO"`
	StartTime string `json:"TIME_FROM"`
	EndTime   string `json:"TIME_TO"`
}

var tb TimetableData

func removeDup(intakeDupList []string) []string {
	intakeMap := make(map[string]bool)
	intakeList := []string{}

	for _, intake := range intakeDupList {
		if _, value := intakeMap[intake]; !value {
			intakeMap[intake] = true
			intakeList = append(intakeList, intake)
		}
	}

	return intakeList
}

func parseJSON(link string) {
	// Create HTTPS Get request from open web service API
	resp, _ := http.Get(link)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Parse JSON into "Timetable" struct data type
	err := json.Unmarshal([]byte(bytes), &tb)
	if err != nil {
		panic(err)
	}
}

func intakeArrayList() []string {
	// Add all intake codes into a slice
	intakeDupList := []string{}

	for i := range tb {
		intakeDupList = append(intakeDupList, tb[i].Intake)
	}

	// Remove redundant intake codes
	return removeDup(intakeDupList)
}
