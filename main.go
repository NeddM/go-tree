package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	directory := "."
	if len(os.Args) >= 2 {
		directory = os.Args[1]
	}

	absPath, err := filepath.Abs(directory)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = listFiles(absPath, "")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func listFiles(dirPath, indent string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for i, file := range files {
		filename := file.Name()
		if i == len(files)-1 {
			fmt.Printf("%s└── %s\n", indent, filename)
			if file.IsDir() {
				err = listFiles(filepath.Join(dirPath, filename), indent+"    ")
			}
		} else {
			fmt.Printf("%s├── %s\n", indent, filename)
			if file.IsDir() {
				err = listFiles(filepath.Join(dirPath, filename), indent+"│   ")
			}
		}
		if err != nil {
			return err
		}
	}

	return nil
}
