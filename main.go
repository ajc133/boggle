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
		if len(line) < 3 {
			continue
		}
		words = append(words, strings.Split(line, "\n")[0])
	}
	return words, nil
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

	words, _ := readWordList("wordlist.txt")

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			startingCoord, err := b.Get(x, y)
			if err != nil {
				panic(err.Error())
			}
			seenCoords := make([]Square, 0)
			results, err := b.Search(startingCoord, seenCoords, words)
			// TODO: dedupe
			if err != nil {
				panic(err.Error())
			}
			printList(results)

		}
	}
}
