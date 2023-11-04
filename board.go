package boggle

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var MINWORDLEN = 3
var WIDTH int = 4

// TODO: handle 'QU'
type Square struct {
	Letter string
	X      int
	Y      int
}

type Board struct {
	Board [][]Square // FIXME: Make fixed array of Width * Width
}

func same(c Square, d Square) bool {
	if c.X == d.X && c.Y == d.Y {
		return true
	}
	return false

}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsSquare(s []Square, e Square) bool {
	for _, a := range s {
		if same(a, e) {
			return true
		}
	}
	return false
}

// TOOD: lower-case everything
func NewBoard() (*Board, error) {
	b := new(Board)
	b.Board = make([][]Square, WIDTH)
	reader := bufio.NewReader(os.Stdin)
	for rowNum := 0; rowNum < WIDTH; rowNum++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.Split(line, "\n")[0]
		if len(line) != WIDTH {
			return nil, fmt.Errorf("Line is not Width characters!")
		}

		rowSlice := strings.Split(line, "")
		squares := make([]Square, 0)
		for colNum, c := range rowSlice {
			squares = append(squares, Square{Letter: c, X: colNum, Y: rowNum})
		}
		b.Board[rowNum] = squares
	}
	return b, nil
}

func (b *Board) PrintBoard() {
	for _, row := range b.Board {
		for _, square := range row {
			fmt.Printf("%s ", square.Letter)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *Board) Get(x int, y int) (Square, error) {
	if x < 0 || x > 3 || y < 0 || y > 3 {
		return Square{}, fmt.Errorf("Coords are out of bounds")
	}
	return b.Board[y][x], nil
}

// Only shows neighbors that have not been visited yet
func (b *Board) GetNewNeighbors(c Square, exclude []Square) []Square {
	neighbors := make([]Square, 0)
	x := c.X
	y := c.Y
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i == x && j == y {
				continue
			}

			square, err := b.Get(i, j)
			if err != nil {
				// Out of bounds
				continue
			}
			if !ContainsSquare(exclude, square) {
				neighbors = append(neighbors, square)
			}
		}
	}

	return neighbors
}

func ConcatSquares(s []Square) (string, error) {
	var builder strings.Builder
	for _, e := range s {
		_, err := builder.WriteString(e.Letter)
		if err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}

func GetPrefixMatches(prefix string, words []string) []string {
	matchingWords := make([]string, 0)

	for _, word := range words {
		if strings.HasPrefix(word, prefix) {
			matchingWords = append(matchingWords, word)
		}
	}

	return matchingWords
}

func WeFoundAWord(word string, prefixMatches []string) bool {
	return ContainsString(prefixMatches, word) && (len(word) >= MINWORDLEN)

}

func (b *Board) Search(currentSquare Square, seenSquares []Square, potentialWords []string) ([]string, error) {
	results := make([]string, 0)

	seenSquares = append(seenSquares, currentSquare)
	wordSoFar, err := ConcatSquares(seenSquares)
	if err != nil {
		return nil, err
	}

	prefixMatches := GetPrefixMatches(wordSoFar, potentialWords)

	// No words match from here on
	if len(prefixMatches) == 0 {
		return []string{}, nil
	}

	if WeFoundAWord(wordSoFar, prefixMatches) {
		results = append(results, wordSoFar)
	}

	neighbors := b.GetNewNeighbors(currentSquare, seenSquares)
	for _, neighborSquare := range neighbors {
		subResults, err := b.Search(neighborSquare, seenSquares, prefixMatches)
		if err != nil {
			return nil, err
		}
		results = append(results, subResults...)
	}
	return results, nil

}
