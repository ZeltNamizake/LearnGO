package main
import(
   "fmt"
   "os"
   "bufio"
   "strconv"
  )

func stdinName() string {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Enter name: ")
  scanner.Scan()
  return scanner.Text()
}
func stdinBirthYear() string {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Enter birth year: ")
  scanner.Scan()
  return scanner.Text()
}

func main() {
  name := stdinName()
  birthYear := stdinBirthYear()
  yearsNow := 2026
  by, err := strconv.Atoi(birthYear)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  
  if by >= yearsNow || by < 1900{
    fmt.Println("Birth year invalid!")
    return
  }
  
  var ageNow int
  ageNow = yearsNow - by
  fmt.Println("Hello", name)
  fmt.Println("You are", ageNow, "years old")
}