package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
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
