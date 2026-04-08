package main

import "fmt"

type Task struct {
	Title string
	Done  bool
}

func markDone(t *Task) {
	t.Done = true
}

func main() {

	/*	t := Task{
			Title: "Learning GO",
			Done:  false,
		}
		println(t.Title)*/

	tasks := []Task{
		{Title: "Learning GO", Done: false},
		{Title: "Reading", Done: false},
		{Title: "Sleep", Done: false},
	}

	markDone(&tasks[0])

	for i, t := range tasks {
		fmt.Printf("%d. %s (Selesai: %v)\n", i+1, t.Title, t.Done)
	}
}
