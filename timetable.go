package main

import (
	"fmt"
	"time"
	"text/tabwriter"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	intake_code string
	w = new(tabwriter.Writer)
)

func Timetable() (string, tview.Primitive) {
	w.Init(timetable, 5, 0, 2, ' ', 0)

	// Check if config file exist
	if !checkConfig() {
		timetable.SetText("Press 'b' to browse and select an intake.")
	} else {
		// Get intake code from config file
		intake_code = readConfig()
		myintake := readConfig()

		// Display timetable
		timetable.SetText(myintake + "\n\n")
		for i := range tb {
			if myintake == tb[i].Intake {
				fmt.Fprintln(
					w, tb[i].Day + "\t" +
					tb[i].Date + "\t" +
					tb[i].StartTime + "-" + tb[i].EndTime + "\t" +
					tb[i].Room + "\t" +
					tb[i].Module + "\t" +
					tb[i].LectID,
				)
			}
		}
		w.Flush()
	}

	timetable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 's' {
			// Create config directory if not exist
			if !checkConfig() {
				createConfig()
			}

			// Set intake code into config file and display message
			if readConfig() != intake_code {
				writeConfig(intake_code)
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
