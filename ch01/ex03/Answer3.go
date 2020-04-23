package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var inefficiencyExeTime = inefficiency()
	var efficiencyExeTime = efficiency()
	fmt.Printf("%d - %d = %dns \n", inefficiencyExeTime, efficiencyExeTime, inefficiencyExeTime-efficiencyExeTime)
}

func efficiency() int64 {
	start := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	return time.Since(start).Nanoseconds()
}

func inefficiency() int64 {
	start := time.Now()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	return time.Since(start).Nanoseconds()
}
