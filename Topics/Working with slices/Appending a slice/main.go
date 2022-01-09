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

func solve(s1, s2 []string) []string {
	// put your code here

	var s = append(s1, s2...)
	return s
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
	strconv.Itoa(1)
	sort.Strings(words)
	return words
}
