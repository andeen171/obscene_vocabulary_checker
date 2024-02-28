package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string
	_, err := fmt.Scanln(&filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	tabooWords := string(file)

	var input string

	for true {
		_, err = fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
			return
		}
		if input == "exit" {
			fmt.Println("Bye!")
			break
		}
		scanner := bufio.NewScanner(strings.NewReader(input))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			word := scanner.Text()
			if strings.Contains(tabooWords, strings.ToLower(word)) {
				fmt.Print(strings.Repeat("*", len(word)))
			} else {
				fmt.Print(word)
			}
			fmt.Print(" ")
		}
	}
}
