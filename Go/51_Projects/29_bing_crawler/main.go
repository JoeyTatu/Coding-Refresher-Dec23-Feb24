package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := []string{"1", "3", "5", "7", "9", "11", "13", "15"}

	var answer int
	for answer = 30; answer <= 50; answer++ {
		fmt.Printf("Checking combinations for answer = %d\n", answer)
		var x, y, z string

		// Brute-force approach to find the combination
		for _, x = range numbers {
			for _, y = range numbers {
				for _, z = range numbers {
					if x != y && x != z && y != z {
						sum := parseAndSum(x, y, z)
						fmt.Printf("Trying combination: %s + %s + %s = %d\n", x, y, z, sum)
						if sum == answer {
							// Solution found
							fmt.Printf("The correct combination is %s + %s + %s = %d\n", x, y, z, answer)
						}
					}
				}
			}
		}
	}
}

func parseAndSum(nums ...string) int {
	sum := 0
	for _, num := range nums {
		val, _ := strconv.Atoi(num)
		sum += val
	}
	return sum
}

