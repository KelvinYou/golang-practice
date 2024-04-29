package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	numThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Number of CPU cores:", numCPU)
	fmt.Println("Number of threads (logical processors):", numThreads)
}
