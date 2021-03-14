/* See LICENSE file for copyright and license details. */

package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Window func() (string, tview.Primitive)

// Global widgets
var (
	app = tview.NewApplication()
	pages = tview.NewPages()
	timetable = tview.NewTextView()
)

func removePage(pageName string) {
	if pages.HasPage(pageName) {
		pages.RemovePage(pageName)
	}
}

func main() {
	parse_JSON("https://s3-ap-southeast-1.amazonaws.com/open-ws/weektimetable")

	windows := []Window{
		Timetable,
		Browse,
	}

	// Initialize UI widgets
	flex := tview.NewFlex()
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

	pages.SetBorder(true)

	// Display hint
	fmt.Fprintf(info, "t:[darkcyan]%s[white]  ", "Timetable")
	fmt.Fprintf(info, "b:[darkcyan]%s[white]  ", "Browse")
	fmt.Fprintf(info, "/:[darkcyan]%s[white]  ", "Search")
	fmt.Fprintf(info, "s:[darkcyan]%s[white]  ", "Search")

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
			removePage("Temp")
			return nil
		case 'b':
			pages.SwitchToPage("Browse")
			removePage("Temp")
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
