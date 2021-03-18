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
	"net/http"
	"io/ioutil"
)

type Timetable_Data []struct {
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

var tb Timetable_Data

func removeDup(intake_dupList []string) []string {
	intakeMap := make(map[string]bool)
	intakeList := []string{}

	for _, intake := range intake_dupList {
		if _, value := intakeMap[intake]; !value {
			intakeMap[intake] = true
			intakeList = append(intakeList, intake)
		}
	}

	return intakeList
}

func parse_JSON(link string) {
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

func intake_arrayList() []string {
	// Add all intake codes into a slice
	intake_dupList := []string{}

	for i := range tb {
		intake_dupList = append(intake_dupList, tb[i].Intake)
	}

	// Remove redundant intake codes
	return removeDup(intake_dupList)
}
