/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "os"
  "bufio"
  "github.com/dlclark/regexp2"
	"github.com/spf13/cobra"
  "WordleGuesser/cmd/builders"
)

var badLetters []string

// guessCmd represents the guess command
var guessCmd = &cobra.Command{
	Use:   "guess",
	Short: "Prompt user for a guess",
	Run: func(cmd *cobra.Command, args []string) {

    for {

      // Prompt for word
      var word string           // To hold the current guess 
      var wordSlice []string     // Slice for parsed guess
      fmt.Print("Enter word: ")
      fmt.Scanln(&word)

      // Prompt for validation
      var validation string
      var validationSlice []string
      fmt.Print("Enter validation: ")
      fmt.Scanln(&validation)

      // Parsing Word to add to slice
      for _, letter := range word {
        // fmt.Printf("%c\n", letter)
        wordSlice = append(wordSlice, string(letter))
      }

      // Parsing Validation to add to slice
      for _, num := range validation {
        validationSlice = append(validationSlice, string(num)) 
      }

      analyzeGuess(wordSlice, validationSlice)
    }
	},
}


func analyzeGuess(word, validation []string) {
  // Maps for assigning letters and their corresponding positions
  goodLetters := make(map[int]string)
  correctLetters := make(map[string]int)

  // Sorting letters by the associated validation ids
  var x = 0
  for x < len(word) {
    if validation[x] == "0" {
      badLetters = append(badLetters, word[x])
    } else if validation[x] == "1" {
      goodLetters[x+1] = word[x]
    } else if validation[x] == "2" {
      correctLetters[word[x]] = x+1
    } else {
      fmt.Printf("Unknown validation id: %s", validation[x])
    }
    x++
  }

  // Check for and remove any 'good letters' from the 'bad letters' set
  lookup := make(map[string]struct{})
  for _, letter := range goodLetters {
    lookup[letter] = struct{}{} 
  }

  // Check for and remove any 'correct letters' from the 'bad letters' set
  lookup2 := make(map[string]struct{})
  for cl, _ := range correctLetters {
    lookup2[cl] = struct{}{}
  }

  for index, bl := range badLetters {
    if _, exists := lookup[bl]; exists {
      badLetters = append(badLetters[:index], badLetters[index+1:]...)
    } else if _, exists := lookup2[bl]; exists {
      badLetters = append(badLetters[:index], badLetters[index+1:]...)
    }
  }



  // Create regex pattern from the series of functions beginning with 'wordBuilder'
  pattern := wordBuilder(goodLetters, correctLetters, badLetters)

// DEBUG: Print pattern
//  fmt.Println(pattern)

  // DEBUG: To read pattern before wordlist is generated
//  bufio.NewReader(os.Stdin).ReadBytes('\n')


  checkForWords(pattern)

}

// Function to check for correct letters, if so determine the position of the correct letters
// and assign the correct builder function based on position.
func wordBuilder(goodLetters map[int]string, correctLetters map[string]int, badLetters []string) string {
  var pattern string

  // Check for any correct letters and check the position based on the number of correct letters
  _, length := checkForCorrectLetters(correctLetters)
  
  if length > 0 {
    switch length {
    case 1:
      fmt.Println("1 correct letter has been guessed.")
      pattern = checkPosition_OneLetter(correctLetters, goodLetters, badLetters)
    case 2:
      fmt.Println("2 correct letters have been guessed.")
      pattern = checkPosition_TwoLetters(correctLetters, goodLetters, badLetters)
    case 3:
      fmt.Println("3 correct letters have been guessed.")
       pattern = checkPosition_ThreeLetters(correctLetters, goodLetters, badLetters)
    case 4:
      fmt.Println("4 correct letters have been guessed.")
      pattern = checkPosition_FourLetters(correctLetters, goodLetters, badLetters)
    default:
      // At this point the Wordle should be solvable
      fmt.Println("You solved the Wordle, why are you still here?")
    }
  } else {
    fmt.Println("No correct letters have been guessed.")
    pattern = builders.NoCorrectLetters(goodLetters, badLetters)
  }
  return pattern
}

// The following 4 functions are to parse the position of the correct letters and assign them accordingly
//
// Determine the position of the correct letter and assign the correct regexbuilder function
func checkPosition_OneLetter(correctLetters map[string]int, goodLetters map[int]string, badLetters []string) string {
  var pattern string
  for letter, position := range correctLetters {
    switch position {
    case 1:
      pattern = builders.FirstLetter(goodLetters, letter, badLetters)
    case 2:
      pattern = builders.SecondLetter(goodLetters, letter, badLetters)
    case 3:
      pattern = builders.ThirdLetter(goodLetters, letter, badLetters)
    case 4:
      pattern = builders.FourthLetter(goodLetters, letter, badLetters)
    case 5:
      pattern = builders.FifthLetter(goodLetters, letter, badLetters)
    }
  }
  return pattern
}

