package main

import (
	"fmt"
	"strings"
)

func getNewHosts() map[string][]string {
	str := getResponseString(fgEnfHostApi)
	strs := strings.Split(str, "\n")
	m := make(map[string][]string)
	for _, s := range strs {
		if s == "" {
			continue
		}
		ss := strings.Split(s, ";")
		m[ss[0]] = strings.Split(ss[1], ",")
	}
	return m
}

func createHostsContent(m map[string][]string) string {
	var sb strings.Builder
	for ip, prefixes := range m {
		for _, prefix := range prefixes {
			sb.WriteString(ip + " " + prefix + "fastgit.org\n")
		}
	}
	return strings.Trim(sb.String(), "\n")
}

type HostFunc struct {
}

func (h *HostFunc) Run(args []string) {
	fmt.Print("Fetching hosts from FastGit UK Emergency Network Framework...")
	m := getNewHosts()
	fmt.Println("Success!")
	fmt.Println("Please modify your host by adding following content.")
	content := createHostsContent(m)
	fmt.Println(content)
}
