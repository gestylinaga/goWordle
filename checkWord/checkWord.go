package checkWord

import (
	"fmt"
	"slices"
	"strings"
)

func Check(word string, guess string) string {
  splitWord := strings.Split(word, "")
  splitGuess := strings.Split(guess, "")
  var tempGuess []string
  var tempC string
  var guessed string

  for i, c := range splitGuess {
    if strings.ToLower(splitGuess[i]) == strings.ToLower(splitWord[i]) {
      upper := strings.ToUpper(splitGuess[i])
      tempC = fmt.Sprintf("\033[92m%v ", upper)
      tempGuess = append(tempGuess, tempC)
    } else if slices.Contains(splitWord, c) {
      upper := strings.ToUpper(splitGuess[i])
      tempC = fmt.Sprintf("\033[93m%v ", upper)
      tempGuess = append(tempGuess, tempC)
    } else {
      upper := strings.ToUpper(splitGuess[i])
      tempC = fmt.Sprintf("\033[0m%v ", upper)
      tempGuess = append(tempGuess, tempC)
    }
  }
  //guessed = append(guessed, strings.Join(tempGuess, ""))
  guessed = strings.Join(tempGuess, "")

  return guessed
}
