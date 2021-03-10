root	:=		$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: clean build-linux build-mac build-win

clean:
	rm -f SimpleWebView*

build-linux:
	cd src && GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/../../SimpleWebView

build-mac:
	cd src && GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/../../SimpleWebView.app

build-win:
	cd src && GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags="-s -w -H windowsgui" -o ${root}/../../SimpleWebView.exe
