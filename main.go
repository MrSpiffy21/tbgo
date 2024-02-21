package main

import (
	"github.com/MrSpiffy21/tbgo/ui"
	"github.com/rivo/tview"
	"sync"
)

// API Connection Info
var host, username, password string

// NAP Creation Control
var createNAP, updateNAP, deleteNAP, isNAP, isPBX, napProxyPoll, napProxyPollEnable, napColumn bool

// NAP Control Info
var customerName, configName, digitMap, proxyHost, phoneNumbers, portRange, sipTransport, rDefRouteGroups, napcRouteGroups, napProfile, remoteMothodSIP, remoteMethodRTP, localMethodSIP, localMethodRTP string

// Max Calls
var maxTotalCalls int

func main() {
	var wg sync.WaitGroup
	app := tview.NewApplication()
	pages := tview.NewPages()

	// The menu page
	mainMenu := tview.NewList().
		AddItem("Create NAP", "Start the NAP creation process", 'c', func() {
			ui.showCreateNAPForm(app, pages)
		}).
		AddItem("Update NAP", "Start the NAP update process", 'u', func() {
			ui.showUpdateNAPForm(app, pages)
		}).
		AddItem("Delete NAP", "Start the NAP deletion process", 'd', func() {
			ui.showDeleteNAPForm(app, pages)
		}).
		AddItem("Exit", "Exit the program", 'e', func() {
			app.Stop()
		})

	pages.AddAndSwitchToPage("mainMenu", mainMenu, true)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.SetRoot(pages, true).Run(); err != nil {
			panic(err)
		}
	}()
	wg.Wait()
}
