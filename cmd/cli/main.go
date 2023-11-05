package main

import (
	"bufio"
	"fmt"
	"github.com/ajc133/boggle"
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

func getInput() (string, error) {
	concatenated := ""

	reader := bufio.NewReader(os.Stdin)

	// Read at most 4 lines of 6 characters each (including newline).
	for i := 0; i < 4; i++ {
		fmt.Printf("Enter line %d (up to 5 characters): ", i)
		line, _ := reader.ReadString('\n')
		line = strings.Split(line, "\n")[0]

		// If the line length is less than 5 characters, break the loop.
		if len(line) < 4 {
			return "", fmt.Errorf("Too few characters")
		} else if len(line) > 4 {
			return "", fmt.Errorf("Too many characters")
		}
		concatenated += line
	}

	return concatenated, nil
}

func printList(l []string) {
	for _, word := range l {
		fmt.Println(word)
	}

}

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}
	fmt.Println(input)
	b, err := boggle.NewBoard(input)
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
			seenCoords := make([]boggle.Square, 0)
			results, err := b.Search(startingCoord, seenCoords, words)
			// TODO: dedupe
			if err != nil {
				panic(err.Error())
			}
			printList(results)

		}
	}
}
