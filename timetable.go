package main

import (
	"fmt"
	"text/tabwriter"

	"github.com/rivo/tview"
)

var w = new(tabwriter.Writer)

func Timetable() (string, tview.Primitive) {
	w.Init(timetable, 5, 0, 2, ' ', 0)

	// Check if config file exist
	if !checkConfig() {
		timetable.SetText("Press 'b' to browse and select an intake.")
	} else {
		// Get intake code from config file
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

	return "Timetable", timetable
}
