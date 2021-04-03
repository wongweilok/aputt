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

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Window holds all widgets properties and content
type Window func() (string, tview.Primitive)

// Global widgets
var (
	app       = tview.NewApplication()
	pages     = tview.NewPages()
	timetable = tview.NewTextView()
	search    = tview.NewInputField()
)

// URL of the timetable API
const URL string = "https://s3-ap-southeast-1.amazonaws.com/open-ws/weektimetable"

func removePage(pageName string) {
	if pages.HasPage(pageName) {
		pages.RemovePage(pageName)
	}
}

func main() {
	parseJSON(URL)

	windows := []Window{
		Timetable,
		Browse,
	}

	// Initialize UI widgets
	flex := tview.NewFlex()
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(false)

	// Add pages
	for i, window := range windows {
		pageName, page := window()
		pages.AddPage(pageName, page, true, i == 0)
	}

	pages.SetBorder(true)

	search.SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.ColorBlack)

	// Display hint
	fmt.Fprintf(info, "q:[blue]%s[white]  ", "Quit")
	fmt.Fprintf(info, "t:[blue]%s[white]  ", "Timetable")
	fmt.Fprintf(info, "b:[blue]%s[white]  ", "Browse")
	fmt.Fprintf(info, "/:[blue]%s[white]  ", "Search")
	fmt.Fprintf(info, "s:[blue]%s[white]  ", "Set Default")

	// Organize widgets placement with flex layout
	flex.SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false).
		AddItem(search, 1, 1, false)

	// Set keybindings
	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			app.Stop()
		case 't':
			pages.SwitchToPage("Timetable")
			removePage("Temp")
			return nil
		case 'b':
			pages.SwitchToPage("Browse")
			removePage("Temp")
			return nil
		case '/':
			search.SetText("")
			app.SetFocus(search)
			return nil
		}

		return event
	})

	search.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			app.SetFocus(pages)
			search.SetText("")
			return nil
		} else if event.Key() == tcell.KeyEnter {
			removePage("Temp")
			pageName, page := Temp(search.GetText())
			pages.AddAndSwitchToPage(pageName, page, true)
			app.SetFocus(pages)
			search.SetText("")
		}

		return event
	})

	// Run the application
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
