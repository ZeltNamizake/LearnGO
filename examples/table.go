package main

import (
	"github.com/fatih/color"   // go get github.com/fatih/color
	"github.com/rodaine/table" // go get github.com/rodaine/table
)

func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("#", "Fruit", "Value")

	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	tbl.AddRow(1, "Grape", 100)
	tbl.AddRow(2, "Banana", 75)
	tbl.AddRow(3, "Watermelon", 50)

	tbl.Print()
}
