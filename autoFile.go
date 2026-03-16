package main

import(
 "os"
 "bufio"
 "fmt"
)

func main() {
  
  if len(os.Args) < 2 {
    fmt.Println("Usage: go run autoFile.go <text>")
    return
  }
  
  text := []byte(os.Args[1])
  if err := os.WriteFile("data.txt", text, 0644); err != nil {fmt.Println("Error write:", err); return}
  
  file, err := os.Open("data.txt")
  if err != nil {fmt.Println("Error open:", err); return}
  defer file.Close(); fmt.Println("Open file successfully!")
  
  scanner := bufio.NewScanner(file); fmt.Println("Output:")
  for scanner.Scan() {fmt.Println(scanner.Text())}
  if err := scanner.Err(); err != nil {fmt.Println("Error read:", err); return}
  
}