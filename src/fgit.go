package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

// For -X Arg
var timestamp = "unknown"
var version = "unknown"

func showVersion() {
	fmt.Println("Version:", version)
	fmt.Println("Build Time:", timestamp)
	os.Exit(0)
}

func debugConnection(url string) bool {
	fmt.Print("Test connection...")
	response, err := http.Head(url)
	if err != nil {
		fmt.Println("Failed")
		fmt.Println("Response create failed\n", err)
		return false
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("Failed")
		return false
	} else {
		fmt.Println("Success")
		return true
	}
}

func debug(url string) bool {
	fmt.Println("" +
		"FastGit Debug Command Line Tool\n" +
		"===============================")
	if url != "--help" && url != "-h" {
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}
		fmt.Println("Remote Address:", url)
		fmt.Print("IP Address: ")
		addr, err := net.LookupIP(strings.Replace(strings.Replace(url, "https://", "", -1), "http://", "", -1))
		if err != nil {
			fmt.Println("Unknown")
		} else {
			fmt.Println(addr)
		}

		fmt.Print("Local Address: ")
		resp, err := http.Get("https://api.ip.sb/ip")
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Unknown ->", err)
		} else {
			s, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Unknown ->", err)
			} else {
				fmt.Printf("[%s]\n", strings.Replace(string(s), "\n", "", -1))
			}
		}

		return debugConnection(url)
	} else {
		fmt.Println("" +
			"SYNTAX\n" +
			"    fgit debug [URL<string>] [--help|-h]\n" +
			"REMARKS\n" +
			"    URL is an optional parameter\n" +
			"    We debug https://hub.fastgit.org by default\n" +
			"    If you want to debug another URL, enter URL param\n" +
			"EXAMPLE\n" +
			"    fgit debug\n" +
			"    fgit debug https://fastgit.org")
		return true
	}
}

func convertToFastGit() bool {
	return convertHelper("https://github.com", "https://hub.fastgit.org")
}

func convertToGitHub() bool {
	return convertHelper("https://hub.fastgit.org", "https://github.com")
}

func convertHelper(oldPrefixValue, newPrefixValue string) bool {
	cmd := exec.Command("git", "remote", "-v")
	buf, err := cmd.Output()
	sBuf := string(buf)
	originUrl := ""
	checkErr(err, "Convert failed.", 8)
	tmp := strings.Split(strings.Replace(strings.Replace(sBuf, "\t", " ", -1), "  ", " ", -1), " ")
	for i := range tmp {
		if strings.HasPrefix(tmp[i], oldPrefixValue) {
			originUrl = tmp[i]
			break
		}
	}
	if originUrl == "" {
		return false
	}
	fmt.Println(originUrl)
	cmd = exec.Command("git", "remote", "set-url", "origin", strings.Replace(originUrl, oldPrefixValue, newPrefixValue, 1))
	_, err = cmd.Output()
	checkErr(err, "Convert failed.", 8)
	return true
}

func checkErr(err error, msg string, exitCode int) {
	if err != nil {
		fmt.Println("Exception:", msg)
		fmt.Println("Tracker:", err)
		os.Exit(exitCode)
	}
}

func conv(target string) {
	switch target {
	case "gh", "github":
		convertToGitHub()
	case "fg", "fastgit":
		convertToFastGit()
	case "-h", "--help":
		fmt.Println("" +
			"FastGit Conv Command Line Tool\n" +
			"==============================\n" +
			"REMARKS\n" +
			"    Convert upstream between github or fastgit automatically\n" +
			"    github and gh means convert to github, fastgit and fg means convert to fastgit\n" +
			"SYNTAX\n" +
			"    fgit conv [--help|-h]\n" +
			"    fgit conv [github|gh|fastgit|fg]\n" +
			"ALIASES\n" +
			"    fgit convert\n" +
			"EXAMPLE\n" +
			"    fgit conv gh")
	default:
		fmt.Println("Invalid args for conv. Use --help to get more information.")
	}
}

func main() {
	if len(os.Args) == 1 || (len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h")) {
		fmt.Println("" +
			"FastGit Command Line Tool\n" +
			"=========================\n" +
			"REMARKS\n" +
			"    We will convert GitHub to FastGit automatically\n" +
			"    Do everything like git\n" +
			"    Build by KevinZonda with GoLang\n" +
			"EXTRA-SYNTAX\n" +
			"    fgit debug [URL<string>] [--help|-h]\n" +
			"    fgit get [URL<string>] [Path<string>] [--help|-h]\n" +
			"    fgit conv [Target<string>] [--help|-h]\n " +
			"    If you want to known more about extra-syntax, try to use --help")
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
		default:
			get("", "")
		case 3:
			get(os.Args[2], "")
		case 4:
			get(os.Args[2], os.Args[3])
		}
		os.Exit(0)
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
		isConvertToFastGit = convertToFastGit()
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
		convertToGitHub()
	}
}

func get(url, fpath string) {
	if url == "" || url == "--help" || url == "-h" {
		fmt.Println("" +
			"FastGit Get Command Line Tool\n" +
			"=============================\n" +
			"REMARKS\n" +
			"    Download with FastGit automatically\n" +
			"SYNTAX\n" +
			"    fgit [--help|-h]\n" +
			"    fgit get [URL<string>]\n" +
			"    fgit get [URL<string>] [Path<string>]\n" +
			"ALIASES\n" +
			"    fgit dl\n" +
			"    fgit download\n" +
			"EXAMPLE\n" +
			"    fgit get https://github.com/fastgitorg/fgit-go/archive/master.zip")
		os.Exit(0)
	} else {
		downloadFile(url, fpath)
	}
}

func downloadFile(url, fpath string) {
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
			fmt.Print("File with the same name exists. New file will cover the old file.\nDo you want to continue? [Y/n]")
			fmt.Scanf("%c", &isContinue)

			switch strings.ToLower(string(isContinue)) {
			case "y":
				os.Remove(fpath)
				goto startDown
			case "n":
				fmt.Println("User cancle the operation.")
				os.Exit(0)
			default:
				fmt.Println("Unknown input, exiting...")
				os.Exit(1)
			}

		}
	}

startDown:
	if strings.HasPrefix(url, "https://github.com/") {
		query := strings.Replace(url, "https://github.com/", "", -1)
		querySplit := strings.Split(query, "/")
		if len(querySplit) > 3 {
			// Source -> fastgitorg/fgit-go/blob/master/fgit.go
			// Target -> fastgitorg/fgit-go/master/fgit.go
			if querySplit[2] == "blob" {
				url = "https://raw.fastgit.org/"
				for _i, _s := range querySplit {
					if _i != 2 {
						// not /blob/
						if _i == len(querySplit)-1 {
							url += _s
						} else {
							url += _s + "/"
						}
					}
				}
				fmt.Println("Redirect ->", url)
			}
		}
	}
	newURL := strings.Replace(url, "https://github.com", "https://download.fastgit.org", -1)
	if newURL != url {
		fmt.Println("Redirect ->", newURL)
	}
	fmt.Println("File ->", fpath)
	fmt.Println("Downloading...")
	resp, err := http.Get(newURL)
	checkErr(err, "Http.Get create failed", 1)
	defer resp.Body.Close()

	out, err := os.Create(fpath)
	checkErr(err, "File create failed", 1)
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	checkErr(err, "io.Copy failed!", 1)
	fmt.Println("Finished.")
	os.Exit(0)
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
