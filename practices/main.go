package main

import "fmt"

// Q1: FooBar

func fooBar(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FooBar")
		} else if i%3 == 0 {
			fmt.Println("Foo")
		} else if i%5 == 0 {
			fmt.Println("Bar")
		} else {
			fmt.Println(i)
		}
	}
}

// Q2: Two Sum

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, ok := numMap[complement]; ok {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return nil // No solution found
}

// Q3: Max Profit

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		fmt.Println("price:", price)
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func main() {
	// Q1
	// fooBar(15);

	// Q2
	// nums := []int{2, 7, 11, 15}
	// target := 9
	// result := twoSum(nums, target)
	// fmt.Println(result)

	// Q3
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices)) // Output: 5
}
