package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func showVersion() {
	fmt.Println("Version:", version)
	fmt.Println("Build Time:", timestamp)
	os.Exit(0)
}

func main() {
	if len(os.Args) == 1 || (len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h")) {
		fmt.Println(mainHelpMsg)
		os.Exit(0)
	}

	isConvertToFastGit := false
	isPushOrPull := false

	switch os.Args[1] {
	case "debug":
		var isConnectOk bool
		switch len(os.Args) {
		case 2:
			isConnectOk = debug("https://hub.fastgit.org")
		case 3:
			isConnectOk = debug(os.Args[2])
		default:
			fmt.Println("Invalid args for debug. Use --help to get more information.")
		}
		if isConnectOk {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	case "get", "dl", "download":
		switch len(os.Args) {
		case 3:
			get(os.Args[2], "")
		case 4:
			get(os.Args[2], os.Args[3])
		default:
			get("", "")
		}
		os.Exit(0)
	case "jdl", "jsdget":
		switch len(os.Args) {
		case 3:
			jsdget(os.Args[2], "")
		case 4:
			jsdget(os.Args[2], os.Args[3])
		default:
			jsdget("", "")
		}

	case "conv", "convert":
		switch len(os.Args) {
		default:
			fmt.Println("Invalid args for conv. Use --help to get more information.")
		case 3:
			conv(os.Args[2])
		case 2:
			conv("-h")
		}
		os.Exit(0)
	case "-v", "--version", "version":
		showVersion()
	}

	for i := range os.Args {
		if os.Args[i] == "push" || os.Args[i] == "pull" {
			isPushOrPull = true
			break
		}
	}

	if isPushOrPull {
		isConvertToFastGit = convToFastGit()
	}

	cmd := exec.Command("git")

	// Combine to new command
	for i := range os.Args {
		if i != 0 {
			cmd.Args = append(cmd.Args, strings.Replace(os.Args[i], "https://github.com", "https://hub.fastgit.org", -1))
		}
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	checkErr(err, "Command Start Failed!", 4)

	cmd.Wait()
	if isConvertToFastGit {
		convToGitHub()
	}
}
