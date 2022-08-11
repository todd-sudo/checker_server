.PHONY: run

run: start

start:
	go run main.go


build:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/app_linux_amd64


build_linux_all:
	env GOOS=linux GOARCH=386 go build -o bin/linux/app_linux_386 -v main.go
	env GOOS=linux GOARCH=arm go build -o bin/linux/app_linux_arm -v main.go
	env GOOS=linux GOARCH=arm64 go build -o bin/linux/app_linux_arm64 -v main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/app_linux_amd64 -v main.go

build_windows_all:
	env GOOS=windows GOARCH=amd64 go build -o bin/windows/app_win_64.exe -v main.go
	env GOOS=windows GOARCH=386 go build -o bin/windows/app_win_32.exe -v main.go

build_all_platform:build_linux build_windows
	

.DEFAULT_GOAL := run