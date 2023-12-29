@echo off

set GOOS=linux
set GOARCH=amd64

echo "build start"
go build -o ./weapp
echo "build success"
pause