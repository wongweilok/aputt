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
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

// URL of the timetable API
const URL string = "https://s3-ap-southeast-1.amazonaws.com/open-ws/weektimetable"

func main() {
	parseJSON(URL)

	// Init tabwriter
	tw := new(tabwriter.Writer)
	tw.Init(os.Stdout, 5, 0, 2, ' ', 0)

	// Command line flag
	intake := flag.String("i", "", "Display schedule of given intake code.")
	df := flag.Bool("d", false, "Display default intake schedule.")
	flag.Parse()

	// Display intake schedule as standard output
	if *intake != "" {
		count := 0
		tb = rmDupSchedule(tb)
		for i := range tb {
			if *intake == tb[i].Intake && weekNo == weekOf(tb[i].DateISO) {
				count++
				fmt.Fprintln(
					tw, tb[i].Day+"\t"+
						tb[i].Date+"\t"+
						tb[i].StartTime+"-"+tb[i].EndTime+"\t"+
						tb[i].Room+"\t"+
						tb[i].Module+"\t"+
						tb[i].LectID+"\t"+
						tb[i].Group,
				)
			}
		}
		if count == 0 {
			fmt.Println("No match found. / No classes for this week.")
		}

		tw.Flush()
		os.Exit(0)
	}

	if *df {
		if !checkConfig() {
			fmt.Println("No default intake code was set.")
		} else {
			intakeCode := readConfig()

			count := 0
			tb = rmDupSchedule(tb)
			for i := range tb {
				if intakeCode == tb[i].Intake && weekNo == weekOf(tb[i].DateISO) {
					count++
					fmt.Fprintln(
						tw, tb[i].Day+"\t"+
							tb[i].Date+"\t"+
							tb[i].StartTime+"-"+tb[i].EndTime+"\t"+
							tb[i].Room+"\t"+
							tb[i].Module+"\t"+
							tb[i].LectID+"\t"+
							tb[i].Group,
					)
				}
			}
			if count == 0 {
				fmt.Println("No classes for this week.")
			}

			tw.Flush()
			os.Exit(0)
		}
	}

	// Init and start application
	widget := &Widget{}
	widget.Init()
	widget.Run()
}
