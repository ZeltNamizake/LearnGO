package main

import (
  "fmt"
  "os"
  )
  
func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: go run writeFile.go <text>")
    return
  }
  
  text := []byte(os.Args[1])
  if err := os.WriteFile("text.txt", text, 0644); err != nil {
    fmt.Println("Error:", err)
    return
  }
  fmt.Println("Write file successfully")
}