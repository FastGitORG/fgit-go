package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type GetFunc struct {
}

func (g *GetFunc) Run(args []string) {
	switch len(args) {
	case 1:
		get(args[0], "")
	case 2:
		get(os.Args[0], os.Args[1])
	default:
		get("", "")
	}
	os.Exit(0)
}

func get(url, fpath string) {
	if url == "" || url == "--help" || url == "-h" {
		fmt.Println(getHelpMsg)
	} else {
		getFile(url, fpath)
	}
}

func getFile(url, fpath string) {
	urlSplit := strings.Split(url, "/")
	filename := urlSplit[len(urlSplit)-1]
	if fpath == "" {
		downloadFile(url, filename)
	}

	if isExists(fpath) {
		if isDir(fpath) {
			fpath = path.Join(fpath, filename)
			downloadFile(url, fpath)
		} else {
			isContinue := ' '
			fmt.Print(
				"File with the same name exists. New file will cover the old file.\n" +
					"Do you want to continue? [Y/n]")
			fmt.Scanf("%c", &isContinue)

			switch strings.ToLower(string(isContinue)) {
			case "y":
				os.Remove(fpath)
				break
				// goto startDown
			case "n":
				fmt.Println("User canceled the operation.")
				os.Exit(0)
			default:
				fmt.Println("Unknown input, exiting...")
				os.Exit(1)
			}

		}
	}

	url = parseToGetUrl(url)

	fmt.Println("Redirect ->", url)

	newURL := strings.Replace(url, "https://github.com", downloadMirror, -1)
	if newURL != url {
		fmt.Println("Redirect ->", newURL)
	}
	fmt.Println("File ->", fpath)
	fmt.Println("Downloading...")

	downloadFile(newURL, fpath)

	fmt.Println("Finished.")
}

func parseToGetUrl(url string) string {
	if !strings.HasPrefix(url, "https://github.com/") {
		return url
	}
	query := strings.Replace(url, "https://github.com/", "", -1)
	querySplit := strings.Split(query, "/")
	if len(querySplit) > 3 {
		// Source -> fastgitorg/fgit-go/blob/master/fgit.go
		// Target -> fastgitorg/fgit-go/master/fgit.go
		if querySplit[2] == "blob" {
			url = rawMirror
			for _i, _s := range querySplit {
				if _i != 2 {
					url += "/" + _s
				}
			}
		}
	}
	return url
}
