package main

import "fmt"

type funcBase struct {
}

func (fb *funcBase) Run(args []string) {
	for k, v := range args {
		fmt.Printf("[%v] -> %v", k, v)
	}
}
