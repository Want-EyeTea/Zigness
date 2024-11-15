package main

import (
  "fmt"
  "github.com/dlclark/regexp2"
  "os"
  "bufio"
)

func main() {
  filepath := "../availableWords.txt"
  //pattern := `^(?=.*[l])(?=.*[a])t.{4}(?!.*[es]).*`
  pattern := `^(?!.*[es])(?=.*[l])(?=.*[a])t.{4}.*`

  query := regexp2.MustCompile(pattern, regexp2.None)

  fmt.Printf("Current Query: %s\n", query)


  file , err := os.Open(filepath)
  if err != nil {
    fmt.Println("Could not open file: ", err)
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    
    match, _ := query.MatchString(line)
    if match {
      fmt.Println(line)
    }
  }


  
}


