package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var filter []string
	if len(os.Args) > 2 {
		if os.Args[1] == "test" {
			fmt.Println("Setting up test hooks")
			for _, args := range os.Args[2:] {
				filter = append(filter, args)
			}
			fmt.Println(filter)
			test(filter)
			return
		}
		fmt.Println("Command not supported")
		return
		// for _, args := range os.Args[1:] {
		// 	if args == "test" {
		// 		fmt.Println("Jarvis D3PL0Y1NG some nasty stuff!!")
		// 		deploy()
		// 	}
		// 	fmt.Println(args)
		// }
	}
	fmt.Println("No command found")
}

func test(filter []string) {
	cmd := exec.Command("sls", "invoke", "test")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	for _, f := range filter {
		if strings.Contains(outStr, f) {
			fmt.Println("CAUGHT", f)
			os.Exit(1)
		}
	}
	// fmt.Println(strings.Contains(outStr, "failing"))
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}