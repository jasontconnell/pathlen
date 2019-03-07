package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	p := flag.String("p", ".", "path to measure")
	l := flag.Int("l", 255, "length of path to flag")
	flag.Parse()

	paths := []string{}
	filepath.Walk(*p, func(path string, info os.FileInfo, err error) error {
		if len(path) >= *l {
			paths = append(paths, path)
		}
		return nil
	})

	for _, longpath := range paths {
		fmt.Println(longpath)
	}
}
