package checkWord
import (
	"fmt"
	"slices"
	"strings"
)

// Function to check guess against selected word
func Check(word string, guess string) string {
  // Words split into slices for easier comparison
  splitWord := strings.Split(word, "")
  splitGuess := strings.Split(guess, "")
  var tempGuess []string
  var tempC string
  var guessed string

  for i, c := range splitGuess {
    // Color green for correct letter AND correct position
    if strings.ToLower(splitGuess[i]) == strings.ToLower(splitWord[i]) {
      upper := strings.ToUpper(splitGuess[i])
      tempC = fmt.Sprintf("\033[92m%v ", upper) // green
      // Color yellow for correct letter only
    } else if slices.Contains(splitWord, c) {
      upper := strings.ToUpper(splitGuess[i])
      tempC = fmt.Sprintf("\033[93m%v ", upper) // yellow
      // Color reset (white) for no match
    } else {
      upper := strings.ToUpper(splitGuess[i])
      tempC = fmt.Sprintf("\033[0m%v ", upper) // text color reset
    }
    tempGuess = append(tempGuess, tempC)
  }

  // Temporary slice of colored strings joined and returned
  guessed = strings.Join(tempGuess, "")
  return guessed
}
