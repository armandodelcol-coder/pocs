package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("The program will count the duplicated lines of your standard input.")
	fmt.Println("If you put a zero '0' the program will stop and print the count of duplicated words.")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "0" {
			break
		}

		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}