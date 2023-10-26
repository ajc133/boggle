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

type Board struct {
	Board [4][4]string
}

type Coord struct {
	X int
	Y int
}

func NewBoard() *Board {
	b := new(Board)
	b.Board = [4][4]string{
		{"a", "b", "c", "d"},
		{"g", "f", "g", "h"},
		{"i", "j", "k", "l"},
		{"m", "n", "o", "p"}}
	return b
}

func (b *Board) PrintBoard() {
	for _, row := range b.Board {
		for _, char := range row {
			fmt.Printf("%s ", char)
		}
		fmt.Println()
	}
}

func (b *Board) Get(x int, y int) (string, error) {
	if x < 0 || x > 3 || y < 0 || y > 3 {
		return "", fmt.Errorf("Coords are out of bounds")
	}
	return b.Board[y][x], nil
}

func (b *Board) GetNeighbors(x int, y int) []string {

	neighbors := make([]string, 0)
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i == x && j == y {
				continue
			}

			neighbor, err := b.Get(i, j)
			if err != nil {
				continue
			}
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func main() {
	b := NewBoard()
	b.PrintBoard()

	fmt.Println(b.GetNeighbors(0, 0))
	fmt.Println(b.GetNeighbors(1, 1))

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
