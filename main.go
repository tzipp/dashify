package main

import (
	"os"
	"fmt"
	"path"
	"regexp"
	"log"
)

func main() {
	args := os.Args[1:]
	src := args[0]
	_, file := path.Split(src)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	dst := reg.ReplaceAllString(file, "-")

	fmt.Println(dst)
}
