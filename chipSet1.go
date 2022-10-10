package main

import (
	"fmt"
)

func getChipSet1Functions() map[string]func(int) int {
	m := make(map[string]func(int) int)

	m["white"] = func(value int) int {
		pot.whites += value
		return value
	}

	m["orange"] = func(value int) int {
		return value
	}

	m["green"] = func(value int) int {
		return value
	}

	m["red"] = func(value int) int {
		numOranges := 0
		for _, chip := range pot.ingredients {
			if chip.color == "orange" {
				numOranges += 1
			}
		}

		if numOranges == 1 || numOranges == 2 {
			value += 1
		} else if numOranges > 3 {
			value += 2
		}
		return value
	}

	m["blue"] = func(value int) int {
		var chips []int
		var response int

		switch value {
		case 1:
			chips = getUniqueNumbers(1, len(bag))
		case 2:
			chips = getUniqueNumbers(2, len(bag))
		case 4:
			chips = getUniqueNumbers(4, len(bag))
		}
		fmt.Println("Blue effect: enter the number for the chip you'd like to add to your pot (0 for none)")
		for _, v := range chips {
			fmt.Println(bag[v].color, bag[v].value)
		}
		for {
			fmt.Scanln(&response)

			if response == 0 {
				return value
			}
			for i := 1; i <= len(chips); i++ {
				if response == i {
					transferChip(response)
					return value
				}
			}
			fmt.Println("Please enter a valid number")
		}
	}

	return m
}
