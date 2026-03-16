package main

import ("fmt"; "bufio"; "os")

func input() string {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Enter text: ")
  scanner.Scan()
  return scanner.Text()
}

func main() {
  text := input()
  fmt.Println("Output:", text)
}