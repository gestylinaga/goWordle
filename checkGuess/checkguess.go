package checkGuess
import (
  "fmt"
  "unicode"
)

// Function to check validity of input,
// also prints feedback message in red text
func Check(guess string) (bool) {
  // Checks if all characters are letters
  for _, l := range guess {
    if !unicode.IsLetter(l) {
      fmt.Println("\033[91mInvalid characters")
      fmt.Println("\033[0m") // Text color reset
      return false
    }
  }

  // Checks for proper string length
  if len(guess) > 5 {
    fmt.Println("\033[91mToo many letters")
    fmt.Println("\033[0m") // Text color reset
    return false
  } else if len(guess) < 5 {
    fmt.Println("\033[91mToo few letters")
    fmt.Println("\033[0m") // Text color reset
    return false
  }

  return true
}
