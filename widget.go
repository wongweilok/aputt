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

import "github.com/rivo/tview"

// Widget stores all widget properties and functions
type Widget struct {
	app          *tview.Application
	flex         *tview.Flex
	pages        *tview.Pages
	timetable    *tview.TextView
	info         *tview.TextView
	search       *tview.InputField
	browse       *tview.Table
	customBrowse *tview.Table
}

// Init initializes all widget properties and functions
func (w *Widget) Init() {
	// Init application
	w.app = tview.NewApplication()

	// Timetable
	w.timetable = tview.NewTextView()

	// Info
	w.info = tview.NewTextView()
	w.LoadInfo()

	// Search field
	w.search = tview.NewInputField()
	w.LoadSearch()

	// Browse menu
	w.browse = tview.NewTable()

	// Set Pages
	w.pages = tview.NewPages()
	w.SetPage()

	// Set layout
	w.flex = tview.NewFlex()
	w.SetLayout()

	// Set Keybindings
	w.SetKeybinding()
}

// SetLayout setup the flex window layout
func (w *Widget) SetLayout() {
	w.flex.SetDirection(tview.FlexRow).
		AddItem(w.pages, 0, 1, true).
		AddItem(w.info, 1, 1, false).
		AddItem(w.search, 1, 1, false)
}

// Run starts the application
func (w *Widget) Run() {
	if err := w.app.SetRoot(w.flex, true).Run(); err != nil {
		panic(err)
	}
}
