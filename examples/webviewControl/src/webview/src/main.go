package main

import (
	"os"
	"net"
	"github.com/webview/webview"
)

const defaultHTML = "data:text/html,<html><body><center><h1>WebView Example</h1></center></body></html>"

func main() {
	// UDP Setting
	udpAddr, err := net.ResolveUDPAddr("udp", ":1201")
	if err != nil {
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		os.Exit(1)
	}

	// webview Setting
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("WebView Example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(defaultHTML)

	// Start UDP Handle
	go handleClient(conn, w)

	w.Run()
}

func handleClient(conn *net.UDPConn, w webview.WebView) {
	var buf [512]byte

	for {
		n, _, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			os.Exit(1)
		}
		url := string(buf[0:n])
		if len(url) > 0 {
			if !checkURL(url) {
				url = defaultHTML
			}
			w.Navigate(url)
		}
	}
}

func checkURL(url string) bool {
	return (len(url) > 7 && url[:7] == "http://") || (len(url) > 8 && url[:8] == "https://")
}
