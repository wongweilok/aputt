package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Browse() (string, tview.Primitive) {
	intakes := intake_arrayList()

	w.Init(timetable, 5, 0, 2, ' ', 0)

	// Display list of intake codes with table
	browse := tview.NewTable().
		SetSelectable(true, false)

	for row, i := range intakes {
		tableCell := tview.NewTableCell(i).
			SetTextColor(tcell.ColorWhite)

		browse.SetCell(row, 0, tableCell)
	}

	// Display timetable of the selected intake code
	browse.SetSelectedFunc(func(row, column int) {
		pages.SwitchToPage("Timetable")

		timetable.SetText(intakes[row] + "\n\n")
		for i := range tb {
			if intakes[row] == tb[i].Intake {
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
	})

	return "Browse", browse
}
