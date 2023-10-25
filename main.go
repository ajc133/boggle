package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readWordList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open word list")
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	words := []string{}
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return nil, err
		}
		words = append(words, strings.Split(line, "\n")[0])
	}
	return words, nil
}

func threeLetterMap(words []string) map[string][]string {
	m := make(map[string][]string)
	for _, word := range words {
		substring := word[:3]
		if m[substring] == nil {
			m[substring] = []string{word}
		} else {
			m[substring] = append(m[substring], word)
		}
	}
	return m
}

func main() {

	words, _ := readWordList("wordlist.txt")
	wordMap := threeLetterMap(words)
	for trigraph, words := range wordMap {
		fmt.Println(len(words), trigraph)
	}

}
