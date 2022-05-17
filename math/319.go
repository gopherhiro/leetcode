package math

import "math"

// 319. 灯泡开关
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
