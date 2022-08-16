package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadFile(url, path string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err, "Http.Get create failed!", 1)
	req.Header.Set("User-Agent", "fgit/"+version)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	checkErr(err, "Http request failed!", 1)

	out, err := os.Create(path)
	checkErr(err, "File create failed!", 1)
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	checkErr(err, "io.Copy failed!", 1)
	fmt.Println("Finished.")
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

func checkErr(err error, msg string, exitCode int) {
	if err != nil {
		fmt.Println("Exception:", msg)
		fmt.Println("Tracker:", err)
		os.Exit(exitCode)
	}
}

func replaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

func removeHttpAndHttps(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url[7:]
	}
	if strings.HasPrefix(url, "https://") {
		return url[8:]
	}
	return url
}

func replacePrefix(str, prefix, after string) string {
	if len(str) < len(after) {
		return str
	}
	if str == prefix {
		return after
	}

	if strings.HasPrefix(str, prefix) {
		return after + str[len(prefix)+1:]
	}
	return str
}
