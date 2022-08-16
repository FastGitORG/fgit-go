package oper

import (
	"fgit-go/shared"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ConvFunc struct {
}

func (c *ConvFunc) Run(args []string) {
	switch len(args) {
	case 1:
		conv(args[0])
	case 0:
		conv("-h")
	default:
		fmt.Println("Invalid args for conv. Use --help to get more information.")
	}
	os.Exit(0)
}
func ConvToFastGit() bool {
	return convHelper("https://github.com", shared.GitMirror)
}

func ConvToGitHub() bool {
	return convHelper(shared.GitMirror, "https://github.com")
}

func conv(target string) {
	switch target {
	case "gh", "github":
		ConvToGitHub()
	case "fg", "fastgit":
		ConvToFastGit()
	case "-h", "--help":
		fmt.Println(shared.ConvHelpMsg)
	default:
		fmt.Println("Invalid args for conv. Use --help to get more information.")
	}
}

func convHelper(oldPrefixValue, newPrefixValue string) bool {
	cmd := exec.Command("git", "remote", "-v")
	buf, err := cmd.Output()
	sBuf := string(buf)
	originUrl := ""
	shared.CheckErr(err, "Convert failed.", 8)
	tmp := strings.Split(
		strings.Replace(
			strings.Replace(sBuf, "\t", " ", -1),
			"  ", " ", -1),
		" ")

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
	shared.CheckErr(err, "Convert failed.", 8)
	return true
}
