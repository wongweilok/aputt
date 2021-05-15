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
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// LoadBrowse loads browse menu with intake codes and init browse menu specific settings
func (w *Widget) LoadBrowse() (string, tview.Primitive) {
	intakes := intakeArrayList()

	writer.Init(w.timetable, 5, 0, 2, ' ', 0)

	// Display list of intake codes with table
	w.browse.SetSelectable(true, false)

	for row, i := range intakes {
		tableCell := tview.NewTableCell(i)
		tableCell.SetTextColor(tcell.ColorWhite)

		w.browse.SetCell(row, 0, tableCell)
	}

	// Display timetable of the selected intake code
	w.browse.SetSelectedFunc(func(row, column int) {
		w.pages.SwitchToPage("Timetable")

		intakeCode = intakes[row]

		count := 0
		w.timetable.SetText(intakes[row] + "\n\n")
		tb = rmDupSchedule(tb)
		for i := range tb {
			if intakes[row] == tb[i].Intake && weekNo == weekOf(tb[i].DateISO) {
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
	})

	return "Browse", w.browse
}

// Temp is a custom browse menu that loads specific intake codes based on search query
func (w *Widget) Temp(query string) (string, tview.Primitive) {
	intakes := intakeArrayList()
	shortList := []string{}

	writer.Init(w.timetable, 5, 0, 2, ' ', 0)

	w.customBrowse.SetSelectable(true, false)

	// Filter the intake code list with search keyword
	for _, i := range intakes {
		if strings.Contains(i, query) {
			shortList = append(shortList, i)
		}
	}

	// Display the filtered intake code list
	for row, i := range shortList {
		tableCell := tview.NewTableCell(i)
		tableCell.SetTextColor(tcell.ColorWhite)

		w.customBrowse.SetCell(row, 0, tableCell)
	}

	// Display timetable of the selected intake code
	w.customBrowse.SetSelectedFunc(func(row, column int) {
		w.pages.SwitchToPage("Timetable")

		intakeCode = shortList[row]

		count := 0
		w.timetable.SetText(shortList[row] + "\n\n")
		tb = rmDupSchedule(tb)
		for i := range tb {
			if shortList[row] == tb[i].Intake && weekNo == weekOf(tb[i].DateISO) {
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

		// Remove this temporary page
		w.pages.RemovePage("Temp")
	})

	return "Temp", w.customBrowse
}
