package main

import (
	"flag"
	"fmt"
)

func stats(email string) {
	print("stats")

}

func scan(folder string) {
	fmt.Printf("Found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filePath := getDotFilePath()
	addNewSliceElementToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "idk")
	flag.StringVar(&email, "email", "", "idk")

	flag.Parse()

	if folder != "" {
		return
	}

	stats(email)

}
