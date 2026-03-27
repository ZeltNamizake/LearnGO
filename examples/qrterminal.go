package main

//Install & add library
// $ go get github.com/mdp/qrterminal/v3

import (
  "os"
  "github.com/mdp/qrterminal/v3"
  "fmt"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: go run qrterminal.go <text>")
    return
  }
  config := qrterminal.Config{
      Level: qrterminal.L,
      Writer: os.Stdout,
      BlackChar: qrterminal.BLACK,
      WhiteChar: qrterminal.WHITE,
      QuietZone: 1,
  }
  qrterminal.GenerateWithConfig(os.Args[1], config)
}