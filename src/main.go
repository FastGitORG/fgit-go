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

	switch strings.ToLower(os.Args[1]) {
	case "debug":
		runByArgs(&DebugFunc{})

	case "get", "dl", "download":
		runByArgs(&GetFunc{})

	case "jdl", "jsdget", "jsd":
		runByArgs(&JsdFunc{})

	case "conv", "convert":
		runByArgs(&ConvFunc{})

	case "host":
		runByArgs(&HostFunc{})

	case "-v", "--version", "version":
		showVersion()
	}

	if os.Args[2] == "push" || os.Args[2] == "pull" {
		isConvertToFastGit = convToFastGit()
	}

	cmd := exec.Command("git")

	// Combine to new command
	for i := range os.Args[1:] {
		cmd.Args = append(cmd.Args, strings.Replace(os.Args[i], "https://github.com", "https://hub.fastgit.org", -1))
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

func runByArgs(fb IFuncBase) {
	if len(os.Args) < 3 {
		fb.Run([]string{})
	}
	fb.Run(os.Args[3:])
}
