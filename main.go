package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		cleanIn := cleanInput(scanner.Text())
		if len(cleanIn) == 0 {
			fmt.Println("Your command was:")
			continue
		}
		firstWord := cleanIn[0]
		fmt.Println("Your command was:", firstWord)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}