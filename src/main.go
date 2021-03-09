package main

import (
	"github.com/andlabs/ui"
)

func setupUI() {
	mainWindow := ui.NewWindow("Test Updating UI", 640, 480, true)
	mainWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainWindow.Destroy()
		return true
	})

	vbContainer := ui.NewVerticalBox()
	vbContainer.SetPadded(true)

	inputGroup := ui.NewGroup("Input")
	inputGroup.SetMargined(true)

	vbInput := ui.NewVerticalBox()
	vbInput.SetPadded(true)

	inputForm := ui.NewForm()
	inputForm.SetPadded(true)

	message := ui.NewEntry()
	message.SetText("Hello UI World")
	inputForm.Append("What message do you want to show?", message, false)

	showMessageButton := ui.NewButton("Show message")
	clearMessageButton := ui.NewButton("Clear message")

	vbInput.Append(inputForm, false)
	vbInput.Append(showMessageButton, false)
	vbInput.Append(clearMessageButton, false)

	inputGroup.SetChild(vbInput)

	messageGroup := ui.NewGroup("Message")
	messageGroup.SetMargined(true)

	vbMessage := ui.NewVerticalBox()
	vbMessage.SetPadded(true)

	messageLabel := ui.NewLabel("")

	vbMessage.Append(messageLabel, false)

	messageGroup.SetChild(vbMessage)

	vbContainer.Append(inputGroup, false)
	vbContainer.Append(messageGroup, false)

	mainWindow.SetChild(vbContainer)

	showMessageButton.OnClicked(func(*ui.Button) {
		// Update the UI directly as it is called from the main thread
		messageLabel.SetText(message.Text())
	})

	clearMessageButton.OnClicked(func(*ui.Button) {
		// Update the UI directly as it is called from the main thread
		messageLabel.SetText("")
	})

	mainWindow.Show()
}

func main() {
	ui.Main(setupUI)
}
