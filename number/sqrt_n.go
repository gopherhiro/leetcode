package number

import (
	"fmt"
	"math"
)

// sqrtN 使用二分法计算根号 n
func sqrtN(n float64) float64 {
	if n < 0 { // 处理负数输入
		return math.NaN()
	}
	if n == 0 { // 处理 0
		return 0
	}

	var left, right float64
	epsilon := 1e-10 // 精度控制

	// 确定初始区间
	if n < 1 {
		left = 0
		right = 1
	} else {
		left = 0
		right = n
	}

	for right-left > epsilon { // 当区间宽度大于精度时继续
		mid := (left + right) / 2 // 计算中点
		midSquare := mid * mid    // 计算 mid 的平方

		if math.Abs(midSquare-n) < epsilon { // 如果足够接近 n，直接返回
			return mid
		}
		if midSquare > n { // 中点太大，缩小右边界
			right = mid
		} else { // 中点太小，增大左边界
			left = mid
		}
	}
	return (left + right) / 2 // 返回最终区间的平均值
}

func main() {
	// 测试不同输入
	testCases := []float64{0, 0.25, 2, 3, 4, 100}
	for _, n := range testCases {
		result := sqrtN(n)
		fmt.Printf("sqrt(%g) 计算结果: %.10f, math.Sqrt: %.10f, 误差: %.10f\n",
			n, result, math.Sqrt(n), math.Abs(result-math.Sqrt(n)))
	}
}
