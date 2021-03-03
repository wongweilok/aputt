package main

import (
	"github.com/rivo/tview"
)

func Browse() (string, tview.Primitive) {
	// Temporary place holder for browsing feature
	browse := tview.NewTextView().
		SetText("This is a page for browsing list.")

	browse.SetBorder(true)

	return "Browse", browse
}
