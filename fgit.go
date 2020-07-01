package main

import (
	"bytes"
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

var timestamp string
var version string

func showVersion() {
	fmt.Println("Version:", version)
	fmt.Println("Build Time:", timestamp)
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
	if url != "--help" {
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
			"    fgit debug [URL<string>] [--help]\n" +
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
	fi, err := os.Open(path.Join(".git", "config"))
	checkErr(err, "This is not a git path! Cannot push!", 1)
	defer fi.Close()

	gitConfigByte, err := ioutil.ReadFile(path.Join(".git", "config"))
	checkErr(err, "Cannot read .git config file!", 3)
	gitConfig := string(gitConfigByte)

	isReplaceDo := false
	sb := new(bytes.Buffer)
	iniArray := strings.Split(gitConfig, "\n")
	for i := range iniArray {
		iniArray[i] = strings.Replace(iniArray[i], oldPrefixValue, newPrefixValue, 1)
		isReplaceDo = true
		sb.WriteString(iniArray[i] + "\n")
	}
	fi.Write(sb.Bytes())
	return isReplaceDo
}

func checkErr(err error, msg string, exitCode int) {
	if err != nil {
		fmt.Println("Exception:", msg)
		fmt.Println("Tracker:", err)
		os.Exit(exitCode)
	}
}

func main() {
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "--help") {
		fmt.Println("" +
			"FastGit Command Line Tool\n" +
			"=========================\n" +
			"REMARKS\n" +
			"    We will convert GitHub to FastGit automatically\n" +
			"    Do everything like git\n" +
			"    Build by KevinZonda with GoLang\n" +
			"EXTRA-SYNTAX\n" +
			"    fgit debug [URL<string>] [--help]\n" +
			"    fgit get [URL<string>] [Path<string>] [--help]" +
			"    If you wan to known more about extra-syntax, try to use --help")
		os.Exit(0)
	}

	isConvertToFastGit := false
	isPushOrPull := false

	if os.Args[1] == "debug" {
		var isConnectOk bool
		switch len(os.Args) {
		case 2:
			isConnectOk = debug("https://hub.fastgit.org")
		case 3:
			_i := os.Args[2]
			if strings.HasPrefix(_i, "https://") || strings.HasPrefix(_i, "https://") {
				_i = "http://" + _i
			}
			isConnectOk = debug(_i)
		default:
			fmt.Println("Invalid args for debug. If help wanted, use --help arg.")
		}
		if isConnectOk {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	if os.Args[1] == "get" || os.Args[1] == "dl" || os.Args[1] == "download" {
		switch len(os.Args) {
		default:
			get("", "")
		case 3:
			get(os.Args[2], "")
		case 4:
			get(os.Args[2], os.Args[3])
		}
		os.Exit(0)
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
	if url == "" || url == "--help" {
		fmt.Println("" +
			"FastGit Get Command Line Tool\n" +
			"=============================\n" +
			"REMARKS\n" +
			"    Download with FastGit automatically\n" +
			"SYNTAX\n" +
			"    fgit [--help]\n" +
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
	newUrl := strings.Replace(url, "https://github.com", "https://download.fastgit.org", -1)
	if newUrl != url {
		fmt.Println("Redirect ->", newUrl)
	}
	fmt.Println("Downloading...")
	resp, err := http.Get(newUrl)
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
