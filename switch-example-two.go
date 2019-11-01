package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		/*
		 freebsd, Openbsd
		 plan9, windows
		*/
		fmt.Printf("%S. \n", os)
	}
}
