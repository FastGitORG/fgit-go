package main

import (
	"fgit-go/config"
	"fgit-go/oper"
	"fgit-go/shared"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	timestamp = "unknown"
	version   = "unknown"
)

func showVersion() {
	fmt.Println("Version:", version)
	fmt.Println("Build Time:", timestamp)
	os.Exit(0)
}

func main() {
	if len(os.Args) == 1 || (len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h")) {
		fmt.Println(shared.MainHelpMsg)
		os.Exit(0)
	}

	config.ReadConfig()

	isConvertToFastGit := false

	switch strings.ToLower(os.Args[1]) {
	case "debug":
		runByArgs(&oper.DebugFunc{})

	case "get", "dl", "download":
		runByArgs(&oper.GetFunc{})

	case "conv", "convert":
		runByArgs(&oper.ConvFunc{})

	case "-v", "--version", "version":
		showVersion()
	}

	if os.Args[2] == "push" || os.Args[2] == "pull" {
		isConvertToFastGit = oper.ConvToFastGit()
	}

	cmd := exec.Command("git")

	// Combine to new command
	for i := range os.Args[1:] {
		cmd.Args = append(cmd.Args,
			strings.Replace(os.Args[i], "https://github.com", shared.GitMirror, -1))
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	shared.CheckErr(err, "Command Start Failed!", 4)

	cmd.Wait()
	if isConvertToFastGit {
		oper.ConvToGitHub()
	}
}

func runByArgs(fb IFuncBase) {
	if len(os.Args) < 3 {
		fb.Run([]string{})
	}
	fb.Run(os.Args[3:])
}
