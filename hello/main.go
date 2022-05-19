package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(10)

	if n == 7 {
		os.Exit(1)
	}

	fmt.Printf("Hello world %d\n", n)
}
