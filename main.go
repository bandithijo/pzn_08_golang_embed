package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed logo.jpg
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("logo_next.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println("File name", entry.Name())
			file, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println("File content:", string(file))
		}
	}
}
