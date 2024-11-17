package main

import (
  "fmt"
)

func main() {
  guessMap := map[int]string {
    1 : "f",
    2 : "l",
    5 : "e",
  }
  badString := []string {"b", "e", "s"}
  fmt.Println("BEFORE")
  fmt.Println(badString)
  
  lookup := make(map[string]struct{})

  for _, letter := range guessMap {
    lookup[letter] = struct{}{}
  }

  for i, letter := range badString {
    if _, exists := lookup[letter]; exists {
      fmt.Printf("Found the letter %s in the badString slice, at index %d!\n", letter, i)
      badString = append(badString[:i], badString[i+1:]...)
    }
  }  

  fmt.Println("AFTER:")
  fmt.Println(badString)
}


