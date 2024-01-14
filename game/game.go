package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"goWordle/checkGuess"
	"goWordle/checkWord"
	"goWordle/wordChoice"
)

// Function to display game state info
func displayResults(guesses int, guessedList []string) {
  fmt.Printf("\n")
  fmt.Printf("Guesses left: %d\n", guesses)
  fmt.Println("Past guesses:")
  for i := 0; i < len(guessedList); i++{
    fmt.Println(guessedList[i])
    fmt.Printf("\033[0m") // Text color reset
  }
  fmt.Printf("\n")
}

// Main game function
func main() {
  fmt.Printf(`
             _ _ _           _ _     
     ___ ___| | | |___ ___ _| | |___ 
    | . | . | | | | . |  _| . | | -_|
    |_  |___|_____|___|_| |___|_|___|
    |___|                            
  ` + "\n")
  fmt.Println("Wordle written in Go! goWordle!")
  fmt.Println("\nGuess the 5 letter word, in 6 tries!")
  fmt.Println("\033[93mLetters in Yellow are correct, BUT in the wrong place")
  fmt.Println("\033[92mLetters in Green are correct, AND in the correct place")
  fmt.Println("\033[0m") // Text color reset

  // Random word chosen
  word := wordChoice.Word()
  guesses := 6
  var guessed string
  var guessedList []string

  // Game loop
  for {
    fmt.Println("Your guess?")

    // Reading user input
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    err := scanner.Err()
    guess := scanner.Text()

    if err != nil {
      log.Fatal(err)
    }

    // Checking validity of input before proceeding with game
    if checkGuess.Check(guess) {
      guesses -= 1
      guessed = checkWord.Check(word, guess)
      guessedList = append(guessedList, guessed)
    }
    // Displays game state: Guesses left and past guesses
    displayResults(guesses, guessedList)

    // Win condition check; breaks loop
    if strings.ToLower(guess) == strings.ToLower(word) {
      fmt.Printf("Correct! The word was: %v\n", word)
      fmt.Println("\033[92mYOU WIN!") // green
      fmt.Println("\033[0m") // Text color reset
      break
    }

    // Lose condition check; breaks loop
    if (guesses == 0) && (strings.ToLower(guess) != strings.ToLower(word)) {
      fmt.Println("No guesses left...")
      fmt.Printf("The word was: %v\n", word)
      fmt.Println("\033[91mYOU LOSE!") // red
      fmt.Println("\033[0m") // Text color reset
      break
    }
  }
}
