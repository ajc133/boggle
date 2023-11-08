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

// TODO: make a utils package
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

	words, err := readWordList("./assets/static/wordlist.txt")
	if err != nil {
		panic(err.Error())
	}

	// TOOD: Figure out how to set index.html

	r.LoadHTMLGlob("./assets/templates/*")
	// TODO: Make these routes play nice. Maybe enumerate the static files?
	r.Static("/static", "./assets/static")
	r.GET("/solve", func(c *gin.Context) {
		letters := c.Query("letters")

		board, err := boggle.NewBoard(letters)
		if err != nil {
			// FIXME: stop doing 200 once you figure out how to get htmx to show errors
			c.HTML(http.StatusOK, "error.tmpl", gin.H{"error": err})

			// TODO: gin logging
			log.Println(err)
			return
		}

		results, err := board.SearchAll(words)
		if err != nil {
			// FIXME: stop doing 200 once you figure out how to get htmx to show errors
			c.HTML(http.StatusOK, "error.tmpl", gin.H{"error": err})
			// TODO: gin logging
			log.Println(err)
		}

		c.HTML(200, "answers.tmpl", results)
	})

	err = r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
