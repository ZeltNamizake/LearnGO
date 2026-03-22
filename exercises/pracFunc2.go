package main

import (
  "fmt"
  "bufio"
  "strconv"
  "os"
  "errors"
  )

// Input User
func input() (x int, operator string, y int, err error) {
  // Get & Save input
  // .Scan = Save Value
  // .Text = Output Value
  scanner := bufio.NewScanner(os.Stdin)
  
  //Input X
  fmt.Print("Enter num (x): ")
  scanner.Scan()
  numX, errConv := strconv.Atoi(scanner.Text())
  if errConv != nil {
    err = errConv
    return
  }
  x = numX
  
  //Input Operator
  fmt.Print("Enter operator (+, -, *, /): ")
  scanner.Scan()
  operator = scanner.Text()
  
  //Input Y
  fmt.Print("Enter num (y): ")
  scanner.Scan()
  numY, errConv := strconv.Atoi(scanner.Text())
  if errConv != nil {
    err = errConv
    return
  }
  y = numY
  
  return x, operator, y, nil
  
}

// Function Calc
func calc(x int, operator string, y int) (int, error) {
  switch operator {
  case "+":
   return x + y, nil
    
  case "-":
   return x - y, nil
    
  case "*":
   return x * y, nil
    
  case "/":
    if y == 0 {
      // Output erros
      return 0, errors.New("invalid Value Y = 0")
    }
    return x / y, nil
  default:
      // Output erros
  return 0, errors.New("Operator invalid!")
  }
}

func main() {
  x, operator, y, err := input()
  if err != nil {
    fmt.Println(err)
    return
  }
  
  res, err := calc(x, operator, y)
  if err != nil {
    fmt.Println(err)
    return
  }
    fmt.Println("Result:", res)
  
}