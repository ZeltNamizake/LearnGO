package lib

import (
  "fmt"
  "os"
  "os/exec"
  "bufio"
  )
  
func Ping() {
  scanner := bufio.NewScanner(os.Stdin)
  
  fmt.Print("Enter target: ")
  scanner.Scan()
  target := scanner.Text()
  
  cmd := exec.Command("ping", "-c", "4", target)
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    fmt.Println(err)
    return
  }
  
  stderr, err := cmd.StderrPipe()
  if err != nil {
	fmt.Println(err)
	return
  }
  
  if err = cmd.Start(); err != nil {
    fmt.Println(err)
    return
  }
  
  go func() {
  scannerErr := bufio.NewScanner(stderr)
	for scannerErr.Scan() {
		fmt.Println(scannerErr.Text())
   }
  }()
  
  scannerStdout := bufio.NewScanner(stdout)
  for scannerStdout.Scan() {
    line := scannerStdout.Text()
    fmt.Println(line)
  }
  
  if err := cmd.Wait(); err != nil {
    fmt.Println("Error:", err)
    return
  }
}
