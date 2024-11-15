package builders

import "fmt"
  

func NoCorrectLetters(goodLetters, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  for _, letter := range badLetters {

    badLetterString += letter
  }
  pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  
  // Create look aheads for each of the good letters
  for _, letter := range goodLetters {

    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  // Add the final "any-and-all letters" fill-in for the any remainder of the word
  pattern += ".*"

  return pattern

}


func FirstLetter(goodLetters []string, correctLetter string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add in the negative look ahead for the bad letters
  for _, letter := range badLetters {

    badLetterString += letter
  }
  pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 

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
func SecondLetter(goodLetters []string, correctLetter string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add negative look aheads
  for _, letter := range badLetters {

    badLetterString += letter
  }
  pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  
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
func ThirdLetter(goodLetters []string, correctLetter string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add negative look aheads
  for _, letter := range badLetters {

    badLetterString += letter
  }
  pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 

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
func FourthLetter(goodLetters []string, correctLetter string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add negative look aheads
  for _, letter := range badLetters {

    badLetterString += letter
  }
  pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 


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
func FifthLetter(goodLetters []string, correctLetter string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add negative look aheads
  for _, letter := range badLetters {

    badLetterString += letter
  }
  pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 

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

func FirstAndSecond(goodLetters []string, correctLetters []string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add negative look aheads
  for _, letter := range badLetters {

    badLetterString += letter
  }
  pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 
  

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

func FirstAndThird(goodLetters []string, correctLetters []string, badLetters []string) string {
  var badLetterString string

  pattern := "^"

  // Add negative look aheads
  for _, letter := range badLetters {

    badLetterString += letter
  }
  pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 

  
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

// START HERE 
//
//
func FirstAndFourth(goodLetters []string, correctLetters []string, badLetters []string) string {
  var badLetterString string

  pattern := "^"
  
  // Add negative look aheads
  for _, letter := range badLetters {

    badLetterString += letter
  }
  pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 

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

func FirstAndFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s.{3}%s", correctLetters[0], correctLetters[1])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func SecondAndThird(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^.{1}%s%s", correctLetters[0], correctLetters[1])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func SecondAndFourth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^.{1}%s.{1}%s", correctLetters[0], correctLetters[1])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func SecondAndFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^.{1}%s.{2}%s", correctLetters[0], correctLetters[1])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func ThirdAndFourth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^.{2}%s%s", correctLetters[0], correctLetters[1])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

//func ThirdAndFifth(goodLetters []string, correctLetters []string) string {
//  pattern := fmt.Sprintf("^.{2}%s.{1}%s", correctLetters[0], correctLetters[1])
//  
//  for _, letter := range goodLetters {
//    pattern += fmt.Sprintf("(?=.*[%s])", letter)
//  }
//
//  pattern += ".*"
//
//  return pattern
//
//}

func ThirdAndFifth(goodLetters []string, correctLetters []string, badLetters []string) string {
  // DEBUG: Not seeing bad letters
  fmt.Println("BAD LETTERS @ ThirdAndFifth()")
  for _, letter := range badLetters {
    fmt.Printf("%s\n", letter)
  }

  var badLetterString string

  pattern := "^"
  
  // Add negative look aheads
  for _, letter := range badLetters {
    badLetterString += letter
  }
  pattern += fmt.Sprintf("(?!.*[%s])", badLetterString) 

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

func FourthAndFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^.{3}%s%s", correctLetters[0], correctLetters[1])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}


// Triple Combo Regex Builders //

func FirstSecondThird(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func FirstSecondFourth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s%s.{1}%s", correctLetters[0], correctLetters[1], correctLetters[2])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func FirstSecondFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s%s.{2}%s", correctLetters[0], correctLetters[1], correctLetters[2])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func FirstThirdFourth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s.{1}%s%s", correctLetters[0], correctLetters[1], correctLetters[2])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func FirstThirdFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s.{1}%s.{1}%s", correctLetters[0], correctLetters[1], correctLetters[2])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func FirstFourthFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s.{2}%s%s", correctLetters[0], correctLetters[1], correctLetters[2])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func SecondThirdFourth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^.{1}%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func SecondThirdFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^.{1}%s%s.{1}%s", correctLetters[0], correctLetters[1], correctLetters[2])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func SecondFourthFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^.{1}%s.{1}%s%s", correctLetters[0], correctLetters[1], correctLetters[2])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

func ThirdFourthFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^.{2}%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern

}

// Quad Combo Regex Builders //

func FirstSecondThirdFourth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2], correctLetters[3])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern
}

func FirstSecondThirdFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s%s%s.{1}%s", correctLetters[0], correctLetters[1], correctLetters[2], correctLetters[3])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern
}

func FirstSecondFourthFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s%s.{1}%s%s", correctLetters[0], correctLetters[1], correctLetters[2], correctLetters[3])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern
}

func FirstThirdFourthFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^%s.{1}%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2], correctLetters[3])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern
}

func SecondThirdFourthFifth(goodLetters []string, correctLetters []string) string {
  pattern := fmt.Sprintf("^.{1}%s%s%s%s", correctLetters[0], correctLetters[1], correctLetters[2], correctLetters[3])
  
  for _, letter := range goodLetters {
    pattern += fmt.Sprintf("(?=.*[%s])", letter)
  }

  pattern += ".*"

  return pattern
}




















