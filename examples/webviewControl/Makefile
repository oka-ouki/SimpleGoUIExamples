root	:=		$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: clean build-linux build-mac build-win

clean:
	rm -f SimpleUI*
	cd ${root}/src/webview && make clean

build-linux:
	cd ${root}/src && GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/SimpleUI
	cd ${root}/src/webview && make build-linux

build-mac:
	cd ${root}/src && GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/SimpleUI.app
	cd ${root}/src/webview && make build-mac

build-win:
	cd ${root}/src && GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ${root}/SimpleUI.exe
	cd ${root}/src/webview && make build-win
