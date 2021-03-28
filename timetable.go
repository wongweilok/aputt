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

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	intakeCode string
	w          = new(tabwriter.Writer)
)

// Timetable returns its properties and content
func Timetable() (string, tview.Primitive) {
	w.Init(timetable, 5, 0, 2, ' ', 0)

	// Check if config file exist
	if !checkConfig() {
		timetable.SetText("Press 'b' to browse and select an intake.")
	} else {
		// Get intake code from config file
		intakeCode = readConfig()
		myintake := readConfig()

		// Display timetable
		timetable.SetText(myintake + "\n\n")
		for i := range tb {
			if myintake == tb[i].Intake {
				fmt.Fprintln(
					w, tb[i].Day+"\t"+
						tb[i].Date+"\t"+
						tb[i].StartTime+"-"+tb[i].EndTime+"\t"+
						tb[i].Room+"\t"+
						tb[i].Module+"\t"+
						tb[i].LectID,
				)
			}
		}
		w.Flush()
	}

	timetable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 's' {
			// Create config directory if not exist
			if !checkConfigDir() {
				createConfigDir()
			}

			// Set intake code into config file and display message
			if !checkConfig() {
				writeConfig(intakeCode)
				search.SetText("Current intake code has been set as default.")
				go clearText()
			} else if readConfig() != intakeCode {
				writeConfig(intakeCode)
				search.SetText("Current intake code has been set as default.")
				go clearText()
			} else {
				search.SetText("Current intake code is already the default.")
				go clearText()
			}

			return nil
		}
		return event
	})

	return "Timetable", timetable
}

func clearText() {
	// Clear message after 3 seconds
	time.Sleep(3 * time.Second)
	app.QueueUpdateDraw(func() {
		if len(search.GetText()) >= 24 {
			search.SetText("")
		}
	})
}
