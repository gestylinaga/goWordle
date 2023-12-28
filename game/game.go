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

func displayResults(guesses int, guessedList []string) {
  fmt.Printf("Guesses left: %d\n", guesses)
  fmt.Println("Past guesses:")
  for i := 0; i < len(guessedList); i++{
    fmt.Println(guessedList[i])
    fmt.Printf("\033[0m")
  }
  fmt.Println("\033[0m")
}

func main() {
  fmt.Println("Wordle written in Go! goWordle!")
  fmt.Println("\nGuess the 5 letter word, in 6 tries!")
  fmt.Println("\033[93mLetters in Yellow are correct, BUT in the wrong place")
  fmt.Println("\033[92mLetters in Green are correct, AND in the correct place")
  fmt.Println("\033[0m")

  word := wordChoice.Word()
  guesses := 6
  var guessed string
  var guessedList []string

  for {
    fmt.Println("Your guess?")
    fmt.Println("\033[0m")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    err := scanner.Err()
    guess := scanner.Text()

    if err != nil {
      log.Fatal(err)
    }

    if checkGuess.Check(guess) {
      guesses -= 1
      guessed = checkWord.Check(word, guess)
      guessedList = append(guessedList, guessed)

    }
    displayResults(guesses, guessedList)

    // Win condition check
    if strings.ToLower(guess) == strings.ToLower(word) {
      fmt.Printf("Correct! The word was: %v\n", word)
      fmt.Println("\033[92mYOU WIN!") // green
      fmt.Println("\033[0m")
      break
    }

    // Lose condition check
    if (guesses == 0) && (strings.ToLower(guess) != strings.ToLower(word)) {
      fmt.Println("No guesses left...")
      fmt.Printf("The word was: %v\n", word)
      fmt.Println("\033[91mYOU LOSE!") // red
      fmt.Println("\033[0m")
      break
    }
  }
}
