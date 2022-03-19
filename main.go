package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	var filePath string
	words := readWords(filePath)
	for {
		if wordByWordScan(words) {
			break
		}
	}
	fmt.Println("Bye!")
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
	return words
}
func wordByWordScan(words []string) bool {
	var tabooSentence string
	fmt.Scan(&tabooSentence)

	if tabooSentence == "exit" {
		return true
	}

	regex := regexp.MustCompile("[, -!?:]+")
	tabooSentenceWords := regex.Split(tabooSentence, -1)
	//fmt.Print(tabooSentenceWords)
	for _, tabooSentenceWord := range tabooSentenceWords {
		if contains(words, strings.ToLower(tabooSentenceWord)) {
			var censoredWord string
			//fmt.Println(tabooSentenceWord)
			for i := 0; i < len(tabooSentenceWord); i++ {
				censoredWord += "*"
			}
			re := regexp.MustCompile(tabooSentenceWord)
			tabooSentence = re.ReplaceAllString(tabooSentence, censoredWord)
			//fmt.Println(tabooSentence)
		}
	}

	fmt.Println(tabooSentence)

	return false
}
func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}