func checkPosition_TwoLetters(correctLetters map[string]int, goodLetters map[int]string, badLetters []string) string {
  for _, letter := range badLetters {
    fmt.Printf("%s\n", letter)
  }
  //


  var pattern string
  var letterPositions []int
  var lettersForRegex []string

  for letter, position := range correctLetters {
    letterPositions = append(letterPositions, position) 
    lettersForRegex = append(lettersForRegex, letter)
  }

  switch {
  case matches(letterPositions, []int{1,2}):
    pattern = builders.FirstAndSecond(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,3}):
    pattern = builders.FirstAndThird(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,4}):
    pattern = builders.FirstAndFourth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,5}):
    pattern = builders.FirstAndFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{2,3}):
    pattern = builders.SecondAndThird(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{2,4}):
    pattern = builders.SecondAndFourth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{2,5}):
    pattern = builders.SecondAndFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{3,4}):
    pattern = builders.ThirdAndFourth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{3,5}):
    pattern = builders.ThirdAndFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{4,5}):
    pattern = builders.FourthAndFifth(goodLetters, lettersForRegex, badLetters)
  }
  return pattern
}


func checkPosition_ThreeLetters(correctLetters map[string]int, goodLetters map[int]string, badLetters []string) string {
  var pattern string
  var letterPositions []int
  var lettersForRegex []string

  for letter, position := range correctLetters {
    letterPositions = append(letterPositions, position)
    lettersForRegex = append(lettersForRegex, letter)
  }

  switch {
  case matches(letterPositions, []int{1,2,3}):
    pattern = builders.FirstSecondThird(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,2,4}):
    pattern = builders.FirstSecondFourth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,2,5}):
    pattern = builders.FirstSecondFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,3,4}):
    pattern = builders.FirstThirdFourth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,3,5}):
    pattern = builders.FirstThirdFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,4,5}):
    pattern = builders.FirstFourthFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{2,3,4}):
    pattern = builders.SecondThirdFourth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{2,3,5}):
    pattern = builders.SecondThirdFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{2,4,5}):
    pattern = builders.SecondFourthFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{3,4,5}):
    pattern = builders.ThirdFourthFifth(goodLetters, lettersForRegex, badLetters)
  }
  return pattern 
}


func checkPosition_FourLetters(correctLetters map[string]int, goodLetters map[int]string, badLetters []string) string {
  var pattern string
  var letterPositions []int
  var lettersForRegex []string

  for letter, position := range correctLetters {
    letterPositions = append(letterPositions, position)
    lettersForRegex = append(lettersForRegex, letter)
  }

  switch {
  case matches(letterPositions, []int{1,2,3,4}):
    pattern = builders.FirstSecondThirdFourth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,2,3,5}):
    pattern = builders.FirstSecondThirdFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,2,4,5}):
    pattern = builders.FirstSecondFourthFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{1,3,4,5}):
    pattern = builders.FirstThirdFourthFifth(goodLetters, lettersForRegex, badLetters)
  case matches(letterPositions, []int{2,3,4,5}):
    pattern = builders.SecondThirdFourthFifth(goodLetters, lettersForRegex, badLetters)
  }
  return pattern
}


// Determine whether or not any correct letters have been guessed and stored in the map
func checkForCorrectLetters(correctLetters map[string]int) (bool, int){
  if len(correctLetters) == 0 {
    return false, 0
  } else {
    return true, len(correctLetters)
  }
}

// Helper function for the checkPostion_* functions to determine the correct combination of letter positions  
func matches(slice1, slice2 []int) bool {
  for index, value := range slice1 {
    if value != slice2[index] {
      return false
    }
  }
  return true
}


// Check the list of words using the given regex query for filtering
func checkForWords(pattern string) {

  query := regexp2.MustCompile(pattern, regexp2.None)

  file, err := os.Open("../availableWords.txt")
    if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()    // Closing file when function completes

  // Creating scanner to read file line-by-line
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {

    line := scanner.Text()

    match, _ := query.MatchString(line) 
    if match {
      fmt.Println(line)
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
      fmt.Println("Error occurred while reading file:", err)
    }
  }
}

func init() {
  rootCmd.AddCommand(guessCmd)
}





















