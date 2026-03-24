package main

import (
	"bufio"
	"fmt"
	"os"
        "time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <wordlist>")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
                time.Sleep(1 * time.Second)
	}
}
