package main

import (
	"fmt"
	"strings"

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

		intake_code = intakes[row]

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

func Temp(query string) (string, tview.Primitive) {
	intakes := intake_arrayList()
	short_list := []string{}

	w.Init(timetable, 5, 0, 2, ' ', 0)

	custom_browse := tview.NewTable().
		SetSelectable(true, false)

	// Filter the intake code list with search keyword
	for _, i := range intakes {
		if strings.Contains(i, query) {
			short_list = append(short_list, i)
		}
	}

	// Display the custom intake code list
	for row, i := range short_list {
		tableCell := tview.NewTableCell(i).
			SetTextColor(tcell.ColorWhite)

		custom_browse.SetCell(row, 0, tableCell)
	}

	// Display timetable of the selected intake code
	custom_browse.SetSelectedFunc(func(row, column int) {
		pages.SwitchToPage("Timetable")

		intake_code = short_list[row]

		timetable.SetText(short_list[row] + "\n\n")
		for i := range tb {
			if short_list[row] == tb[i].Intake {
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

		// Remove this temporary page
		pages.RemovePage("Temp")
	})

	return "Temp", custom_browse
}
