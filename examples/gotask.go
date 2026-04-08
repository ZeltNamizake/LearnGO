package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Task struct {
	Title string
	Done  bool
}

func inputTitle() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Title task: ")
	scanner.Scan()
	return scanner.Text()
}

func saveTasks(filename string, tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "")
	_ = ioutil.WriteFile(filename, data, 0644)
}

func main() {
	tasks := []Task{}
	title := inputTitle()

	tasks = append(tasks, Task{Title: title, Done: false})
	saveTasks("tasks.json", tasks)

	fmt.Println("Task successfully save!")
}
