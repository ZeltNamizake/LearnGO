package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  )

/*func checkNumber(num int) string {
  if num > 0 {
    return "Positive"
  } else if num < 0 {
    return "Negative"
  } else {
    return "Zero"
  }
}*/

/*func divide(a int, b int) (res int, err error) {
  if b == 0 {
    err = errors.New("Invalid value 0")
    return
  }
  res = a / b
  return
}*/

/*func getMinMax(a int, b int) (int, int) {
  if a < b {
    return a, b
  }
  return b, a
}*/

/*func checkEventOdd(num int) string {
  if num % 2 == 0 {
    return "Genap"
  }
  return "Ganjil"
}*/

// Input Number
func inputNum() (number int, err error) {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Enter text: ")
  scanner.Scan()
  str := scanner.Text()
  num, errConv := strconv.Atoi(str)
  if errConv != nil {
    err = errConv
    return
  }
  number = num
  return
}

// Check Even or Odd
func checkEvenOdd(num int) string {
  if num % 2 == 0 {
    return "Even"
  }
  return "Odd"
}


func main() {
  if num, err := inputNum(); err != nil {
    fmt.Println(err)
  }else {
  fmt.Println(checkEvenOdd(num))
  }
}