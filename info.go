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

import "fmt"

// KeyInfo stores key information
type KeyInfo struct {
	key    rune
	action string
}

// LoadInfo displays hint of keys and its action
func (w *Widget) LoadInfo() {
	// Info widget settings
	w.info.SetDynamicColors(true)
	w.info.SetWrap(false)

	keyInfo := []KeyInfo{
		{'q', "Quit"},
		{'t', "Timetable"},
		{'b', "Browse"},
		{'/', "Search"},
		{'s', "Set Default"},
	}

	for i := range keyInfo {
		fmt.Fprintf(
			w.info, "%s:[blue]%s[white]  ",
			string(keyInfo[i].key),
			keyInfo[i].action,
		)
	}
}
