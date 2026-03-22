package main
import (
  "fmt"
  )
  
func main() {
  
  // Normal Loop
  fmt.Println("Exercise - 1")
  for i := 1; i <= 5; i++ {
    fmt.Println(i)
  }
  
  // Decline
  fmt.Println("\nExercise - 2")
  for i := 5; i >= 1; i-- {
    fmt.Println(i)
  }
  
  // Jump +2
  fmt.Println("\Exercise - 3")
  for i := 2; i <= 10; i +=2{
    fmt.Println(i)
  }
  
  // Jump -2
  fmt.Println("Exercise - 4")
  for i := 10; i >= 2; i -=2{
    fmt.Println(i)
  }
}