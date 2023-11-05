package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ajc133/boggle"
	"github.com/gin-gonic/gin"
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

func main() {
	r := gin.Default()

	// TOOD: Figure out how to set index.html

	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")
	r.GET("/solve", func(c *gin.Context) {
		letters, ok := c.GetQueryArray("letters")
		if !ok {
			// FIXME: stop doing 200 once you figure out how to get htmx to show errors
			c.HTML(http.StatusOK, "error.tmpl", gin.H{"error": "Did not find 'letters' query array"})

			// TODO: gin logging
			log.Println("No 'letters' param passed")
			return
		}
		input := strings.Join(letters, "")

		// 17 because 'qu' is one letter
		if len(input) < 16 || len(input) > 17 {
			// FIXME: stop doing 200 once you figure out how to get htmx to show errors
			c.HTML(http.StatusOK, "error.tmpl", gin.H{"error": "Fill in every box"})

			// TODO: gin logging
			log.Println("Wrong number of letters")
			return
		}
		board, err := boggle.NewBoard(input)
		if err != nil {
			panic(err.Error())
		}

		words, err := readWordList("wordlist.txt")
		if err != nil {
			panic(err.Error())
		}

		var results []string
		for x := 0; x < 4; x++ {
			for y := 0; y < 4; y++ {
				startingCoord, err := board.Get(x, y)
				if err != nil {
					panic(err.Error())
				}
				seenCoords := make([]boggle.Square, 0)
				currentResults, err := board.Search(startingCoord, seenCoords, words)
				results = append(results, currentResults...)
				// TODO: dedupe
				if err != nil {
					panic(err.Error())
				}

			}
		}

		c.HTML(200, "answers.tmpl", results)
	})

	err := r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
