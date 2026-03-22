package main

import (
  "fmt"
  "os"
  )
  
func main () {
  if file, err := os.Open("text.txt"); err != nil {
    fmt.Println("Error:", err)
    return
  }else {
    defer file.Close()
    fmt.Println("Open file successfully!")
  }
  
}