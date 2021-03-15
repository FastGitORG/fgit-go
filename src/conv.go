package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func convertToFastGit() bool {
	return convertHelper("https://github.com", "https://hub.fastgit.org")
}

func convertToGitHub() bool {
	return convertHelper("https://hub.fastgit.org", "https://github.com")
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
