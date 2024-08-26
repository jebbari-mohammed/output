package fs

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func PrintAscii(input string, table [][]string) {
	args := os.Args[1:]
	output := args[0][9:]
	finaltxt := ""
	outputfile, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	inputSplit := strings.Split(input, "\\n")
	for _, str := range inputSplit {
		for _, char := range str {
			if int(char) > 126 || int(char) < 32 {
				fmt.Println("invalid input")
				return
			}
		}
	}
	var allNewLine bool

	for i := 0; i < len(inputSplit); i++ {
		if inputSplit[i] != "" {
			allNewLine = true
		}
	}
	if !allNewLine {
		inputSplit = inputSplit[1:]
	}

	for _, str := range inputSplit {
		if str == "" {
			finaltxt = finaltxt + "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range str {
				finaltxt = finaltxt + table[int(char)-32][i]
			}
			finaltxt = finaltxt + "\n"
		}
	}
	if _, err := io.WriteString(outputfile, finaltxt); err != nil {
		panic(err)
	}
}
