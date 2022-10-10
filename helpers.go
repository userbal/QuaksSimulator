package main

import (
	"fmt"
	"math/rand"
)

func getUniqueNumbers(num int, limit int) []int {
	var nums []int
	for i := 0; i < num; i++ {
		for {
			r := rand.Intn(limit)
			unique := true
			for _, n := range nums {
				if r == n {
					unique = true
				}
			}
			if unique {
				nums = append(nums, r)
				continue
			}
		}
	}
	return nums
}

func showBag() {
	fmt.Println("Bag:")
	for _, chip := range bag {
		fmt.Println(chip.color, chip.value)
	}
}
