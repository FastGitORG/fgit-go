package main

import (
	"fmt"
	"os"
	"strings"
)

func jsdget(url string, fpath string) {
	if url == "" || url == "--help" || url == "-h" {
		fmt.Println("" +
			"FastGit JsdGet Command Line Tool\n" +
			"=============================\n" +
			"REMARKS\n" +
			"    Download with jsDelivr automatically\n" +
			"SYNTAX\n" +
			"    fgit [--help|-h]\n" +
			"    fgit jsdget [URL<string>]\n" +
			"    fgit jsdget [URL<string>] [Path<string>]\n" +
			"ALIASES\n" +
			"    fgit jdl\n" +
			"EXAMPLE\n" +
			"    fgit jsdget https://github.com/fastgitorg/fgit-go/archive/master.zip")
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
