package fs

import (
	"os"
	"strings"
)

func PutInTable(Banner string) ([][]string, error) {

	file, err := os.ReadFile(Banner)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(file), "\n")
	var table [][]string
	cou := 0
	for i := 1; i < len(lines); i++ {
		if cou < len(lines) {
			table = append(table, []string{})
		}
		for l := 0; l < 8; l++ {
			table[cou] = append(table[cou], lines[i])
			i++
		}
		cou++
	}
	return table, nil
}

func PutInTableT(Banner string) ([][]string, error) {
	file, err := os.ReadFile(Banner)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(file), "\r\n")
	var table [][]string
	cou := 0
	for i := 1; i < len(lines); i++ {
		if cou < len(lines) {
			table = append(table, []string{})
		}
		for l := 0; l < 8; l++ {
			table[cou] = append(table[cou], lines[i])
			i++
		}
		cou++
	}
	return table, nil
}
