package stack

import (
	"fmt"
	"testing"
)

func TestCanSeePersonsCount(t *testing.T) {
	heights := []int{2, 1, 2, 4, 3}
	r := canSeePersonsCount(heights)
	fmt.Println(r)
}

func TestDailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	r := dailyTemperatures(temperatures)
	fmt.Println(r)
}
