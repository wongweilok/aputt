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
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// LoadBrowse loads browse menu with intake codes and init browse menu specific settings
func (w *Widget) LoadBrowse() (string, tview.Primitive) {
	intakes := intakeArrayList()

	// Browse widget settings
	w.browse.SetSelectable(true, false)
	w.browse.SetBorderPadding(0, 0, 1, 0)

	// Display list of intake codes with table
	for row, i := range intakes {
		tableCell := tview.NewTableCell(i)
		tableCell.SetTextColor(tcell.ColorWhite)

		w.browse.SetCell(row, 0, tableCell)
	}

	// Display timetable of the selected intake code
	w.browse.SetSelectedFunc(func(row, column int) {
		w.pages.SwitchToPage("Timetable")

		intakeCode = intakes[row]

		w.DisplaySchedule(intakeCode)
	})

	return "Browse", w.browse
}

// Temp is a custom browse menu that loads specific intake codes based on search query
func (w *Widget) Temp(query string) (string, tview.Primitive) {
	intakes := intakeArrayList()
	shortList := []string{}

	// CustomBrowse widget settings
	w.customBrowse = tview.NewTable()
	w.customBrowse.SetSelectable(true, false)
	w.customBrowse.SetBorderPadding(0, 0, 1, 0)

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

		w.DisplaySchedule(intakeCode)

		// Remove this temporary page
		w.pages.RemovePage("Temp")
	})

	return "Temp", w.customBrowse
}
