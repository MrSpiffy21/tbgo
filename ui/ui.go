package ui

import "github.com/rivo/tview"

func showCreateNAPForm(app *tview.Application, pages *tview.Pages) {
	// Reset/Create Form Here
	modal := tview.NewModal().
		SetText("NAP Creation is not yet implemented.").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SwitchToPage("mainMenu")
		})
	pages.AddAndSwitchToPage("createNap", modal, true)
}

func showUpdateNAPForm(app *tview.Application, pages *tview.Pages) {
	// Reset/Create Form Here
	modal := tview.NewModal().
		SetText("NAP Update is not yet implemented.").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SwitchToPage("mainMenu")
		})
	pages.AddAndSwitchToPage("updateNAP", modal, true)
}

func showDeleteNAPForm(app *tview.Application, pages *tview.Pages) {
	// Reset/Create Form Here
	modal := tview.NewModal().
		SetText("NAP Deletion is not yet implemented.").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SwitchToPage("mainMenu")
		})
	pages.AddAndSwitchToPage("deleteNAP", modal, true)
}
