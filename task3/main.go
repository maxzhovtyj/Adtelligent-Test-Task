package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	arg := "internal"

	err := ReadFilesFromDir(arg)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFilesFromDir(arg string) error {
	dir, err := os.ReadDir(arg)
	if err != nil {
		return fmt.Errorf("%s - %v", dir, err)
	}

	for _, curr := range dir {
		if curr.IsDir() {
			currDirPath := fmt.Sprintf("%s/%s", arg, curr.Name())

			err = ReadFilesFromDir(currDirPath)
			if err != nil {
				return fmt.Errorf("%s - %v", curr.Name(), err)
			}
		} else {
			fmt.Println(curr.Name())
		}
	}

	return nil
}
