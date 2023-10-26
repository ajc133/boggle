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
		if len(line) < 3 || len(line) > 10 {
			continue
		}
		words = append(words, strings.Split(line, "\n")[0])
	}
	return words, nil
}

func getMatches(prefix string, words []string) []string {
	matchingWords := make([]string, 0)

	for _, word := range words {
		if strings.HasPrefix(word, prefix) {
			matchingWords = append(matchingWords, word)
		}
	}

	return matchingWords
}

func getInput(reader *bufio.Reader) (string, error) {
	fmt.Print("\n\ninput text: ")
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Split(line, "\n")[0], nil
}

func printList(l []string) {
	for _, word := range l {
		fmt.Println(word)
	}

}

func main() {
	b, err := NewBoard()
	if err != nil {
		panic(err.Error())
	}
	b.PrintBoard()

    for _, neighbor := range b.GetNeighbors(Coord{0, 1}) {
        fmt.Println(neighbor)
    }

	// words, _ := readWordList("wordlist.txt")

	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	input, err := getInput(reader)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		os.Exit(1)
	// 	}

	// 	printList(getMatches(input, words))
	// }

}
