package checkGuess

import (
  "fmt"
  "unicode"
)

func Check(guess string) (bool) {
  for _, l := range guess {
    if !unicode.IsLetter(l) {
      fmt.Println("\033[91mInvalid characters")
      fmt.Println("\033[0m")
      return false
    }
  }

  if len(guess) > 5 {
    fmt.Println("\033[91mToo many letters")
    fmt.Println("\033[0m")
    return false
  } else if len(guess) < 5 {
    fmt.Println("\033[91mToo few letters")
    fmt.Println("\033[0m")
    return false
  }

  return true
}
