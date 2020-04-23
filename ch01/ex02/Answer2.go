package main

import (
	"fmt"
	"os"
)

func main() {
	for i, s := range os.Args {
		fmt.Printf("index[%d]: %s \n", i, s)
	}
}
