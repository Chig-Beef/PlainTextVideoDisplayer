package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Expected file to view as arg")
		return
	}

	fileName := args[1]
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Couldn't read specified file for viewing")
		panic(err)
	}

	video := decodeVideo(file)
	video.play()
}
