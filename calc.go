package main

import(
  "fmt"
  "os"
  "strconv"
  )
  
func main() {
 
 if len(os.Args) < 4 {
   fmt.Println("Usage: <num> <operator> <num>")
   return
 }
 
 value1 := os.Args[1]
 operator := os.Args[2]
 value2 := os.Args[3]
 
 num1, err := strconv.Atoi(value1)
 if err != nil {
   fmt.Println("Error:", err)
   return
 }
 
 num2, err := strconv.Atoi(value2)
 if err != nil {
   fmt.Println("Error:", err)
   return
 }
 
 switch operator {
 case "+" :
   fmt.Println(num1 + num2)
   
 case "-" :
   fmt.Println(num1 - num2)
   
 case "*" :
   fmt.Println(num1 * num2)
   
 case "/" :
   if num2 == 0 {
     fmt.Println("Tidak bisa membagi dengan nol!")
     return
   }
   
   fmt.Println(num1 / num2)
  
 default:
  fmt.Println("Operator Invalid")
 }

}