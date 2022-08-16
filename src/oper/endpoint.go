package oper

import (
	"fgit-go/shared"
	"fmt"
	"os"
)

type EndPointFunc struct {
}

func (c *EndPointFunc) Run(args []string) {
	fmt.Println(
		"FastGit EndPoint Command Line Tool\n" +
			"==================================")
	fmt.Printf("https://github.com                -> %s\n", shared.GitMirror)
	fmt.Printf("https://raw.githubusercontent.com -> %s\n", shared.RawMirror)
	fmt.Printf("GitHub Download                   -> %s\n", shared.DownloadMirror)
	os.Exit(0)
}
