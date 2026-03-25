package main

import (
  "fmt"
  "sync"
  )
  
func program1(start int, end int, wg *sync.WaitGroup) {
  defer wg.Done()
  for i := start; i <= end; i++ {
    fmt.Println(i)
  }
}
func program2(start int , end int, wg *sync.WaitGroup) {
  defer wg.Done()
  for i := start; i <= end; i++ {
    fmt.Println(i)
  }
}
func program3(start int , end int, wg *sync.WaitGroup) {
  defer wg.Done()
  for i := start; i <= end; i++ {
    fmt.Println(i)
  }
}

func main() {
  var wg sync.WaitGroup
  
  wg.Add(3) // Add 3 Workers for run 3 function
  go program1(1, 3, &wg)
  go program2(3, 6, &wg)
  go program3(6, 9, &wg)
  wg.Wait() // wait workers to finish.
  
  fmt.Println("Done!")
  
}