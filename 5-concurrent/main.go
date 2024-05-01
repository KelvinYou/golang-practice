package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// without concurrent
	startTime := time.Now()
	result := sumOfSquares(nums)
	endTime := time.Now()
	fmt.Println("Sum of squares:", result)
	fmt.Println("startTime:", startTime)
	fmt.Println("endTime:", endTime)
	fmt.Println("Total time taken without concurrency:", endTime.Sub(startTime))

	// with concurrent
	startTime = time.Now()
	result = concurrentSumOfSquares(nums)
	endTime = time.Now()
	fmt.Println("Sum of squares:", result)
	fmt.Println("startTime:", startTime)
	fmt.Println("endTime:", endTime)
	fmt.Println("Total time taken:", endTime.Sub(startTime))
}

func sumOfSquares(nums []int) int {
	sum := 0
	for _, num := range nums {
		square := num * num
		sum += square
	}
	return sum
}

func concurrentSumOfSquares(nums []int) int {
	resultChan := make(chan int)
	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d handling task: square of %d\n", goroutineID(), n)
			resultChan <- n * n
		}(num)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	sum := 0
	for square := range resultChan {
		sum += square
	}

	return sum
}

// goroutineID returns an ID of the current goroutine
func goroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	var id int
	fmt.Sscanf(string(buf[:n]), "goroutine %d", &id)
	return id
}
