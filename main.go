package main

import (
	"os"
	"fmt"
	"path"
	"regexp"
	"log"
	"strings"
)

func main() {
	args := os.Args[1:]
	src := args[0]
	dir, file := path.Split(src)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	dstFile := reg.ReplaceAllString(file, "-")
	parts := []string{dir, dstFile}
	dst := strings.Join(parts, "")

	// os.Rename(src, dst)
	fmt.Println(src)
	fmt.Println(dst)
}
