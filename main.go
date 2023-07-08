/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"ftp_client/cmd"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	LabelHost     = "Host: "
	LabelUser     = "Username: "
	LabelPassword = "Password: "
	LabelPort     = "Port: "
	LabelConnect  = "Connect"
)

func main() {
	cmd.Execute()

	app := tview.NewApplication()

	inputHost := tview.NewInputField().SetLabel(LabelHost).SetFieldWidth(30).SetFieldTextColor(tcell.ColorBlack.TrueColor())
	inputUser := tview.NewInputField().SetLabel(LabelUser).SetFieldWidth(30).SetFieldTextColor(tcell.ColorBlack.TrueColor())
	inputPw := tview.NewInputField().SetLabel(LabelPassword).SetFieldWidth(30).SetFieldTextColor(tcell.ColorBlack.TrueColor())
	inputPort := tview.NewInputField().SetLabel(LabelPort).SetFieldWidth(30).SetFieldTextColor(tcell.ColorBlack.TrueColor())
	btnConnect := tview.NewButton(LabelConnect).SetLabelColor(tcell.ColorBlack.TrueColor())

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		// Top
		AddItem(tview.NewFlex().
			//AddItem(grid, 0, 1, false), 0, 1, false.
			AddItem(inputHost, 0, 1, false).
			AddItem(inputUser, 0, 1, false).
			AddItem(inputPw, 0, 1, false).
			AddItem(inputPort, 0, 1, false).
			AddItem(btnConnect, 0, 1, false).
			//AddItem(tview.NewTextView().SetLabelWidth(10), 0, 1, false).
			AddItem(tview.NewTextView().SetLabelWidth(10), 0, 1, false), 0, 1, false).

		// Middle
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Local"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Remote"), 0, 1, false), 0, 20, false).

		// Bottom
		AddItem(tview.NewBox().SetBorder(true).SetTitle("File Transfer"), 0, 4, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
