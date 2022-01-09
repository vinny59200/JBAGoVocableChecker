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

func solve(s []int) {
	// put your code here
	for index := range s {
		// processing slice elements
		fmt.Println(s[index] * 2)
	}
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
