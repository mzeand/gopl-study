package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type dup struct {
	num   int
	files map[string]int
}

func main() {
	counts := make(map[string]*dup)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if line != "" {
				if _, ok := counts[line]; !ok {
					counts[line] = &dup{
						files: make(map[string]int),
					}
				}
				counts[line].num++
				counts[line].files[filename]++
			}
		}
	}
	for line, d := range counts {
		if d.num > 1 {
			fmt.Printf("%d\t%s\t%v\n", d.num, line, d.files)
		}
	}
}
