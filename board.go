package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Width int = 4

type Board struct {
	Board [][]string // FIXME: Make fixed array of Width * Width
}

type Coord struct {
	X int
	Y int
}

func NewBoard() (*Board, error) {
	b := new(Board)
	b.Board = make([][]string, 0)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < Width; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.Split(line, "\n")[0]
		if len(line) != Width {
			return nil, fmt.Errorf("Line is not Width characters!")
		}

		row := strings.Split(line, "")
		b.Board = append(b.Board, row)
	}
	return b, nil
}

func (b *Board) PrintBoard() {
	for _, row := range b.Board {
		for _, char := range row {
			fmt.Printf("%s ", char)
		}
		fmt.Println()
	}
}

func (b *Board) Get(c Coord) (string, error) {
	if c.X < 0 || c.X > 3 || c.Y < 0 || c.Y > 3 {
		return "", fmt.Errorf("Coords are out of bounds")
	}
	return b.Board[c.Y][c.X], nil
}

func (b *Board) GetNeighbors(c Coord) []Coord {
	neighbors := make([]Coord, 0)
	for i := c.X - 1; i < c.X+2; i++ {
		for j := c.Y - 1; j < c.Y+2; j++ {
			if i == c.X && j == c.Y {
				continue
			}

            potentialCoord := Coord{i,j}
			_, err := b.Get(potentialCoord)
			if err != nil {
				// Out of bounds
				continue
			}
			neighbors = append(neighbors, potentialCoord)
		}
	}

	return neighbors
}
