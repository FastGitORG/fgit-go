package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

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
	if len(os.Args) == 1 {
		fmt.Println("" +
			"FastGit Command Line Tool\n" +
			"=========================\n" +
			"We will convert GitHub to FastGit automatically\n" +
			"Do everything like git\n" +
			"Build by KevinZonda with GoLang")
		os.Exit(0)
	}

	isConvertToFastGit := false
	isPush := false
	for i := range os.Args {
		if os.Args[i] == "push" {
			isPush = true
			break
		}
	}

	if isPush {
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
