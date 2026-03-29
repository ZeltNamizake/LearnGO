package main

import (
  "fmt"
  "time"
  "net"
  "os"
  "sync"
)


func portScan(host string, port int) bool {
  address := fmt.Sprintf("%s:%d", host, port)
  conn, err := net.DialTimeout("tcp", address, 1 * time.Second)
  if err != nil {
    return false
  }
  conn.Close()
  return true
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: go run portScan.go <host>")
    return
  }
  
commonPorts := []int{
  21,
  22,
  23,
  25,
  53,
  80,
  110,
  143,
  443,
  1080,
  3306,
  8080,
}
  
  host := os.Args[1]
  var wg sync.WaitGroup
  sem := make(chan struct{}, 5)
  
  for _, p := range commonPorts {
    wg.Add(1)
    sem <- struct{}{}
    
    go func(port int) {
      defer wg.Done()
      defer func() { <- sem }()
      
      if portScan(host, port) {
       fmt.Printf("[OPEN] %d\n", port)
      } else {
        fmt.Printf("[CLOSED] %d\n", port)
      }
      
    }(p)
  }
    wg.Wait()
}