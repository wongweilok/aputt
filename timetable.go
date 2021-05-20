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
	"fmt"
	"text/tabwriter"
	"time"

	"github.com/rivo/tview"
)

var (
	dIntakeCode string // Default intake code
	intakeCode  string // Currently view intake code
	writer      = new(tabwriter.Writer)
	_, weekNo   = time.Now().ISOWeek() // Get week number of current time
)

// LoadTimetable displays timetable schedule
func (w *Widget) LoadTimetable() (string, tview.Primitive) {
	// Timetable widget settings
	w.timetable.SetBorderPadding(0, 0, 1, 0)

	// Check if config file exist
	if !checkConfig() {
		w.timetable.SetText("Press 'b' to browse and select an intake.")
	} else {
		// Get intake code from config file
		dIntakeCode = readConfig()
		intakeCode = dIntakeCode

		// Display timetable schedule
		w.DisplaySchedule(intakeCode)
	}

	return "Timetable", w.timetable
}

// Get week number of a date
func weekOf(dateISO string) int {
	date, _ := time.Parse(time.RFC3339, dateISO)
	_, weekNo := date.ISOWeek()

	return weekNo
}

// Remove duplicate timetable schedule
func rmDupSchedule(tb []TimetableData) []TimetableData {
	type key struct {
		Intake    string
		Module    string
		Day       string
		Room      string
		LectID    string
		Date      string
		DateISO   string
		StartTime string
		EndTime   string
		Group     string
	}

	var tbUnique []TimetableData
	tbMap := make(map[key]int)

	for _, slot := range tb {
		k := key{
			slot.Intake,
			slot.Module,
			slot.Day,
			slot.Room,
			slot.LectID,
			slot.Date,
			slot.DateISO,
			slot.StartTime,
			slot.EndTime,
			"",
		}

		if i, ok := tbMap[k]; ok {
			// Replace group number to 'All' for repetitive schedule
			k = key{
				slot.Intake,
				slot.Module,
				slot.Day,
				slot.Room,
				slot.LectID,
				slot.Date,
				slot.DateISO,
				slot.StartTime,
				slot.EndTime,
				"All",
			}

			tbUnique[i] = TimetableData(k)
		} else {
			tbMap[k] = len(tbUnique)
			tbUnique = append(tbUnique, slot)
		}
	}

	return tbUnique
}

func (w *Widget) DisplaySchedule(intakeCode string) {
	// Initialize tabwriter
	writer.Init(w.timetable, 5, 0, 2, ' ', 0)

	count := 0
	w.timetable.SetText(intakeCode + "\n\n")
	tb = rmDupSchedule(tb)
	for i := range tb {
		if intakeCode == tb[i].Intake && weekNo == weekOf(tb[i].DateISO) {
			count++
			fmt.Fprintln(
				writer, tb[i].Day+"\t"+
					tb[i].Date+"\t"+
					tb[i].StartTime+"-"+tb[i].EndTime+"\t"+
					tb[i].Room+"\t"+
					tb[i].Module+"\t"+
					tb[i].LectID+"\t"+
					tb[i].Group,
			)
		}
	}
	if count == 0 {
		fmt.Fprintln(writer, "No classes for this week.")
	}
	writer.Flush()
}
