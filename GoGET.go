package main

import (
  "fmt"
  "net/http"
  "net/url"
  "bufio"
  "os"
  "strings"
  )
  
func url_target() string {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Enter URL: ")
  scanner.Scan()
  input := strings.TrimSpace(scanner.Text())
  
  parsed, err := url.ParseRequestURI(input)
	if err != nil || !(parsed.Scheme == "http" || parsed.Scheme == "https") {
		fmt.Println("URL Invalid!")
		return ""
	}
	return input
}

func header_target() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(`
Select the header you want to display (case-insensitive):
- Content-Type
- Date
- Server
- Set-Cookie
- Content-Length
- Press Enter or type "all" to display all headers
`)
	fmt.Print("Enter Header Target: ")
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
  
func reqGet(url string, headers string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Status:", res.Status)

	header := strings.ToLower(headers)

	switch header {
	case "date":
		fmt.Println("Date:", res.Header.Get("Date"))
	case "content-type":
		fmt.Println("Content-Type:", res.Header.Get("Content-Type"))
	case "server":
		fmt.Println("Server:", res.Header.Get("Server"))
	case "set-cookie":
		fmt.Println("Set-Cookie:", res.Header.Get("Set-Cookie"))
	case "content-length":
		fmt.Println("Content-Length:", res.Header.Get("Content-Length"))
	case "", "all":
		fmt.Println("All Headers:")
		for key, values := range res.Header {
			for _, value := range values {
				fmt.Println(key, ":", value)
			}
		}
	default:
		value := res.Header.Get(headers)
		if value != "" {
			fmt.Printf("%s: %s\n", headers, value)
		} else {
			fmt.Println("Header not found:", headers)
		}
	}
}
  
func main () {
  urlTarget := url_target()
  if urlTarget == "" {
    return
  }
  
  headerTarget := header_target()
  reqGet(urlTarget, headerTarget)
}