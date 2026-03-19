package main

import (
  "fmt"
  "os"
  "toolKit-V1/lib"
  )
  
func Usage() {
  fmt.Println(`
Usage: go run main.go <command>
COMMAND:
  - ping
`)
}

func main() {
  if len(os.Args) < 2 {
    Usage()
    return
  }
  cmd := os.Args[1]
  
  switch cmd {
  case "ping":
   if len(os.Args) >= 3 {
     fmt.Println("Error")
     return
   }
    lib.Ping()
  
  default:
   fmt.Println("Invalid command!")
  }
}