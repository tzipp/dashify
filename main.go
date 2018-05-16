package main

import (
	"os"
	"path"
	"regexp"
	"log"
	"strings"
	"fmt"
)

func main() {
	args := os.Args[1:]
	src := args[0]
	dir, file := path.Split(src)
	fmt.Println("Directory: ", dir)
	fmt.Println("Source file: ", file)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	dstFile := reg.ReplaceAllString(file, "-")
	parts := []string{dir, dstFile}
	dst := strings.Join(parts, "")

	fmt.Println("Destination: ", dst)

	os.Rename(src, dst)
}
