package main

import (
	"fmt"
	"os"
	"strings"
)

type JsdFunc struct {
}

func (j *JsdFunc) Run(args []string) {
	switch len(args) {
	case 0:
		jsdget("", "")
	case 1:
		jsdget(args[0], "")
	case 2:
		jsdget(args[0], args[1])
	}
}

func jsdget(url string, fpath string) {
	if url == "" || url == "--help" || url == "-h" {
		fmt.Println(jsdHelpMsg)
		os.Exit(0)
	} else {
		downloadFile(parseToJsdUrl(url), fpath)
	}
}

func parseToJsdUrl(url string) string {
	u := strings.Split(url, "//")
	switch len(u) {
	case 1:
		break
	case 2:
		url = u[1]
	case 3:
		fmt.Print("Cannot parse url")
		os.Exit(1)
	}
	if !strings.HasPrefix(url, "raw.githubusercontent.com") {
		fmt.Print("Url is not supported!")
		os.Exit(1)
	}
	i := strings.Index(url, "/")
	url = url[i+1:]
	// <OWN>/<Repo>/<Branch>/<Path>
	// <OWN>/<Repo>@<Branch>/<Path>
	url = "https://cdn.jsdelivr.net/gh/" + replaceNth(url, "/", "@", 2)
	fmt.Print("Url -> ")
	fmt.Println(url)
	return url
}
