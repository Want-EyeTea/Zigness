package builders

import "fmt"
  

func NoCorrectLetters(goodLetters map[int]string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Create look aheads for each of the good letters
  for _, letter := range goodLetters {

    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern

}


func FirstLetter(goodLetters map[int]string, correctLetter string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Create look aheads for each of the good letters
  for _, letter := range goodLetters {

    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters to there corresponding position(s)
  pattern += fmt.Sprintf("%s.{4}", correctLetter)


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
} 

// If the correct letter is in the second position
func SecondLetter(goodLetters map[int]string, correctLetter string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }
  
  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letter in the corresponding position
  pattern += fmt.Sprintf(".{1}%s", correctLetter)
  
  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
} 


// If the correct letter is in the third position
func ThirdLetter(goodLetters map[int]string, correctLetter string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letter in the corresponding position
  pattern += fmt.Sprintf(".{2}%s", correctLetter)

  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
} 


// If the correct letter is in the fourth position
func FourthLetter(goodLetters map[int]string, correctLetter string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letter in the corresponding position
  pattern += fmt.Sprintf("^.{3}%s", correctLetter)

  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
} 


// If the correct letter is in the fifth position
func FifthLetter(goodLetters map[int]string, correctLetter string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letter in the corresponding position
  pattern += fmt.Sprintf("^.{4}%s", correctLetter)

  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
} 


// Regex builders for combos of 2 //

func FirstAndSecond(goodLetters map[int]string, correctLetters []string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }
  
  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("^%s%s.{3}", correctLetters[0], correctLetters[1])

  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FirstAndThird(goodLetters map[int]string, correctLetters []string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }
  
  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("^%s.{1}%s", correctLetters[0], correctLetters[1])

  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FirstAndFourth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {
  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("^%s.{2}%s", correctLetters[0], correctLetters[1])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern

}

func FirstAndFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {
  
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("^%s.{3}%s", correctLetters[0], correctLetters[1])

  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern

}

func SecondAndThird(goodLetters map[int]string, correctLetters []string, badLetters []string) string {
  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{1}%s%s", correctLetters[0], correctLetters[1])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern

}

func SecondAndFourth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {
  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{1}%s.{1}%s", correctLetters[0], correctLetters[1])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern

}

func SecondAndFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {
  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{1}%s.{2}%s", correctLetters[0], correctLetters[1])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern

}

func ThirdAndFourth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{2}%s%s", correctLetters[0], correctLetters[1])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}


func ThirdAndFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{2}%s.{1}%s", correctLetters[0], correctLetters[1])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FourthAndFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{3}%s%s", correctLetters[0], correctLetters[1])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}



// Triple Combo Regex Builders //


func FirstSecondThird(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FirstSecondFourth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("%s%s.{1}%s", correctLetters[0], correctLetters[1], correctLetters[2])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FirstSecondFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("%s%s.{2}%s", correctLetters[0], correctLetters[1], correctLetters[2])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FirstThirdFourth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("%s.{1}%s%s", correctLetters[0], correctLetters[1], correctLetters[2])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FirstThirdFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("%s.{1}%s.{1}%s", correctLetters[0], correctLetters[1], correctLetters[2])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FirstFourthFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("%s.{2}%s%s", correctLetters[0], correctLetters[1], correctLetters[2])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func SecondThirdFourth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{1}%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func SecondThirdFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{1}%s%s.{1}%s", correctLetters[0], correctLetters[1], correctLetters[2])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func SecondFourthFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{1}%s.{1}%s%s", correctLetters[0], correctLetters[1], correctLetters[2])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func ThirdFourthFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{2}%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}


// Quad Combo Regex Builders //

func FirstSecondThirdFourth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("%s%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2], correctLetters[3])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FirstSecondThirdFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("%s%s%s.{1}%s", correctLetters[0], correctLetters[1], correctLetters[2], correctLetters[3])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FirstSecondFourthFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("%s%s.{1}%s%s", correctLetters[0], correctLetters[1], correctLetters[2], correctLetters[3])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func FirstThirdFourthFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf("%s.{1}%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2], correctLetters[3])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}

func SecondThirdFourthFifth(goodLetters map[int]string, correctLetters []string, badLetters []string) string {

  var badLetterString string

  pattern := "^"
  
  // Add in the negative look ahead for the bad letters
  if len(badLetters) > 0 {

    for _, letter := range badLetters {

      badLetterString += letter
    }
    pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  }

  // Create negative look aheads for each of the 'wrong positions' of the good letters
  for position, letter := range goodLetters {
    if position == 1 {
      pattern += fmt.Sprintf("(?![%s])", letter)
    } else {
      pattern += fmt.Sprintf("(?!.{%d}%s)", position - 1, letter)
    }
  }

  // Add look aheads for each of the available 'good letters'
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Set the correct letters in the corresponding position
  pattern += fmt.Sprintf(".{1}%s%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2], correctLetters[3])


  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern
}



















