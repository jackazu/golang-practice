package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {

	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for _, name := range searchDuplicateLineHaveFile(files, counts) {
				fmt.Println(name)
			}
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func searchDuplicateLineHaveFile(files []string, counts map[string]int) []string {
	duplicateLineHaveFileNames := []string{}
	for _, fileName := range files {
		text, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(text), "\n") {
			if counts[line] > 1 {
				duplicateLineHaveFileNames = append(duplicateLineHaveFileNames, fileName)
				break
			}
		}
	}

	return duplicateLineHaveFileNames
}
