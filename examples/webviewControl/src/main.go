package main

import (
	"os"
	"fmt"
	"net"
	"time"
	"os/exec"
	"runtime"
	"path/filepath"
	"github.com/andlabs/ui"
)

const webviewAppName = "SimpleWebView"

func setupUI() {
	exe, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}

	// Log Setting
	t := time.Now()
	os.Mkdir(filepath.Join(filepath.Dir(exe), "log"), 0777)
	logFilePath := filepath.Join(filepath.Dir(exe), "log", t.Format("20060102") + ".log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		os.Exit(1)
	}
	defer func() {
		logFile.Close()
	}()
	writeLog(logFile, "Start")

	// UDP Setting
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:1201")
	if err != nil {
		writeLog(logFile, fmt.Sprint(err))
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		writeLog(logFile, fmt.Sprint(err))
		os.Exit(1)
	}

	// UI Setting
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

	url := ui.NewEntry()
	url.SetText("https://github.com")
	inputForm.Append("What website do you want to access?", url, false)

	showMessageButton := ui.NewButton("Show URL")
	clearMessageButton := ui.NewButton("Clear URL")

	vbInput.Append(inputForm, false)
	vbInput.Append(showMessageButton, false)
	vbInput.Append(clearMessageButton, false)

	inputGroup.SetChild(vbInput)

	urlGroup := ui.NewGroup("URL")
	urlGroup.SetMargined(true)

	vbMessage := ui.NewVerticalBox()
	vbMessage.SetPadded(true)

	urlLabel := ui.NewLabel("")

	vbMessage.Append(urlLabel, false)

	urlGroup.SetChild(vbMessage)

	vbContainer.Append(inputGroup, false)
	vbContainer.Append(urlGroup, false)

	mainWindow.SetChild(vbContainer)

	showMessageButton.OnClicked(func(*ui.Button) {
		// Update the UI directly as it is called from the main thread
		urlLabel.SetText("Navigating to " + url.Text())
		_, err = conn.Write([]byte(url.Text()))
		if err != nil {
			writeLog(logFile, fmt.Sprint(err))
		}
	})

	clearMessageButton.OnClicked(func(*ui.Button) {
		// Update the UI directly as it is called from the main thread
		urlLabel.SetText("Default Page")
		_, err = conn.Write([]byte("nil"))
		if err != nil {
			writeLog(logFile, fmt.Sprint(err))
		}
	})

	// webview Start
	go startWebview(filepath.Dir(exe), logFile)

	mainWindow.Show()
}

// Logging
func writeLog(f *os.File, s string) {
	f.WriteString("[" + time.Now().Format("2006/01/02 15:04:05") + "] " + s + "\n")
}

func startWebview(path string, logFile *os.File) {
	switch runtime.GOOS {
	case "linux":
		path = filepath.Join(path, webviewAppName)
	case "windows":
		path = filepath.Join(path, webviewAppName + ".exe")
	case "darwin":
		path = filepath.Join(path, webviewAppName + ".app")
	default:
		writeLog(logFile, runtime.GOOS)
		os.Exit(1)
	}
	err := exec.Command(path).Start()
	if err != nil {
		writeLog(logFile, fmt.Sprint(err))
		os.Exit(1)
	}
}

func main() {
	ui.Main(setupUI)
}
