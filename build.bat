@echo off
echo ===========================
echo = Go Cross Compile Script =
echo = by Kevin                =
echo ===========================

echo -^> Removing old files
del /s /Q release > nul

SET CGO_ENABLED=0
echo -^> Compiling AMD64
SET GOARCH=amd64

SET GOOS=windows
echo --^> Compiling Windows
go build -ldflags="-s -w" -o release\fgit-windows-amd64.exe src\fgit.go 
SET GOOS=darwin
echo --^> Compiling Darwin
go build -ldflags="-s -w" -o release\fgit-darwin-amd64 src\fgit.go 
SET GOOS=linux
echo --^> Compiling Linux
go build -ldflags="-s -w" -o release\fgit-linux-amd64 src\fgit.go 

SET GOARCH=386
echo -^> Compiling 386
SET GOOS=windows
echo --^> Compiling Windows
go build -ldflags="-s -w" -o release\fgit-windows-386.exe src\fgit.go 
SET GOOS=linux
echo --^> Compiling Linux
go build -ldflags="-s -w" -o release\fgit-linux-386 src\fgit.go 


SET GOARCH=arm
echo -^> Compiling ARM
SET GOOS=linux
echo --^> Compiling Linux
go build -ldflags="-s -w" -o release\fgit-linux-arm src\fgit.go 

SET GOARCH=arm64
echo -^> Compiling ARM64
SET GOOS=linux
echo --^> Compiling Linux
go build -ldflags="-s -w" -o release\fgit-linux-arm64 src\fgit.go 

pause
