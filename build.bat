@echo off
echo ===========================
echo = Go Cross Compile Script =
echo = by Kevin                =
echo ===========================

set CGO_ENABLED=0

:: See more in .\fgit.go main
set /P version=<version
for /F %%i in ('powershell -c Get-Date -Format "yyyy-MM-dd"') do ( set date=%%i)
for /F %%i in ('powershell -c Get-Date -Format "HH:mm:ss"') do ( set time=%%i)
set timestamp=%date% %time%
echo Version: %version%
echo Build Time: %timestamp%

cd src

echo -^> Removing old files
del /s /Q release > nul

echo -^> Compiling AMD64
set GOARCH=amd64

set GOOS=windows
echo --^> Compiling Windows
go build -ldflags="-s -w -X main.version=%version% -X 'main.timestamp=%timestamp%'" -o release\fgit-windows-amd64.exe .\
set GOOS=darwin
echo --^> Compiling macOS
go build -ldflags="-s -w -X main.version=%version% -X 'main.timestamp=%timestamp%'" -o release\fgit-macos-amd64 .\
set GOOS=linux
echo --^> Compiling Linux
go build -ldflags="-s -w -X main.version=%version% -X 'main.timestamp=%timestamp%'" -o release\fgit-linux-amd64 .\

set GOARCH=arm
echo -^> Compiling ARM
set GOOS=linux
echo --^> Compiling Linux
go build -ldflags="-s -w -X main.version=%version% -X 'main.timestamp=%timestamp%'" -o release\fgit-linux-arm .\

set GOARCH=arm64
echo -^> Compiling ARM64
set GOOS=linux
echo --^> Compiling Linux
go build -ldflags="-s -w -X main.version=%version% -X 'main.timestamp=%timestamp%'" -o release\fgit-linux-arm64 .\
echo --^> Compiling macOS
go build -ldflags="-s -w -X main.version=%version% -X 'main.timestamp=%timestamp%'" -o release\fgit-macos-arm64 .\

pause
