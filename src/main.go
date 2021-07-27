package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var words = []string{"Age", "Aggregate", "Bone", "Cave", "Dandy", "Empire", "Funky", "Guitar", "Hello", "Isotope", "Legend", "Low", "New", "Peach", "Pope", "Poppy", "Stone", "Stop"}

/* Returns a string with spaces between each character */
func spaceOut(s string) string {
	if len(s) < 1 {
		panic("Input string must have at least length of 1. Length: " + strconv.Itoa(len(s)))
	}
	if len(s) == 1 {
		return s
	}

	final := string(s[0])
	for i := 1; i < len(s); i++ {
		final += " " + string(s[i])
	}
	return final
}

/* Replaces characters with an underscore if their index in hidden is true */
func hideLetters(s string, hidden []bool) string {
	if len(s) < 1 {
		panic("Input string must have at least length of 1. Length: " + strconv.Itoa(len(s)))
	}

	// Special case if string length is 1
	if len(s) == 1 {
		if hidden[0] {
			return "_"
		} else {
			return string(s[0])
		}
	}

	final := ""
	if hidden[0] {
		final += "_"
	} else {
		final += string(s[0])
	}
	for i := 1; i < len(s); i++ {
		if hidden[i] {
			final += " _"
		} else {
			final += " " + string(s[i])
		}
	}
	return final
}

func main() {
	fmt.Println("H A N G M A N")
	rand.Seed(time.Now().UnixNano())
	word := strings.ToUpper(words[rand.Intn(len(words))])

	hiddenSpaces := len(word)
	hiddenLetters := make([]bool, hiddenSpaces)
	for i := range hiddenLetters {
		hiddenLetters[i] = true
	}
	var previousGuesses []string

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nGuessed: " + ArrayToString(previousGuesses))
		if hiddenSpaces == 0 {
			fmt.Println(spaceOut(word))
			fmt.Println("You win!")
			return
		}

		fmt.Println(hideLetters(word, hiddenLetters))
		fmt.Print("Guess: ")
		text, _ := reader.ReadString('\n')
		text = strings.ToUpper(strings.TrimSpace(text))

		if len(text) != 1 {
			if text == "EXIT" {
				fmt.Println("Goodbye.")
				return
			}
			if text == word {
				fmt.Println("You guessed it!")
				return
			}
		}

		if ArrayIndexOf(previousGuesses, text) != -1 {
			fmt.Println("You have already guessed \"" + text + "\"")
			continue
		}

		correctGuess := false
		for index, element := range word {
			if string(text[0]) == string(element) {
				if !correctGuess {
					correctGuess = true
					fmt.Println("Correct!")
					previousGuesses = append(previousGuesses, string(text[0]))
				}
				hiddenLetters[index] = false
				hiddenSpaces--
			}
		}
		if !correctGuess {
			fmt.Println("Word does not contain " + text)
		}
	}
}
