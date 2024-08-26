package main

import (
	"fmt"
	"os"

	"fs/fs-function"
)

func main() {
	args := os.Args[1:]
	if len(args) > 3 || len(args) == 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	input := args[1]
	Banner := "standard"
	if len(args) == 3 {
		Banner = args[2]
	}

	var table [][]string
	var err error

	switch Banner {
	default:
		fmt.Println("Invalid Banner !!!")
		os.Exit(1)
	case "standard":
		table, err = fs.PutInTable("Banner/standard.txt")
	case "shadow":
		table, err = fs.PutInTable("Banner/shadow.txt")
	case "thinkertoy":
		table, err = fs.PutInTableT("Banner/thinkertoy.txt")
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fs.PrintAscii(input, table)
}
