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

// Window stores specific widgets information
type Window func() (string, tview.Primitive)

// SetPage setup and loads pages
func (w *Widget) SetPage() {
	windows := []Window{
		w.LoadTimetable,
		w.LoadBrowse,
	}

	for i, window := range windows {
		pageName, page := window()
		w.pages.AddPage(pageName, page, true, i == 0)
	}

	w.pages.SetBorder(true)
}
