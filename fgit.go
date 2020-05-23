package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

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
			fmt.Println("Unknown -> ", err)
		} else {
			s, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Unknown -> ", err)
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
		fmt.Println("Exception Detective: ", msg)
		fmt.Println("Tracker: ", err)
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
			isConnectOk = debug(os.Args[2])
		default:
			fmt.Println("Invalid args for debug. If help wanted, use --help arg.")
		}
		if isConnectOk {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
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
