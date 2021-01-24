/* See LICENSE file for copyright and license details. */

package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"text/tabwriter"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Timetable []struct {
	Intake string `json:"INTAKE"`
	Module string `json:"MODID"`
	Day string `json:"DAY"`
	Location string `json:"LOCATION"`
	Room string `json:"ROOM"`
	LectID string `json:"LECTID"`
	LectName string `json:"NAME"`
	Date string `json:"DATESTAMP"`
	StartTime string `json:"TIME_FROM"`
	EndTime string `json:"TIME_TO"`
}

func removeDup(intakes []string) []string {
	intakeMap := make(map[string]bool)
	intakeList := []string{}

	for _, intake := range intakes {
		if _, value := intakeMap[intake]; !value {
			intakeMap[intake] = true
			intakeList = append(intakeList, intake)
		}
	}

	return intakeList
}

func main() {
	// Create HTTPS Get request from open web service API
	resp, _ := http.Get("https://s3-ap-southeast-1.amazonaws.com/open-ws/weektimetable")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Parse JSON into Timetable struct
	var tb Timetable
	err := json.Unmarshal([]byte(bytes), &tb)
	if err != nil {
		panic(err)
	}

	// Add all intake codes into a slice
	intakeListDup := []string{}

	for i := range tb {
		intakeListDup = append(intakeListDup, tb[i].Intake)
	}

	// Remove redundant intake codes
	intakes := removeDup(intakeListDup)

	// Initialize UI widgets
	app := tview.NewApplication()
	intakeCodes := tview.NewTable().SetSelectable(true, false)
	timetable := tview.NewTextView()
	flex := tview.NewFlex()
	searchBox := tview.NewInputField().
		SetLabel("Search: ").
		SetFieldWidth(19)

	// Display the intake codes that have timetable available
	for row, i := range intakes {
		tableCell := tview.NewTableCell(i).
			SetTextColor(tcell.ColorWhite)

		intakeCodes.SetCell(row, 0, tableCell)
	}

	w := new(tabwriter.Writer)
	w.Init(timetable, 5, 0, 2, ' ', 0)

	// Display the timetable based on the selected intake code
	intakeCodes.SetSelectedFunc(func(row, column int) {
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

	intakeCodes.SetBorder(true)
	timetable.SetBorder(true)

	// Layout widgets with Flexbox
	flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(searchBox, 1, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(intakeCodes, 0, 1, true).
			AddItem(timetable, 0, 5, false), 0, 25, true), 0, 1, true)

	// Switch focus with Tab key
	intakeCodes.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			app.SetFocus(searchBox)
		}
		return event
	})

	searchBox.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			app.SetFocus(intakeCodes)
		}
		return event
	})

	// Very basic search function
	searchBox.SetDoneFunc(func(key tcell.Key) {
		found := false
		for i := range tb {
			if searchBox.GetText() == tb[i].Intake {
				found = true

				timetable.SetText(tb[i].Intake + "\n\n")
				fmt.Fprintln(
					w, tb[i].Day + "\t" +
					tb[i].Date + "\t" +
					tb[i].StartTime + "-" + tb[i].EndTime + "\t" +
					tb[i].Room + "\t" +
					tb[i].Module + "\t" +
					tb[i].LectID,
				)

				for row, j := range intakes {
					if tb[i].Intake == j {
						intakeCodes.Select(row, 0)
					}
				}
			}
		}
		w.Flush()

		if !found {
			timetable.SetText("No search result found.")
		}
	})

	// Run the application
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
