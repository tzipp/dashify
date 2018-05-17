package main

import (
	"os"
	"regexp"
	"log"
	"strings"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	src := args[0]
	dir, file := filepath.Split(src)
	// fmt.Println("Directory: ", dir)
	// fmt.Println("Source file: ", file)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	ext := filepath.Ext(file)
	name := file[0:len(file)-len(ext)]

	dstName := reg.ReplaceAllString(name, "-")
	parts := []string{dir, dstName, ext}
	dst := strings.Join(parts, "")

	// fmt.Println("Destination: ", dst)

	err = os.Rename(src, dst)
	if err != nil {
		log.Fatal(err)
	}
}
