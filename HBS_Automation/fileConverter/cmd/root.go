// Implement concurrency


/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
  "os/exec"
  "fmt"
  "encoding/csv"
  "path/filepath"
  "strings"
  "github.com/360EntSecGroup-Skylar/excelize"
  "github.com/fatih/color"
  "sync"
  //DEBUG
//  "bufio"

	"github.com/spf13/cobra"
)

// Declare global variables
var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var pauseMutex sync.Mutex
var pauseCond = sync.NewCond(&pauseMutex)
var processingPaused bool
var outputChannel = make(chan string, 100) // Buffered channel for goroutine messages

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fileConverter",
	Short: `
 
                 / ____|                        | |           
   ___ _____   _| |     ___  _ ____   _____ _ __| |_ ___ _ __ 
  / __/ __\ \ / / |    / _ \| '_ \ \ / / _ \ '__| __/ _ \ '__|
 | (__\__ \\ V /| |___| (_) | | | \ V /  __/ |  | ||  __/ |   
  \___|___/ \_/  \_____\___/|_| |_|\_/ \___|_|   \__\___|_|   
                                                              
                                                              
  Rapidly convert .CSV files to .XSLX
  -----------------------------------
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

  // Clear screen on initial execution
  // cmd := exec.Command("clear")
  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

  // Establish variables
  var csvDirectory string
  var xlsxDirectory string 

  // Prompt for CSV file directory
  fmt.Print("Enter directory path of .csv files: ")
  fmt.Scanln(&csvDirectory)

  // Check if given directory exists
  if _, err := os.Stat(csvDirectory); os.IsNotExist(err) {
    fmt.Printf("Error: %s does not exist.\n", red(csvDirectory))
    return
  }

  // Prompt for XLSX output directory
  fmt.Print("\nEnter destination directory for .xlsx files: ")
  fmt.Scanln(&xlsxDirectory)
  
  // Check if given directory exists
  if _, err := os.Stat(xlsxDirectory); os.IsNotExist(err) {
    fmt.Printf("Error: %s does not exist.\n", red(xlsxDirectory))
    return
  }

  // Create space
  fmt.Println("\n")


  // Read CSV files
  csvFiles, err := filepath.Glob(filepath.Join(csvDirectory, "*.csv"))
  if err != nil {
    fmt.Printf("Error reading CSV files: %v\n", err)
    return
  }

  // Check if any CSV files were discovered
  if len(csvFiles) == 0 {
    fmt.Printf("Error: No .csv files found in %s\n", red(csvDirectory))
    return
  }

  // Create a WaitGroup to track output processing
  var outputWg sync.WaitGroup
  outputWg.Add(1)

  // Process output from the outputChannel
  go func() {
    defer outputWg.Done() // Mark as done when this goroutine finishes
    for msg := range outputChannel {
      pauseMutex.Lock()
      for processingPaused {
        pauseCond.Wait() // Block output while waiting for user input
      }
      pauseMutex.Unlock()
      fmt.Print(msg) // Print the message to stdout
    }
  }()

  // Setup a WaitGroup and buffered channel to limit concurrent workers for the file conversion process
  var wg sync.WaitGroup
  maxWorkers := 4
  sem := make(chan struct{}, maxWorkers)

  // Loop over CSV files and pass them to the 'convertCSVtoXLSX()' function for conversion
  for _, csvFile := range csvFiles {
    baseFileName := filepath.Base(csvFile)
    wg.Add(1)

    go func(csvFile string, baseFileName string) {
      defer wg.Done()

      // Acquire for semaphore slot
      sem <- struct{}{}
      defer func() { <-sem }()

      // Wait if paused
      pauseMutex.Lock()
      for processingPaused {
        pauseCond.Wait()
      }
      pauseMutex.Unlock()

      // Convert .csv to .xlsx, and output the results
      outputFile, err := convertCSVtoXLSX(csvFile, xlsxDirectory)
      if err != nil {
        outputChannel <- fmt.Sprintf("Error converting file: %s::%v\n", csvFile, err)
      } else if outputFile != "" {
        outputChannel <- fmt.Sprintf("%s :>>: %s -> %s\n", baseFileName, outputFile, green("COMPLETED"))
      }
    }(csvFile, baseFileName)
  }

  // Wait for all GoRoutines to complete
  wg.Wait()
  close(outputChannel) //Signal the output goroutine to finish

  // Wait for the output processing goroutine to complete
  outputWg.Wait()

  fmt.Println("\nFinished!")

}

// File conversion function
func convertCSVtoXLSX(csvFilePath, xlsxDirectory string) (string, error) {
  // Open the CSV File
  csvFile, err := os.Open(csvFilePath)
  if err != nil {
    fmt.Println("Failed to open file:")
    return "", err
  }
  defer csvFile.Close()
  
  // Read CSV Data
  reader := csv.NewReader(csvFile)
  reader.LazyQuotes = true
  records, err := reader.ReadAll()
  if err != nil {
    fmt.Println("Failed to read file:")
    return "", err
  }

  // Create the new XLXS file
  xlsx := excelize.NewFile()

  // Write data to the XLSX file
  for rowIndex, row := range records {
    for colIndex, value := range row {
      cell := excelize.ToAlphaString(colIndex) + fmt.Sprintf("%d", rowIndex+1)
      xlsx.SetCellValue("Sheet1", cell, value)
    }
  }

  // Isolate file name without path or extension
  baseFileName := filepath.Base(csvFilePath)
  fileName := strings.TrimSuffix(baseFileName, filepath.Ext(baseFileName))

  // Create full path for the new XLSX file and XSLX file name return value
  xlsxOutputPath := filepath.Join(xlsxDirectory, fileName + ".xlsx")
  xlsxFileName := filepath.Join(fileName + ".xlsx")

  // Check if XLSX file already exists in path
  // Establish variable for if a duplicate file is detected

  if _, err := os.Stat(xlsxOutputPath); err == nil {
    var overwriteCheck string

    pauseMutex.Lock() // Lock mutex to block all new future goroutines
    processingPaused = true // Pause all current goroutines
    fmt.Printf("\nWARNING: The file %s already exists. Do you want to overwrite it? (y/n): ", yellow(xlsxFileName))
    fmt.Scanln(&overwriteCheck)
    fmt.Println("") // White space for formatting
    overwriteCheck = strings.ToLower(strings.TrimSpace(overwriteCheck))

    // Resume all other goroutines
    processingPaused = false
    pauseMutex.Unlock() // Release the mutex
    pauseCond.Broadcast()

    if overwriteCheck != "y" {
      fmt.Printf("\nSkipping file: %s\n", fileName)
      // White Space
      fmt.Println("")
      

      return "", err
    }
  }

  // Save the XLSX file
  err = xlsx.SaveAs(xlsxOutputPath)
  if err != nil {
    fmt.Println("Failed to save file:")
    return "", err
  }

  return xlsxFileName, nil
}


// Define Flags and Configuration Settings
func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

