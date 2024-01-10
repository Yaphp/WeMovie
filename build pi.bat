set GOOS=linux
set GOARCH=arm
set GOARM=6
set CGO_ENABLED=0

@echo off
echo "build start"
go build -o ./weapp
echo "build done"