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
		fmt.Print(line)
                words = append(words, strings.Split(line, "\n")[0])
	}
        return words, nil
}
func main() {

        words, _ := readWordList("wordlist.txt")
        for _, word := range words{
                fmt.Println(word)
        }



}
