package main

import (
  "fmt"
  "os"
  )

func main() {
  
  if len(os.Args) < 3 {
    fmt.Println("Usage: rg main.go read")
    os.Exit(1)
  }
  
  cmd := os.Args[1]
  file := os.Args[2]
  if cmd == "read" && file != "" {
    data, err := os.ReadFile(file)
    if err != nil {
      fmt.Println("Error", err)
      return
    }
    fmt.Println(string(data))
  }else{
    fmt.Println("Invalid")
  }
}