package main

import (
  "fmt"
  "os"
  )

func main() {
  
  if len(os.Args) < 4 {
    fmt.Println("Usage: rg main.go create <filename> 'text' ")
    os.Exit(1)
  }
  
  cmd := os.Args[1]
  file := os.Args[2]
  text := os.Args[3]
  
  if cmd == "create" && file != "" && text != "" {
    file, err := os.Create(file)
    if err != nil {
      fmt.Println("Error", err)
      return
    }
    defer file.Close()
    
    _, err = file.WriteString(text)
    if err != nil {
      fmt.Println("Error:", err)
      return
    }
    fmt.Println("File created successfully!")
  }else{
    fmt.Println("Invalid")
  }
}