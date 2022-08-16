package oper

import (
	"fgit-go/shared"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
)

type DebugFunc struct {
}

func (d *DebugFunc) Run(args []string) {
	var isConnectOk bool
	switch len(args) {
	case 0:
		isConnectOk = debug("https://hub.fastgit.org")
	case 1:
		isConnectOk = debug(os.Args[0])
	default:
		fmt.Println("Invalid args for debug. Use --help to get more information.")
	}
	if isConnectOk {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
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
	if url != "--help" && url != "-h" {
		fmt.Println(
			"FastGit Debug Command Line Tool\n" +
				"===============================")
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}
		fmt.Println("Remote Address:", url)
		fmt.Print("IP Address: ")

		addr, err := net.LookupIP(shared.RemoveHttpAndHttps(url))

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
		fmt.Println(shared.DebugHelpMsg)
		return true
	}
}
