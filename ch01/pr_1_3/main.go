package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t1 := measure(naive, os.Args)
	t2 := measure(join, os.Args)

	fmt.Printf("Naive: %d ns\n", t1)
	fmt.Printf("Join : %d ns\n", t2)
	fmt.Printf("Naive - Join = %d ns\n", t1-t2)
}

func naive(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func join(args []string) string {
	return strings.Join(args[1:], " ")
}

func measure(fn func(args []string) string, args []string) int64 {
	start := time.Now()
	fmt.Println(fn(args))
	return time.Since(start).Nanoseconds()
}
