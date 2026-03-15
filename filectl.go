package main

import (
  "fmt"
  "os"
  )
  
func main() {
  
  // Check Arguments
  if len(os.Args) < 3 {
    fmt.Println("Usage:")
    fmt.Println("-write <filename> <text>")
    fmt.Println("-read <filename>")
    fmt.Println("-delete <filename>")
    return
  }
  
  // Arguments
  command := os.Args[1]
  filename := os.Args[2]
  
  // Condition Arguments
  switch command {
  case "-write": // Create & Write File
    if len(os.Args) < 4 {
      fmt.Println("Enter your text!")
      return
    }
    
    // Text File
    text := os.Args[3]
    
    file, err := os.Create(filename)
    if err != nil {
      fmt.Println("Error:", err)
      return
    }
    defer file.Close()
    
    _, err = file.WriteString(text)
    if err != nil {
      fmt.Println("Error:", err)
    }
    fmt.Println("Write file successfully!")
  
  case "-read": // Read File
    data, err := os.ReadFile(filename)
    if err != nil {
      fmt.Println("Error:", err)
      return
    }
    
    fmt.Println(string(data))
  
  case "-delete": // Delete File
    err := os.Remove(filename)
    if err != nil {
      fmt.Println("Error:", err)
    }
    
    fmt.Println("Delete file successfully!")
  
  default:
    fmt.Println("Command invalid!")
  }
  
}