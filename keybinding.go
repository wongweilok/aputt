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
	"time"

	"github.com/gdamore/tcell/v2"
)

// SetKeybinding sets keybindings and its function
func (w *Widget) SetKeybinding() {
	w.pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			w.app.Stop()
		case 't':
			w.pages.SwitchToPage("Timetable")
			w.removePage("Temp")
			return nil
		case 'b':
			w.pages.SwitchToPage("Browse")
			w.removePage("Temp")
			return nil
		case '/':
			w.search.SetText("")
			w.app.SetFocus(w.search)
			return nil
		}

		return event
	})

	w.search.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			w.app.SetFocus(w.pages)
			w.search.SetText("")
			return nil
		} else if event.Key() == tcell.KeyEnter {
			w.removePage("Temp")
			pageName, page := w.Temp(w.search.GetText())
			w.pages.AddAndSwitchToPage(pageName, page, true)
			w.app.SetFocus(w.pages)
			w.search.SetText("")
		}

		return event
	})

	w.timetable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 's' {
			// Create config directory if not exist
			if !checkConfigDir() {
				createConfigDir()
			}

			// Set intake code into config file and display message
			if !checkConfig() {
				writeConfig(intakeCode)
				w.search.SetText("Current intake code has been set as default.")
				go w.clearText()
			} else if readConfig() != intakeCode {
				writeConfig(intakeCode)
				w.search.SetText("Current intake code has been set as default.")
				go w.clearText()
			} else {
				w.search.SetText("Current intake code is already the default.")
				go w.clearText()
			}

			return nil
		}
		return event
	})

}

// removePage checks and removes given page
func (w *Widget) removePage(pageName string) {
	if w.pages.HasPage(pageName) {
		w.pages.RemovePage(pageName)
	}
}

// clearText clears displayed messages
func (w *Widget) clearText() {
	// Clear message after 3 seconds
	time.Sleep(3 * time.Second)
	w.app.QueueUpdateDraw(func() {
		if len(w.search.GetText()) >= 24 {
			w.search.SetText("")
		}
	})
}
