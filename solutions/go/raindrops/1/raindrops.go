package raindrops

import (
	"strconv"
	"strings"
)


func Convert(number int) string {
	divisibleNumbers := getDivisibleNumbers(number)

	if len(divisibleNumbers) == 0 {
		return strconv.Itoa(number)
	}

	raindrops := make([]string, 0, 3)
	for _, num := range divisibleNumbers {
		switch {
		case num == 3:
			raindrops = append(raindrops, "Pling")
		case num == 5:
			raindrops = append(raindrops, "Plang")
		case num == 7:
			raindrops = append(raindrops, "Plong")
		}
	}
	return strings.Join(raindrops, "")
}

func getDivisibleNumbers(number int) []int {
	divisibleNumbers := make([]int, 0, 3)
	if testDivisibility(number, 3) {
		divisibleNumbers = append(divisibleNumbers, 3)
	}
	if testDivisibility(number, 5) {
		divisibleNumbers = append(divisibleNumbers, 5)
	}
	if testDivisibility(number, 7) {
		divisibleNumbers = append(divisibleNumbers, 7)
	}
	return divisibleNumbers
}

func testDivisibility(dividend int, divisor int) bool {
	if (dividend % divisor) == 0 {
		return true
	}
	return false
}
