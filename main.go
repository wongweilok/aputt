/* See LICENSE file for copyright and license details. */

package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Window func() (string, tview.Primitive)

var app = tview.NewApplication()

func main() {
	windows := []Window{
		Timetable,
		Browse,
	}

	// Initialize UI widgets
	flex := tview.NewFlex()
	pages := tview.NewPages()
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(false)
	search := tview.NewInputField().
		SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.ColorBlack)

	// Add pages
	for i, window := range windows {
		pageName, page := window()
		pages.AddPage(pageName, page, true, i == 0)
	}

	// Display hint
	fmt.Fprintf(info, "t:[darkcyan]%s[white]  ", "Timetable")
	fmt.Fprintf(info, "b:[darkcyan]%s[white]  ", "Browse")
	fmt.Fprintf(info, "/:[darkcyan]%s[white]  ", "Search")

	// Organize widgets placement with flex layout
	flex.SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false).
		AddItem(search, 1, 1, false)

	// Set keybindings
	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 't':
			pages.SwitchToPage("Timetable")
			return nil
		case 'b':
			pages.SwitchToPage("Browse")
			return nil
		case '/':
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
		}

		return event
	})

	// Run the application
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
