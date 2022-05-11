package dp

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	r := coinChangeL(coins, amount)
	r1 := coinChangeM(coins, amount)
	r2 := coinChangeR(coins, amount)
	fmt.Println(r, r1, r2)
}
