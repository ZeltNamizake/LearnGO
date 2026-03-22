package main

import (
  "fmt"
  "os"
  "bufio"
)
  
func main() {
  file, err := os.Open("file.txt")
  if err != nil {fmt.Println("Error:", err); return}
  defer file.Close(); fmt.Println("Open file successfully!\n")
  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {fmt.Println(scanner.Text())}
  
  if err := scanner.Err(); err != nil {fmt.Println("Error:", err)}
  
}