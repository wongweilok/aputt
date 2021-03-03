package main

import (
	"fmt"
	"text/tabwriter"

	"github.com/rivo/tview"
)

func Timetable() (string, tview.Primitive) {
	timetable := tview.NewTextView()

	parse_to_array("https://s3-ap-southeast-1.amazonaws.com/open-ws/weektimetable")

	w := new(tabwriter.Writer)
	w.Init(timetable, 5, 0, 2, ' ', 0)

	myintake := "UC2F2008SE"

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

	timetable.SetBorder(true)

	return "Timetable", timetable
}
