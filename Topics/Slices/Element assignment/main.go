package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(s []string, i int) {
	// Put your code here
	//s = make([]string, i)

	s[i] = "my string value"
	fmt.Print()
}

func readWords(filePath string) []string {
	fmt.Scan(&filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, strings.ToLower(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Strings(words)
	strconv.Itoa(1)
	return words
}
