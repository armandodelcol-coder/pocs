package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func countLinesStandardInput(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "0" {
			break
		}

		counts[input.Text()]++
	}
}

func countLinesFiles(f *os.File, counts map[string]int, whichFileIsDuplicaded map[string][]string, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		whichFileIsDuplicaded[input.Text()] = append(whichFileIsDuplicaded[input.Text()], fileName)
	}
}

func removeDuplicates(slice []string) []string {
	var newSlice []string
	counts := make(map[string]int)

	for _, val := range slice {
		counts[val]++
		if counts[val] < 2 {
			newSlice = append(newSlice, val)
		}
	}

	return newSlice
}

func main() {
	counts := make(map[string]int)
	whichFileIsDuplicaded := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLinesStandardInput(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLinesFiles(f, counts, whichFileIsDuplicaded, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			files := whichFileIsDuplicaded[line]
			files = removeDuplicates(files)
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(files, " "))
		}
	}
}