// https://www.hackerrank.com/challenges/mark-and-toys/problem

package main

import (
	"fmt"
	"sort"
)

func maximumToys(prices []int32, k int32) int32 {
	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})

	var totalPrice = int32(0)
	var toysCount = int32(0)

	for i := 0; i < len(prices); i++ {
		if totalPrice+prices[i] < k {
			totalPrice += prices[i]
			toysCount++
		} else {
			break
		}
	}

	return toysCount
}

func main() {
	fmt.Println(maximumToys([]int32{1, 12, 5, 111, 200, 1000, 10}, 50))
}
