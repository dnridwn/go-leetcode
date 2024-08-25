package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/flytam/filenamify"
)

func main() {
	handleGenerateLeetCodeProblemFile()
}

func handleGenerateLeetCodeProblemFile() {
	// read arguments from command
	problemName := flag.String("name", "", "Put the LeetCode problem name")
	flag.Parse()

	// validate LeetCode problem name
	if *problemName == "" {
		fmt.Println("Please put the LeetCode problem name")
		return
	}

	// creating folder
	validPathName, err := filenamify.Filenamify(*problemName, filenamify.Options{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// stop if folder is exists
	if _, err := os.Open(validPathName); !os.IsNotExist(err) {
		fmt.Println("Folder " + validPathName + " is exists")
		return
	}

	// create folder
	if err := os.Mkdir(validPathName, 0755); err != nil {
		fmt.Println(err.Error())
		return
	}

	// define default file contents
	fileContents := "package " + validPathName + "\n"

	// list of file to process
	files := []string{
		validPathName + "/" + validPathName + ".go",
		validPathName + "/" + validPathName + "_test.go",
	}

	// create and write to file
	for _, file := range files {
		fileOs, err := os.Create(file)
		if err != nil {
			fileOs.Close()
			fmt.Println(err.Error())
			break
		}

		if _, err := fileOs.WriteString(fileContents); err != nil {
			fileOs.Close()
			fmt.Println(err.Error())
			break
		}

		fileOs.Close()
	}

	// print success
	fmt.Println("Successfully generating files")
}
