package main

import (
	"fmt"
	"os"
)

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
