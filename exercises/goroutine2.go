package main

import (
 "fmt"
 "net/http"
 "bufio"
 "os"
 "sync"
 "time"
)

func checkStatusCode(domain string) {
  res, err := http.Get("https://" + domain)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  defer res.Body.Close()
  fmt.Printf("https://%s - Status code: %s\n", domain, res.Status)
  
}

func urlWordlist() {
  file, err := os.Open("domain.txt")
  if err != nil {
    fmt.Println("Error:", err)
  }
  defer file.Close()
  
  var wg sync.WaitGroup // wait for all goroutine
  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()

    wg.Add(1) // add task

    go func(t string) {
        defer wg.Done() // mark done

        fmt.Println("Start", t) // start log
        time.Sleep(2 * time.Second) // simulate delay

        checkStatusCode(t) // run task

        fmt.Println("End", t) // end log

    }(text) // pass value to avoid capture bug
}

wg.Wait() // wait all goroutines

}

func main() {
  urlWordlist()
}