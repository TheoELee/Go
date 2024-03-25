package main

import (
	"fmt"
	"os"
)

func checkForError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := os.ReadFile("./test.txt")
	checkForError(err)
	fmt.Print(data)

	file, err := os.Open("/text.txt")
	checkForError(err)
	fmt.Print(file)
}