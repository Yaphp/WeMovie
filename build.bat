@echo off

set GOOS=linux
set GOARCH=amd64

go build -o ./wemovie wemovie

echo "build success"
pause