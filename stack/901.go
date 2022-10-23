package stack

type StockSpanner struct {
	stack [][2]int
}

func Constructorq() StockSpanner {
	return StockSpanner{
		stack: make([][2]int, 0),
	}
}

// 901. Online Stock Span
// 901. 股票价格跨度
// 思路：单调栈（递减）
// time O(1) space O(n)
func (p *StockSpanner) Next(price int) int {
	ans := 1
	// stack pair value : [0-price, 1-answer]
	for len(p.stack) != 0 && p.stack[len(p.stack)-1][0] <= price {
		ans += p.stack[len(p.stack)-1][1]
		p.stack = p.stack[:len(p.stack)-1]
	}
	current := [2]int{price, ans}
	p.stack = append(p.stack, current)
	return ans
}
